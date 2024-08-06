/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/code-generator/cmd/validation-gen/validators"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"
)

var (
	fieldPkg      = "k8s.io/apimachinery/pkg/util/validation/field"
	errorListType = types.Name{Package: fieldPkg, Name: "ErrorList"}
	fieldPathType = types.Name{Package: fieldPkg, Name: "Path"}
	errorfType    = types.Name{Package: "fmt", Name: "Errorf"}
	runtimePkg    = "k8s.io/apimachinery/pkg/runtime"
	schemeType    = types.Name{Package: runtimePkg, Name: "Scheme"}
)

// genValidations produces a file with autogenerated validations.
type genValidations struct {
	generator.GoGenerator
	outputPackage       string
	inputToPkg          map[string]string // Maps input packages to generated validation packages
	rootTypes           []*types.Type
	discovered          *typeDiscoverer
	imports             namer.ImportTracker
	validator           validators.DeclarativeValidator
	hasValidationsCache map[*typeNode]bool
}

// NewGenValidations cretes a new generator for the specified package.
func NewGenValidations(outputFilename, outputPackage string, rootTypes []*types.Type, discovered *typeDiscoverer, inputToPkg map[string]string, validator validators.DeclarativeValidator) generator.Generator {
	return &genValidations{
		GoGenerator: generator.GoGenerator{
			OutputFilename: outputFilename,
		},
		outputPackage:       outputPackage,
		inputToPkg:          inputToPkg,
		rootTypes:           rootTypes,
		discovered:          discovered,
		imports:             generator.NewImportTrackerForPackage(outputPackage),
		validator:           validator,
		hasValidationsCache: map[*typeNode]bool{},
	}
}

func (g *genValidations) Namers(_ *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genValidations) Filter(_ *generator.Context, t *types.Type) bool {
	// We want to emit code for all root types.
	for _, rt := range g.rootTypes {
		if rt == t {
			return true
		}
	}
	// We want to emit for any other type that is transitively part of a root
	// type and has validations.
	n := g.discovered.typeNodes[t]
	return n != nil && g.hasValidations(n)
}

