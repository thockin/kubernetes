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
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

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
	typesPackage  string
	outputPackage string
	peerPackages  []string
	inputToPkg    map[string]string // Maps input packages to generated validation packages
	initTypes     []*types.Type
	typeNodes     map[*types.Type]*typeNode
	imports       namer.ImportTracker
	validator     validators.DeclarativeValidator
}

// NewGenValidations cretes a new generator for the specified package.
func NewGenValidations(outputFilename, typesPackage, outputPackage string, initTypes []*types.Type, typeNodes map[*types.Type]*typeNode, peerPkgs []string, inputToPkg map[string]string, validator validators.DeclarativeValidator) generator.Generator {
	return &genValidations{
		GoGenerator: generator.GoGenerator{
			OutputFilename: outputFilename,
		},
		typesPackage:  typesPackage,
		outputPackage: outputPackage,
		peerPackages:  peerPkgs,
		inputToPkg:    inputToPkg,
		initTypes:     initTypes,
		typeNodes:     typeNodes,
		imports:       generator.NewImportTrackerForPackage(outputPackage),
		validator:     validator,
	}
}

func (g *genValidations) Namers(_ *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *genValidations) Filter(_ *generator.Context, t *types.Type) bool {
	_, ok := g.typeNodes[t]
	return ok
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
	for _, t := range g.initTypes {
		tn, ok := g.typeNodes[t]
		if !ok {
			continue
		}
		if tn == nil {
			// Should never happen.
			klog.Fatalf("found nil typeNode for type %v", t)
		}

		// TODO: It would be nice if these were not hard-coded.
		specType := t
		var specField string
		if spec := tn.lookupField("spec"); spec != nil {
			specType = spec.underlyingType
			specField = spec.name
		}
		var statusType *types.Type
		var statusField string
		if status := tn.lookupField("status"); status != nil {
			statusType = status.underlyingType
			statusField = status.name
		}

		targs := generator.Args{
			"rootType":    t,
			"specType":    specType,
			"specField":   specField,
			"statusType":  statusType,
			"statusField": statusField,
			"errorList":   c.Universe.Type(errorListType),
			"fieldPath":   c.Universe.Type(fieldPathType),
			"fmtErrorf":   c.Universe.Type(errorfType),
		}
		//TODO: can this be (*$.rootType|raw$)(nil) ?
		sw.Do("scheme.AddValidationFunc(new($.rootType|raw$), func(obj, oldObj interface{}, subresources ...string) $.errorList|raw$ {\n", targs)
		sw.Do("  if len(subresources) == 0 {\n", targs)
		if specType != t {
			sw.Do("    root := obj.(*$.rootType|raw$)\n", targs)
			sw.Do("    return $.specType|objectvalidationfn$(&root.$.specField$, nil)\n", targs)
		} else {
			sw.Do("    return $.specType|objectvalidationfn$(obj.(*$.rootType|raw$), nil)\n", targs)
		}
		sw.Do("  }\n", targs)

		if statusType != nil {
			sw.Do("  if len(subresources) == 1 && subresources[0] == \"status\" {\n", targs)
			sw.Do("    root := obj.(*$.rootType|raw$)\n", targs)
			sw.Do("    return $.statusType|objectvalidationfn$(&root.$.statusField$, nil)\n", targs)
			sw.Do("  }\n", targs)
		}
		sw.Do("  return $.errorList|raw${field.InternalError(nil, $.fmtErrorf|raw$(\"No validation found for %T, subresources: %v\", obj, subresources))}\n", targs)
		sw.Do("})\n", targs)

		// TODO: Support update validations
		//       This will require correlating old object.
	}
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

func (g *genValidations) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	klog.V(5).Infof("generating for type %v", t)

	var errs []error

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	g.emitValidationFunction(c, t, sw)
	if err := sw.Error(); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func (g *genValidations) emitValidationFunction(c *generator.Context, t *types.Type, sw *generator.SnippetWriter) {
	targs := generator.Args{
		"inType":    t,
		"errorList": c.Universe.Type(errorListType),
		"fieldPath": c.Universe.Type(fieldPathType),
	}

	sw.Do("func $.inType|objectvalidationfn$(in *$.inType|raw$, fldPath *$.fieldPath|raw$) (errs $.errorList|raw$) {\n", targs)
	g.emitValidationForType(c, t, "in", pathPart{}, sw, nil, nil)
	sw.Do("return errs\n", nil)
	sw.Do("}\n\n", nil)
}

// typeDiscoverer contains fields necessary to build a tree of types.
type typeDiscoverer struct {
	validator  validators.DeclarativeValidator
	inputToPkg map[string]string
	inProgress map[*types.Type]bool
	knownTypes map[*types.Type]*typeNode
}

// discoverTypes walks the type graph and populates the result map.
func discoverTypes(validator validators.DeclarativeValidator, inputToPkg map[string]string, t *types.Type, results map[*types.Type]*typeNode) error {
	td := &typeDiscoverer{
		validator:  validator,
		inputToPkg: inputToPkg,
		knownTypes: results,
	}
	return td.discover(t)
}

// typeNode carries validation informatiuon for a single type.
type typeNode struct {
	underlyingType *types.Type
	validations    []validators.FunctionGen
	children       []*childNode // populated when parent is a Struct
	elem           *childNode   // populated when parent is a list
	key            *childNode   // populated when parent is a map
	funcName       types.Name
}

func (n typeNode) lookupField(jsonName string) *childNode {
	for _, c := range n.children {
		if c.jsonName == jsonName {
			return c
		}
	}
	return nil
}

// childNode represents a field in a struct.
// FIXME: merge this back into typeNode
type childNode struct {
	name           string
	jsonName       string
	underlyingType *types.Type
	validations    []validators.FunctionGen

	// iterated validation has to be tracked separately from field's validations.
	eachKey, eachVal []validators.FunctionGen
}

const (
	eachKeyTag = "eachKey"
	eachValTag = "eachVal"
)

// discover walks the type graph, starting at t, and registers all types into
// knownTypes.  The specified comments represent the parent context for this
// type - the type comments for a type definition or the field comments for a
// field.
func (td *typeDiscoverer) discover(t *types.Type) error {
	// If we already know this type, we are done.
	if _, ok := td.knownTypes[t]; ok {
		return nil
	}

	klog.V(5).InfoS("discovering", "type", t)

	thisNode := &typeNode{
		underlyingType: t,
	}

	// Publish it right away in case we hit it recursively.
	//FIXME: don't store Builtins?
	td.knownTypes[t] = thisNode

	// Extract any type-attached validation rules.
	if validations, err := td.validator.ExtractValidations(t.Name.Name, t, t.CommentLines); err != nil {
		return err
	} else {
		if len(validations) > 0 {
			klog.V(5).InfoS("  found type-attached validations", "n", len(validations))
			thisNode.validations = validations
		}
	}

	// FIXME: don't descend into other pkgs - add test for cross-pkgs
	switch t.Kind {
	case types.Builtin:
		//FIXME: nothing to do?  return early?
	case types.Pointer:
		klog.V(5).InfoS("  type is a pointer", "type", t.Elem)
		if t.Elem.Kind == types.Pointer {
			klog.Fatalf("type %v: pointers to pointers are not supported", t)
		}
		if err := td.discover(t.Elem); err != nil {
			return err
		}
	case types.Slice, types.Array:
		klog.V(5).InfoS("  type is a list", "type", t.Elem)
		if err := td.discover(t.Elem); err != nil {
			return err
		}
		thisNode.elem = &childNode{
			underlyingType: t.Elem,
		}
	case types.Map:
		klog.V(5).InfoS("  type is a map", "type", t.Elem)
		if err := td.discover(t.Key); err != nil {
			return err
		}
		thisNode.key = &childNode{
			underlyingType: t.Elem,
		}

		if err := td.discover(t.Elem); err != nil {
			return err
		}
		thisNode.elem = &childNode{
			underlyingType: t.Elem,
		}
	case types.Struct:
		klog.V(5).InfoS("  type is a struct")
		fn, ok := td.getValidationFunctionName(t)
		if !ok {
			//FIXME: this seems like an error, but is it?  Or just "opaque from here"
			return nil
		}
		thisNode.funcName = fn

		for _, field := range t.Members {
			name := field.Name
			if len(name) == 0 {
				// embedded fields
				if field.Type.Kind == types.Pointer {
					name = field.Type.Elem.Name.Name
				} else {
					name = field.Type.Name.Name
				}
			}
			jsonName := "<unknown-json-name>"
			if tags, ok := lookupJSONTags(field); ok {
				jsonName = tags.name
			}
			//FIXME: only do exported fields, add a test
			klog.V(5).InfoS("  field", "name", name)

			if err := td.discover(field.Type); err != nil {
				return err
			}

			child := &childNode{
				name:           name,
				jsonName:       jsonName,
				underlyingType: field.Type,
			}

			switch field.Type.Kind {
			case types.Map:
				//TODO: also support +k8s:eachKey
				if tagVals, found := gengo.ExtractCommentTags("+", field.CommentLines)[eachKeyTag]; found {
					for _, tagVal := range tagVals {
						fakeComments := []string{tagVal}
						// Extract any embedded key-validation rules.
						if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[keys]", field.Name), field.Type.Key, fakeComments); err != nil {
							return err
						} else {
							if len(validations) > 0 {
								klog.V(5).InfoS("  found key-validations", "n", len(validations))
								child.eachKey = append(child.eachKey, validations...)
							}
						}
					}
				}
				//TODO: also support +k8s:eachVal
				if tagVals, found := gengo.ExtractCommentTags("+", field.CommentLines)[eachValTag]; found {
					for _, tagVal := range tagVals {
						fakeComments := []string{tagVal}
						// Extract any embedded list-validation rules.
						if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[vals]", field.Name), field.Type.Elem, fakeComments); err != nil {
							return err
						} else {
							if len(validations) > 0 {
								klog.V(5).InfoS("  found list-validations", "n", len(validations))
								child.eachVal = append(child.eachVal, validations...)
							}
						}
					}
				}
			case types.Slice, types.Array:
				//TODO: also support +k8s:eachVal
				if tagVals, found := gengo.ExtractCommentTags("+", field.CommentLines)[eachValTag]; found {
					for _, tagVal := range tagVals {
						fakeComments := []string{tagVal}
						// Extract any embedded list-validation rules.
						if validations, err := td.validator.ExtractValidations(fmt.Sprintf("%s[vals]", field.Name), field.Type.Elem, fakeComments); err != nil {
							return err
						} else {
							if len(validations) > 0 {
								klog.V(5).InfoS("  found list-validations", "n", len(validations))
								child.eachVal = append(child.eachVal, validations...)
							}
						}
					}
				}
			}

			// Extract any field-attached validation rules.
			if validations, err := td.validator.ExtractValidations(name, field.Type, field.CommentLines); err != nil {
				return err
			} else {
				if len(validations) > 0 {
					klog.V(5).InfoS("  found field-attached value-validations", "n", len(validations))
					child.validations = append(child.validations, validations...)
				}
			}
			thisNode.children = append(thisNode.children, child)
		}
	case types.Alias:
		klog.V(5).InfoS("  type is an alias", "type", t.Underlying)
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
			klog.Fatalf("type %v: aliases to pointers are not supported", t)
		}
		if err := td.discover(t.Underlying); err != nil {
			return err
		}
		fn, ok := td.getValidationFunctionName(t)
		if !ok {
			//FIXME: this seems like an error, but is it?  Or just "opaque from here"
			return nil
		}
		thisNode.funcName = fn
	default:
		klog.Fatalf("unhandled type: %v (%v)", t, t.Kind)
	}

	return nil
}