func (g *genValidations) Imports(_ *generator.Context) (imports []string) {
	var importLines []string
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func (g *genValidations) isOtherPackage(pkg string) bool {
	if pkg == g.outputPackage {
		return false
	}
	if strings.HasSuffix(pkg, `"`+g.outputPackage+`"`) {
		return false
	}
	return true
}

func (g *genValidations) Init(c *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	scheme := c.Universe.Type(schemeType)
	schemePtr := &types.Type{
		Kind: types.Pointer,
		Elem: scheme,
	}
	sw.Do("func init() { localSchemeBuilder.Register(RegisterValidations)}\n\n", nil)

	sw.Do("// RegisterValidations adds validation functions to the given scheme.\n", nil)
	sw.Do("// Public to allow building arbitrary schemes.\n", nil)
	sw.Do("func RegisterValidations(scheme $.|raw$) error {\n", schemePtr)
	for _, t := range g.rootTypes {
		node := g.discovered.typeNodes[t]
		if node == nil {
			// Should never happen.
			klog.Fatalf("found nil node for root-type %v", t)
		}

		// TODO: It would be nice if these were not hard-coded.
		var statusType *types.Type
		var statusField string
		if status := node.lookupField("status"); status != nil {
			statusType = status.node.valueType
			statusField = status.name
		}

		targs := generator.Args{
			"rootType":    t,
			"statusType":  statusType,
			"statusField": statusField,
			"errorList":   c.Universe.Type(errorListType),
			"fieldPath":   c.Universe.Type(fieldPathType),
			"fmtErrorf":   c.Universe.Type(errorfType),
		}
		//TODO: can this be (*$.rootType|raw$)(nil) ?
		sw.Do("scheme.AddValidationFunc(new($.rootType|raw$), func(obj, oldObj interface{}, subresources ...string) $.errorList|raw$ {\n", targs)
		sw.Do("  if len(subresources) == 0 {\n", targs)
		sw.Do("    return $.rootType|objectvalidationfn$(obj.(*$.rootType|raw$), nil)\n", targs)
		sw.Do("  }\n", targs)

		if statusType != nil {
			sw.Do("  if len(subresources) == 1 && subresources[0] == \"status\" {\n", targs)
			if g.hasValidations(g.discovered.typeNodes[statusType]) {
				sw.Do("    root := obj.(*$.rootType|raw$)\n", targs)
				sw.Do("    return $.statusType|objectvalidationfn$(&root.$.statusField$, nil)\n", targs)
			} else {
				sw.Do("    return nil // $.statusType|raw$ has no validation\n", targs)
			}
			sw.Do("  }\n", targs)
		}
		sw.Do("  return $.errorList|raw${field.InternalError(nil, $.fmtErrorf|raw$(\"no validation found for %T, subresources: %v\", obj, subresources))}\n", targs)
		sw.Do("})\n", targs)

		// TODO: Support update validations
		//       This will require correlating old object.
	}
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

func (g *genValidations) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	klog.V(5).Infof("emitting validation code for type %v", t)

	var errs []error
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	g.emitValidationFunction(c, t, sw)
	if err := sw.Error(); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func (g *genValidations) hasValidations(n *typeNode) bool {
	if r, found := g.hasValidationsCache[n]; found {
		return r
	}
	r := g.hasValidationsMiss(n)
	g.hasValidationsCache[n] = r
	return r
}

// Called in case of a cache miss.
func (g *genValidations) hasValidationsMiss(n *typeNode) bool {
	if len(n.typeValidations) > 0 {
		return true
	}
	allChildren := n.fields
	if n.key != nil {
		allChildren = append(allChildren, n.key)
	}
	if n.elem != nil {
		allChildren = append(allChildren, n.elem)
	}
	for _, c := range allChildren {
		if len(c.fieldValidations)+len(c.keyValidations)+len(c.elemValidations) > 0 {
			return true
		}
		if g.hasValidations(c.node) {
			return true
		}
	}
	return false
}

func (g *genValidations) emitValidationFunction(c *generator.Context, t *types.Type, sw *generator.SnippetWriter) {
	targs := generator.Args{
		"inType":    t,
		"errorList": c.Universe.Type(errorListType),
		"fieldPath": c.Universe.Type(fieldPathType),
	}

	sw.Do("func $.inType|objectvalidationfn$(obj *$.inType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
	node := g.discovered.typeNodes[t]
	if node == nil {
		// Should never happen.
		klog.Fatalf("found nil node for type %v", t)
	}
	fakeChild := &childNode{
		node:      node,
		childType: t,
	}
	g.emitValidationForChild(c, fakeChild, true, sw)
	sw.Do("return errs\n", nil)
	sw.Do("}\n\n", nil)
}

// typeDiscoverer contains fields necessary to build graphs of types.
type typeDiscoverer struct {
	validator  validators.DeclarativeValidator
	inputToPkg map[string]string

	// typeNodes holds a map of gengo Type to typeNode for all of the types
	// encountered during discovery.
	typeNodes map[*types.Type]*typeNode
}

// NewTypeDiscoverer creates and initializes a NewTypeDiscoverer.
func NewTypeDiscoverer(validator validators.DeclarativeValidator, inputToPkg map[string]string) *typeDiscoverer {
	return &typeDiscoverer{
		validator:  validator,
		inputToPkg: inputToPkg,
		typeNodes:  map[*types.Type]*typeNode{},
	}
}

// typeNode represents a node in the type-graph, annotated with information
// about validations.  Everything in this type, transitively, is assoctiated
// with the type, and not any specific instance of that type (e.g. when used as
// a field in a struct.
type typeNode struct {
	valueType *types.Type // never a pointer, but may be a map, slice, struct, etc.
	funcName  types.Name  // populated when this type is "opaque"

	fields []*childNode // populated when this type is a struct
	key    *childNode   // populated when this type is a map
	elem   *childNode   // populated when this type is a map or slice

	// These are validations defined on the type.
	typeValidations []validators.FunctionGen
}

// Dump returns a multi-line string which represents a typeNode and its
// children.  This can be used to debug the type graph.
func (n *typeNode) Dump() string {
	buf := bytes.Buffer{}
	visited := map[*typeNode]bool{}
	buf.WriteString(fmt.Sprintf("type %s {\n", n.valueType))
	n.doDump(&buf, 1, visited)
	buf.WriteString("}")
	return buf.String()
}

func (n *typeNode) dumpIndent(buf *bytes.Buffer, indent int) {
	for i := 0; i < indent; i++ {
		buf.WriteString("    ")
	}
}

func (n *typeNode) doDump(buf *bytes.Buffer, indent int, visited map[*typeNode]bool) {
	if visited[n] {
		n.dumpIndent(buf, indent)
		buf.WriteString("(recursive)\n")
		return
	}
	visited[n] = true

	for _, val := range n.typeValidations {
		n.dumpIndent(buf, indent)
		fn, args := val.SignatureAndArgs()
		buf.WriteString(fmt.Sprintf("type-validation: %v(%+v)\n", fn, args))
	}
	n.dumpChildren(buf, indent, visited)
}

func (n *typeNode) dumpChildren(buf *bytes.Buffer, indent int, visited map[*typeNode]bool) {
	for _, fld := range n.fields {
		n.dumpIndent(buf, indent)
		buf.WriteString(fmt.Sprintf("field %s: %s {\n", fld.name, fld.childType))
		for _, val := range fld.fieldValidations {
			fn, args := val.SignatureAndArgs()
			n.dumpIndent(buf, indent+1)
			buf.WriteString(fmt.Sprintf("field-validation: %v(%+v)\n", fn, args))
		}
		for _, val := range fld.keyValidations {
			fn, args := val.SignatureAndArgs()
			n.dumpIndent(buf, indent+1)
			buf.WriteString(fmt.Sprintf("key-validation: %v(%+v)\n", fn, args))
		}
		for _, val := range fld.elemValidations {
			fn, args := val.SignatureAndArgs()
			n.dumpIndent(buf, indent+1)
			buf.WriteString(fmt.Sprintf("val-validation: %v(%+v)\n", fn, args))
		}
		fld.node.doDump(buf, indent+1, visited)
		n.dumpIndent(buf, indent)
		buf.WriteString("}\n")
	}
}

// childNode represents a type which is used in another type (e.g. a struct
// field).
type childNode struct {
	name      string      // the field name in the parent, populated when this node is a struct field
	jsonName  string      // always populated when name is populated
	childType *types.Type // the real type of the child (may be a pointer)
	node      *typeNode   // the node of the child's value type

	fieldValidations []validators.FunctionGen // validations on the field
	keyValidations   []validators.FunctionGen // validations on each key of a map field
	elemValidations  []validators.FunctionGen // validations on each value of a list or map
}

// lookupField returns the childNode with the specified JSON name.
func (n typeNode) lookupField(jsonName string) *childNode {
	for _, fld := range n.fields {
		if fld.jsonName == jsonName {
			return fld
		}
	}
	return nil
}

const (
	// This tag defines a validation which is to be run on each key in a map.
	eachKeyTag = "eachKey"
	// This tag defines a validation which is to be run on each value in a map
	// or slice.
	eachValTag = "eachVal"
)

// DiscoverType walks the given type recursively, building a type-graph in this
// typeDiscoverer.  If this is called multiple times for different types, the
// multiple graphs will be stored, and where types overlap, they will be
// merged.
func (td *typeDiscoverer) DiscoverType(t *types.Type) error {
	fldPath := field.NewPath(t.Name.String())
	if node, err := td.discover(t, fldPath); err != nil {
		return err
	} else {
		fmt.Println(node.Dump()) //FIXME: remove
	}
	return nil
}

// discover walks the given type recursively and returns a typeNode
// representing it.
func (td *typeDiscoverer) discover(t *types.Type, fldPath *field.Path) (*typeNode, error) {
	if t.Kind == types.Pointer {
		if t.Elem.Kind == types.Pointer {
			return nil, fmt.Errorf("field %s (%s): pointers to pointers are not supported", fldPath.String(), t)
		}
		// Remove pointerness.
		t = t.Elem
	}
	// If we have done this type already, we can stop here and break any
	// recursion.
	if node := td.typeNodes[t]; node != nil {
		return node, nil
	}

	// If we are descending into a named type, reboot the field path for better
	// logging.  Otherwise the field path might come in as something like
	// <type1>.<field1>.<field2> which is true, but not super useful.
	switch t.Kind {
	case types.Alias, types.Struct:
		fldPath = field.NewPath(t.Name.String())
	}

	thisNode := &typeNode{
		valueType: t,
	}
	td.typeNodes[t] = thisNode

	// Extract any type-attached validation rules.
	// TODO: do eachVal and eachKey for aliases to slices/maps
	if validations, err := td.validator.ExtractValidations(t.Name.Name, t, t.CommentLines); err != nil {
		return nil, fmt.Errorf("%v: %w", fldPath, err)
	} else {
		if len(validations) > 0 {
			klog.V(5).InfoS("  found type-attached validations", "n", len(validations))
			thisNode.typeValidations = validations
		}
	}

	// If this is an opaque, named type, we can call its validation function.
	switch t.Kind {
	case types.Alias, types.Struct:
		if fn, ok := td.getValidationFunctionName(t); !ok {
			return thisNode, nil
		} else {
			thisNode.funcName = fn
		}
	}

	switch t.Kind {
	case types.Builtin:
		// Nothing more to do.
	case types.Alias:
		// Note: By the language definition, what gengo calls "Aliases" (really
		// just "type definitions") have underlying types of the type literal.
		// In other words, if we define `type T1 string` and `type T2 T1`, the
		// underlying type of T2 is string, not T1.  This means that:
		//    1) We will emit code for both underlying types. If the underlying
		//       type is a struct with many fields, we will emit two identical
		//       functions.
		//    2) Validating a field of type T2 will NOT call any validation
		//       defined on the type T1.
		//    3) In the case of a type definition whose RHS is a struct which
		//       has fields with validation tags, the validation for those fields
		//       WILL be called from the generated for for the new type.
		if t.Underlying.Kind == types.Pointer {
			return nil, fmt.Errorf("field %s (%v): typedefs of pointers are not supported", fldPath.String(), t)
		}
		if _, err := td.discover(t.Underlying, fldPath); err != nil {
			return nil, err
		}
	case types.Struct:
		// Discover into this struct, recursively.
		if err := td.discoverStruct(thisNode, fldPath); err != nil {
			return nil, err
		}
	case types.Slice, types.Array:
		// Discover the element type.
		if node, err := td.discover(t.Elem, fldPath.Key("vals")); err != nil {
			return nil, err
		} else {
			thisNode.elem = &childNode{
				childType: t.Elem,
				node:      node,
			}
		}
	case types.Map:
		// Discover the key type.
		if node, err := td.discover(t.Key, fldPath.Key("keys")); err != nil {
			return nil, err
		} else {
			thisNode.key = &childNode{
				childType: t.Key,
				node:      node,
			}
		}

		// Discover the element type.
		if node, err := td.discover(t.Elem, fldPath.Key("vals")); err != nil {
			return nil, err
		} else {
			thisNode.elem = &childNode{
				childType: t.Elem,
				node:      node,
			}
		}
	default:
		return nil, fmt.Errorf("field %s (%v): kind %v is not supported", fldPath.String(), t, t.Kind)
	}

	return thisNode, nil
}

// discoverStruct walks a struct type recursively.
func (td *typeDiscoverer) discoverStruct(thisNode *typeNode, fldPath *field.Path) error {
	var fields []*childNode

	// Discover into each field of this struct.
	for _, memb := range thisNode.valueType.Members {
		name := memb.Name
		if len(name) == 0 { // embedded fields
			if memb.Type.Kind == types.Pointer {
				name = memb.Type.Elem.Name.Name
			} else {
				name = memb.Type.Name.Name
			}
		}
		// Only do exported fields.
		if unicode.IsLower([]rune(name)[0]) {
			continue
		}
		// If we try to emit code for this field and find no JSON name, we
		// will abort.
		jsonName := ""
		if tags, ok := lookupJSONTags(memb); ok {
			jsonName = tags.name
		}

		klog.V(5).InfoS("  field", "name", name, "jsonName", jsonName, "type", memb.Type)

		// Discover the field type.
		childPath := fldPath.Child(name)
		childType := memb.Type
		var child *childNode
		if node, err := td.discover(childType, childPath); err != nil {
			return err
		} else {
			child = &childNode{
				name:      name,
				jsonName:  jsonName,
				childType: childType,
				node:      node,
			}
		}

		// Extract any field-attached validation rules.
		if validations, err := td.validator.ExtractValidations(name, memb.Type, memb.CommentLines); err != nil {
			return fmt.Errorf("field %s: %w", childPath.String(), err)
		} else {
			if len(validations) > 0 {
				klog.V(5).InfoS("  found field-attached value-validations", "n", len(validations))
				child.fieldValidations = append(child.fieldValidations, validations...)
			}
		}

		// Add any other field-attached "special" validators.
		switch childType.Kind {
		case types.Slice, types.Array:
			//TODO: also support +k8s:eachVal
			if tagVals, found := gengo.ExtractCommentTags("+", memb.CommentLines)[eachValTag]; found {
				for _, tagVal := range tagVals {
					fakeComments := []string{tagVal}
					// Extract any embedded list-validation rules.
					if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[vals]", memb.Name), childType.Elem, fakeComments); err != nil {
						return fmt.Errorf("%v: %w", childPath.Key("vals"), err)
					} else {
						if len(validations) > 0 {
							klog.V(5).InfoS("  found list-validations", "n", len(validations))
							child.elemValidations = validations
						}
					}
				}
			}
		case types.Map:
			//TODO: also support +k8s:eachKey
			if tagVals, found := gengo.ExtractCommentTags("+", memb.CommentLines)[eachKeyTag]; found {
				for _, tagVal := range tagVals {
					fakeComments := []string{tagVal}
					// Extract any embedded key-validation rules.
					if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[keys]", memb.Name), childType.Key, fakeComments); err != nil {
						return fmt.Errorf("%v: %w", childPath.Key("keys"), err)
					} else {
						if len(validations) > 0 {
							klog.V(5).InfoS("  found key-validations", "n", len(validations))
							child.keyValidations = validations
						}
					}
				}
			}
			//TODO: also support +k8s:eachVal
			if tagVals, found := gengo.ExtractCommentTags("+", memb.CommentLines)[eachValTag]; found {
				for _, tagVal := range tagVals {
					fakeComments := []string{tagVal}
					// Extract any embedded list-validation rules.
					if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[vals]", memb.Name), childType.Elem, fakeComments); err != nil {
						return fmt.Errorf("%v: %w", childPath.Key("vals"), err)
					} else {
						if len(validations) > 0 {
							klog.V(5).InfoS("  found list-validations", "n", len(validations))
							child.elemValidations = validations
						}
					}
				}
			}
		}

		fields = append(fields, child)
	}

	thisNode.fields = fields
	return nil
}