func (td *typeDiscoverer) getValidationFunctionName(t *types.Type) (types.Name, bool) {
	pkg, ok := td.inputToPkg[t.Name.Package]
	if !ok {
		return types.Name{}, false
	}
	return types.Name{Package: pkg, Name: "Validate_" + t.Name.Name}, true
}

// emitValidationForType writes code for inType, calling type-attached
// validations and then descending into the type (e.g. fstruct fields).
// inType is always a value type, with pointerness removed, and varName
// accomodates for that.
func (g *genValidations) emitValidationForType(c *generator.Context, inType *types.Type, varName string, path pathPart, sw *generator.SnippetWriter, eachKey, eachVal []validators.FunctionGen) {
	if inType.Kind == types.Pointer {
		klog.Fatalf("unexpected pointer: %v (%s)", inType, varName)
	}

	// Emit code for type-attached validations.
	tn := g.typeNodes[inType]
	if len(tn.validations) > 0 {
		g.emitCallsToValidators(c, tn.validations, varName, path, false, sw)
	}

	targs := generator.Args{
		"var": varName,
	}

	// Descend into the type.
	switch inType.Kind {
	case types.Builtin:
		// Nothing further.
	case types.Alias:
		// Nothing further.
	case types.Struct:
		for _, child := range tn.children {
			childVarName := varName
			if len(child.name) > 0 {
				childVarName = varName + "." + child.name
			} else {
				klog.Fatalf("missing child name for field in %v", inType)
			}
			childPath := pathPart{Child: child.jsonName}
			childIsPtr := child.underlyingType.Kind == types.Pointer
			//FIXME: sometimes this need a newline before it and sometimes not?
			sw.Do("// $.fieldName$\n", generator.Args{"fieldName": child.name})
			if len(child.validations) > 0 {
				// When calling registered validators, we always pass the
				// underlying value-type.  E.g. if the field's type is string,
				// we pass string, and if the field's type is *string, we also
				// pass string (checking for nil, first).  This means those
				// validators don't have to know the difference, but it also
				// means that large structs will be passed by value.  If this
				// turns out to be a real problem, we could change this to pass
				// everything by pointer.
				g.emitCallsToValidators(c, child.validations, childVarName, childPath, childIsPtr, sw)
			}

			// Get to the real type.
			t := child.underlyingType
			if t.Kind == types.Pointer {
				t = t.Elem
			}

			if t.Kind == types.Struct || t.Kind == types.Alias {
				// If this field is another type, call its validation function.
				// Checking for nil is handled inside this call.
				g.emitCallToOtherTypeFunc(c, t, childVarName, childPath, childIsPtr, sw)
			} else {
				// Descend into this field.
				g.emitValidationForType(c, t, childVarName, childPath, sw, child.eachKey, child.eachVal)
			}
			sw.Do("\n", nil)
		}
	case types.Slice, types.Array:
		// TODO: get rid of tn.elem and tn.key - redundant with underlyingType.Elem/Key
		child := tn.elem
		if tn.elem.underlyingType != inType.Elem {
			panic("oops")
		}
		elemPath := pathPart{Index: "i"}
		elemIsPtr := inType.Elem.Kind == types.Pointer

		//FIXME: figure out if we can make this a wrapper-function and do it in one call to validate.ValuesInSlice()
		sw.Do("for i, val := range $.var$ {\n", targs)

		// Validate each value.
		validations := child.validations
		validations = append(validations, eachVal...)
		if len(validations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			g.emitCallsToValidators(c, validations, "val", elemPath, elemIsPtr, sw)
		}

		// Get to the real type.
		t := inType.Elem
		if t.Kind == types.Pointer {
			t = t.Elem
		}

		if t.Kind == types.Struct || t.Kind == types.Alias {
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, t, "val", elemPath, elemIsPtr, sw)
		} else {
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		sw.Do("}\n", nil)
	case types.Map:
		keyChild := tn.key
		keyPath := pathPart{} // TODO: we need a way to denote "invalid key"
		keyIsPtr := inType.Key.Kind == types.Pointer

		valChild := tn.elem
		valPath := pathPart{Key: "key"}
		valIsPtr := inType.Elem.Kind == types.Pointer

		sw.Do("for key, val := range $.var$ {\n", targs)

		// Validate each key.
		keyValidations := keyChild.validations
		keyValidations = append(keyValidations, eachKey...)
		if len(keyValidations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			g.emitCallsToValidators(c, keyValidations, "key", keyPath, keyIsPtr, sw)
		}

		// Get to the real type.
		t := inType.Key
		if t.Kind == types.Pointer {
			t = t.Elem
		}

		if t.Kind == types.Struct || t.Kind == types.Alias {
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, t, "key", keyPath, keyIsPtr, sw)
		} else {
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		// Validate each value.
		valValidations := valChild.validations
		valValidations = append(valValidations, eachVal...)
		if len(valValidations) > 0 {
			// When calling registered validators, we always pass the
			// underlying value-type.  E.g. if the field's type is string,
			// we pass string, and if the field's type is *string, we also
			// pass string (checking for nil, first).  This means those
			// validators don't have to know the difference, but it also
			// means that large structs will be passed by value.  If this
			// turns out to be a real problem, we could change this to pass
			// everything by pointer.
			g.emitCallsToValidators(c, valValidations, "val", valPath, valIsPtr, sw)
		}

		// Get to the real type.
		t = inType.Elem
		if t.Kind == types.Pointer {
			t = t.Elem
		}

		if t.Kind == types.Struct || t.Kind == types.Alias {
			// If this field is another type, call its validation function.
			// Checking for nil is handled inside this call.
			g.emitCallToOtherTypeFunc(c, t, "val", valPath, valIsPtr, sw)
		} else {
			// No need to go further.  Struct- or alias-typed fields might have
			// validations attached to the type, but anything else (e.g.
			// string) can't, and we already emitted code for the field
			// validations.
		}

		sw.Do("}\n", nil)
	default:
		klog.Fatalf("unhandled type: %v (%s)", inType, inType.Kind)
	}
}