// getValidationFunctionName looks up the name of the specified type's
// validation function.
//
// TODO: Currently this is a "blind" call - we hope that the expected function
// exists, but we don't verify that, and we only emit calls into packages which
// are being processed by this generator. For cross-package calls we will need
// to verify the target, either by naming convention + fingerprint or by
// explicit comment-tags or something.
func (td *typeDiscoverer) getValidationFunctionName(t *types.Type) (types.Name, bool) {
	pkg, ok := td.inputToPkg[t.Name.Package]
	if !ok {
		return types.Name{}, false
	}
	return types.Name{Package: pkg, Name: "Validate_" + t.Name.Name}, true
}

// emitValidationForChild emits code for the specified childNode, calling
// type-attached validations and then descending into the type (e.g. struct
// fields).
//
// Emitted code can assume that the value in question is always named "obj" and
// the field path to this value is named "fldPath".  objIsPtr indicates whether
// the "obj" variable is a pointer.
func (g *genValidations) emitValidationForChild(c *generator.Context, thisChild *childNode, objIsPtr bool, sw *generator.SnippetWriter) {
	thisNode := thisChild.node
	inType := thisNode.valueType

	targs := generator.Args{
		"inType":    inType,
		"errorList": c.Universe.Type(errorListType),
		"fieldPath": c.Universe.Type(fieldPathType),
	}

	didSome := false // for prettier output later

	// Emit code for type-attached validations.
	if validations := thisNode.typeValidations; len(validations) > 0 {
		sw.Do("// type $.inType|raw$\n", targs)
		emitCallsToValidators(c, validations, objIsPtr, sw)
		sw.Do("\n", nil)
		didSome = true
	}

	// Descend into the type.
	switch inType.Kind {
	case types.Builtin:
		// Nothing further.
	case types.Alias:
		// Nothing further.
	case types.Struct:
		for _, fld := range thisNode.fields {
			if len(fld.name) == 0 {
				panic(fmt.Sprintf("missing field name in type %s (field-type %s)", thisNode.valueType, fld.childType))
			}
			// Missing JSON name is checked iff we have code to emit.

			targs := targs.WithArgs(generator.Args{
				"fieldName": fld.name,
				"fieldJSON": fld.jsonName,
				"fieldType": fld.childType,
			})

			childIsPtr := fld.childType.Kind == types.Pointer

			// Accumulate into a buffer so we don't emit empty functions.
			buf := bytes.NewBuffer(nil)
			bufsw := sw.Dup(buf)

			validations := fld.fieldValidations
			if len(validations) > 0 {
				// When calling registered validators, we always pass the
				// underlying value-type.  E.g. if the field's type is string,
				// we pass string, and if the field's type is *string, we also
				// pass string (checking for nil, first).  This means those
				// validators don't have to know the difference, but it also
				// means that large structs will be passed by value.  If this
				// turns out to be a real problem, we could change this to pass
				// everything by pointer.
				emitCallsToValidators(c, validations, childIsPtr, bufsw)
			}

			// Get to the real type.
			switch fld.node.valueType.Kind {
			case types.Struct, types.Alias:
				// If this field is another type, call its validation function.
				// Checking for nil is handled inside this call.
				g.emitCallToOtherTypeFunc(c, fld.node, childIsPtr, bufsw)
			default:
				// Descend into this field.
				g.emitValidationForChild(c, fld, childIsPtr, bufsw)
			}

			if buf.Len() > 0 {
				if len(fld.jsonName) == 0 {
					panic(fmt.Sprintf("missing JSON name for field %s.%s", fld.node.valueType, fld.name))
				}

				if didSome {
					sw.Do("\n", nil)
				}
				sw.Do("// field $.inType|raw$.$.fieldName$\n", targs)
				sw.Do("errs = append(errs,\n", targs)
				sw.Do("  func(obj $.fieldType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
				sw.Append(buf)
				sw.Do("    return\n", targs)
				sw.Do("  }(obj.$.fieldName$, fldPath.Child(\"$.fieldJSON$\"))...)\n", targs)
				sw.Do("\n", nil)
			} else {
				sw.Do("// field $.inType|raw$.$.fieldName$ has no validation\n", targs)
			}
			didSome = true
		}
	case types.Slice, types.Array:
		targs := targs.WithArgs(generator.Args{
			"elemType": inType.Elem,
		})

		elemIsPtr := inType.Elem.Kind == types.Pointer

		// Accumulate into a buffer so we don't emit empty functions.
		elemBuf := bytes.NewBuffer(nil)
		elemSW := sw.Dup(elemBuf)

		// Validate each value.
		validations := thisChild.elemValidations
		if len(validations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			emitCallsToValidators(c, validations, elemIsPtr, elemSW)
		}

		switch thisNode.elem.node.valueType.Kind {
		case types.Struct, types.Alias:
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, thisNode.elem.node, elemIsPtr, elemSW)
		default:
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		if elemBuf.Len() > 0 {
			sw.Do("for i, val := range obj {\n", targs)
			sw.Do("  errs = append(errs,\n", targs)
			sw.Do("    func(obj $.elemType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
			sw.Append(elemBuf)
			sw.Do("      return\n", targs)
			sw.Do("    }(val, fldPath.Index(i))...)\n", targs)
			sw.Do("}\n", nil)
		}
	case types.Map:
		targs := targs.WithArgs(generator.Args{
			"keyType": inType.Key,
			"valType": inType.Elem,
		})

		keyIsPtr := inType.Key.Kind == types.Pointer
		valIsPtr := inType.Elem.Kind == types.Pointer

		// Accumulate into a buffer so we don't emit empty functions.
		keyBuf := bytes.NewBuffer(nil)
		keySW := sw.Dup(keyBuf)

		// Validate each key.
		keyValidations := thisChild.keyValidations
		if len(keyValidations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			emitCallsToValidators(c, keyValidations, keyIsPtr, keySW)
		}

		switch thisNode.key.node.valueType.Kind {
		case types.Struct, types.Alias:
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, thisNode.key.node, keyIsPtr, keySW)
		default:
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		// Accumulate into a buffer so we don't emit empty functions.
		valBuf := bytes.NewBuffer(nil)
		valSW := sw.Dup(valBuf)

		// Validate each value.
		valValidations := thisChild.elemValidations
		if len(valValidations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			emitCallsToValidators(c, valValidations, valIsPtr, valSW)
		}

		switch thisNode.elem.node.valueType.Kind {
		case types.Struct, types.Alias:
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, thisNode.elem.node, valIsPtr, valSW)
		default:
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		kName, vName := "_", "_"
		if keyBuf.Len() > 0 {
			kName = "key"
		}
		if valBuf.Len() > 0 {
			vName = "val"
		}
		if keyBuf.Len()+valBuf.Len() > 0 {
			sw.Do("for $.key$, $.val$ := range obj {\n", targs.WithArgs(generator.Args{"key": kName, "val": vName}))
			if keyBuf.Len() > 0 {
				sw.Do("  errs = append(errs,\n", targs)
				sw.Do("    func(obj $.keyType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
				sw.Append(keyBuf)
				sw.Do("      return\n", targs)
				sw.Do("    }(key, fldPath)...)\n", targs)
			}
			if valBuf.Len() > 0 {
				sw.Do("  errs = append(errs,\n", targs)
				sw.Do("    func(obj $.valType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
				sw.Append(valBuf)
				sw.Do("      return\n", targs)
				sw.Do("    }(val, fldPath.Key(key))...)\n", nil) // TODO: what if the key is not a string?
			}
			sw.Do("}\n", nil)
		}
	default:
		klog.Fatalf("unhandled type: %v (%s)", inType, inType.Kind)
	}
}

// emitCallToOtherTypeFunc generates a call to the specified node's generated
// validation function for a field in some parent context.
//
// Emitted code can assume that the value in question is always named "obj" and
// the field path to this value is named "fldPath".  objIsPtr indicates whether
// the "obj" variable is a pointer.
func (g *genValidations) emitCallToOtherTypeFunc(c *generator.Context, node *typeNode, objIsPtr bool, sw *generator.SnippetWriter) {
	// If this type has no validations (transitively) then we don't need to do
	// anything.
	if !g.hasValidations(node) {
		return
	}

	addr := "" // adjusted below if needed
	if objIsPtr {
		sw.Do("if obj != nil {\n", nil)
		defer func() {
			sw.Do("}\n", nil)
		}()
	} else {
		addr = "&"
	}

	targs := generator.Args{
		"addr":     addr,
		"funcName": c.Universe.Type(node.funcName),
	}
	sw.Do("errs = append(errs, $.funcName|raw$($.addr$obj, fldPath)...)\n", targs)
}

// emitCallsToValidators emits calls to a list of validation functions for
// a single field or type. validations is a list of functions to call, with
// arguments.
//
// Emitted code can assume that the value in question is always named "obj" and
// the field path to this value is named "fldPath".  objIsPtr indicates whether
// the "obj" variable is a pointer.
func emitCallsToValidators(c *generator.Context, validations []validators.FunctionGen, objIsPtr bool, sw *generator.SnippetWriter) {
	// Helper func
	sort := func(in []validators.FunctionGen) []validators.FunctionGen {
		fatal := make([]validators.FunctionGen, 0, len(in))
		fatalPtr := make([]validators.FunctionGen, 0, len(in))
		nonfatal := make([]validators.FunctionGen, 0, len(in))
		nonfatalPtr := make([]validators.FunctionGen, 0, len(in))

		for _, fg := range in {
			isFatal := (fg.Flags().IsSet(validators.IsFatal))
			isPtrOK := (fg.Flags().IsSet(validators.PtrOK))

			if isFatal {
				if isPtrOK {
					fatalPtr = append(fatalPtr, fg)
				} else {
					fatal = append(fatal, fg)
				}
			} else {
				if isPtrOK {
					nonfatalPtr = append(nonfatalPtr, fg)
				} else {
					nonfatal = append(nonfatal, fg)
				}
			}
		}
		result := fatalPtr
		result = append(result, fatal...)
		result = append(result, nonfatalPtr...)
		result = append(result, nonfatal...)
		return result
	}

	validations = sort(validations)

	insideNilCheck := false
	for _, v := range validations {
		ptrOK := (v.Flags().IsSet(validators.PtrOK))
		isFatal := (v.Flags().IsSet(validators.IsFatal))

		fn, extraArgs := v.SignatureAndArgs()
		targs := generator.Args{
			"funcName": c.Universe.Type(fn),
			"deref":    "", // updated below if needed
		}
		if objIsPtr && !ptrOK {
			if !insideNilCheck {
				sw.Do("if obj != nil {\n", targs)
				insideNilCheck = true
			}
			targs["deref"] = "*"
		} else {
			if insideNilCheck {
				sw.Do("}\n", nil)
				insideNilCheck = false
			}
		}

		emitCall := func() {
			sw.Do("$.funcName|raw$(fldPath, $.deref$obj", targs)
			for _, arg := range extraArgs {
				sw.Do(", "+toGolangSourceDataLiteral(arg), nil)
			}
			sw.Do(")", targs)
		}

		if isFatal {
			sw.Do("if e := ", nil)
			emitCall()
			sw.Do("; len(e) != 0 {\n", nil)
			sw.Do("errs = append(errs, e...)\n", nil)
			sw.Do("    return // fatal\n", nil)
			sw.Do("}\n", nil)
		} else {
			sw.Do("errs = append(errs, ", nil)
			emitCall()
			sw.Do("...)\n", nil)
		}
	}
	if insideNilCheck {
		sw.Do("}\n", nil)
	}
}

func toGolangSourceDataLiteral(value any) string {
	// For safety, be strict in what values we output to visited source, and ensure strings
	// are quoted.
	switch value.(type) {
	case uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, bool:
		return fmt.Sprintf("%v", value)
	case string:
		// If the incoming string was quoted, we still do it ourselves, JIC.
		str := value.(string)
		if s, err := strconv.Unquote(str); err == nil {
			str = s
		}
		return fmt.Sprintf("%q", str)
	}
	// TODO: check this during discovery and emit an error with more useful information
	klog.Fatalf("unsupported extraArg type: %T", value)
	return ""
}