// emitCallToOtherTypeFunc generates a call to a different generated validation
// function for a field in some parent context.  inType is the value type
// being validated with pointerness removed.  isVarPtr indicates that the value
// was a pointer in the parent context.  varName is the name of this value in
// the parent context, and path is used to build the field.Path for the call.
func (g *genValidations) emitCallToOtherTypeFunc(c *generator.Context, inType *types.Type, varName string, path pathPart, isVarPtr bool, sw *generator.SnippetWriter) {
	if isVarPtr {
		sw.Do("if $.var$ != nil {\n", generator.Args{"var": varName})
		defer func() {
			sw.Do("}\n", nil)
		}()
	} else {
		varName = "&" + varName
	}

	tn := g.typeNodes[inType]
	targs := generator.Args{
		"var":  varName,
		"path": path,
		"fn":   c.Universe.Type(tn.funcName),
	}
	sw.Do("errs = append(errs, $.fn|raw$($.var$, ", targs)
	if len(path.Child) > 0 {
		sw.Do("fldPath.Child(\"$.path.Child$\")", targs)
	} else if len(path.Index) > 0 {
		sw.Do("fldPath.Index($.path.Index$)", targs)
	} else if len(path.Key) > 0 {
		sw.Do("fldPath.Key($.path.Key$)", targs)
	} else {
		sw.Do("fldPath", targs)
	}
	sw.Do(")...)\n", targs)
}

// emitCallsToValidators generates calls to a list of validation functions for
// a single field or type. validations is a list of functions to call, with
// arguments.  varName is the name of this value in the parent context, and
// path is used to build the field.Path for the call.  isVarPtr indicates that
// the value  was a pointer in the parent context.
func (g *genValidations) emitCallsToValidators(c *generator.Context, validations []validators.FunctionGen, varName string, path pathPart, isVarPtr bool, sw *generator.SnippetWriter) {
	nFatal := 0
	// TODO: We previously relied on Fatal validators also being priority, but
	// when we have forEach, the ordering can be messed up.  Instead, we should
	// sort by fatalness before running them, or do two passes.
	for i, v := range validations {
		moreValidations := i != len(validations)-1

		fn, extraArgs := v.SignatureAndArgs()
		targs := generator.Args{
			"var":  varName,
			"path": path,
			"fn":   c.Universe.Type(fn),
		}
		closeThisValidation := func() {}
		if isVarPtr && (v.Flags()&validators.PtrOK == 0) {
			// TODO: This test will be emitted for each validation. We could
			// restructure this to collect all of the calls, sort by PtrOK, and
			// emit this one time.
			sw.Do("if $.var$ != nil {\n", targs)
			closeThisValidation = func() {
				sw.Do("}\n", nil)
			}
			targs["var"] = "*" + varName
		}

		emitCall := func() {
			sw.Do("$.fn|raw$(", targs)
			if len(path.Child) > 0 {
				sw.Do("fldPath.Child(\"$.path.Child$\")", targs)
			} else if len(path.Index) > 0 {
				sw.Do("fldPath.Index($.path.Index$)", targs)
			} else if len(path.Key) > 0 {
				sw.Do("fldPath.Key($.path.Key$)", targs)
			} else {
				sw.Do("fldPath", targs)
			}
			sw.Do(", $.var$", targs)
			for _, arg := range extraArgs {
				sw.Do(", "+toGolangSourceDataLiteral(arg), nil)
			}
			sw.Do(")", targs)
		}
		if (v.Flags()&validators.IsFatal) != 0 && moreValidations {
			nFatal++
			sw.Do("if e := ", nil)
			emitCall()
			sw.Do("; len(e) != 0 {\n", nil)
			sw.Do("errs = append(errs, e...)\n", nil)
			sw.Do("} else {\n", nil)
		} else {
			sw.Do("errs = append(errs, ", nil)
			emitCall()
			sw.Do("...)\n", nil)
		}
		closeThisValidation()
	}
	for i := 0; i < nFatal; i++ {
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
	klog.Fatalf("Unsupported extraArg type: %T", value) // TODO: handle error
	return ""
}

type pathPart struct {
	Child string
	Index string
	Key   string
}
