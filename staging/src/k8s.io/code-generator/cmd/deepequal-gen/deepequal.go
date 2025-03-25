/*
Copyright 2025 The Kubernetes Authors.

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
	"io"
	"unicode"

	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
)

func mkPkgNames(pkg string, names ...string) []types.Name {
	result := make([]types.Name, 0, len(names))
	for _, name := range names {
		result = append(result, types.Name{Package: pkg, Name: name})
	}
	return result
}

func mkSymbolArgs(c *generator.Context, names []types.Name) generator.Args {
	args := generator.Args{}
	for _, name := range names {
		args[name.Name] = c.Universe.Type(name)
	}
	return args
}

var (
	ptrPkg        = "k8s.io/utils/ptr"
	ptrPkgSymbols = mkPkgNames(ptrPkg, "Equal", "To")
)

// genDeepEqual produces a file with DeepEqual methods for types in a package
type genDeepEqual struct {
	generator.GoGenerator
	outPkg  string
	imports namer.ImportTracker
	context *generator.Context
}

// NewDeepEqualGenerator creates a new generator for DeepEqual methods
func NewDeepEqualGenerator(c *generator.Context, pkg *types.Package, outputFileName string) generator.Generator {
	return &genDeepEqual{
		GoGenerator: generator.GoGenerator{
			OutputFilename: outputFileName,
		},
		outPkg:  pkg.Path,
		imports: generator.NewImportTrackerForPackage(pkg.Path),
		context: c,
	}
}

// Namers returns the name system used by this generator
func (g *genDeepEqual) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outPkg, g.imports),
	}
}

// FIXME
// shouldGenerateDeepEqual determines if we should generate a DeepEqual method for a type.
func (g *genDeepEqual) shouldGenerateDeepEqual(t *types.Type) bool {
	// Check if the type's package is in the input packages
	for _, inputPkg := range g.context.Inputs {
		if t.Name.Package == inputPkg {
			return true
		}
	}
	return false
}

// FIXME
// hasDeepEqual checks if a type has a DeepEqual method (from another package).
func (g *genDeepEqual) hasDeepEqual(t *types.Type) bool {
	// If we're generating DeepEqual for this type, return true
	for t.Kind == types.Pointer {
		t = t.Elem
	}
	if g.shouldGenerateDeepEqual(t) {
		return true
	}
	// Otherwise check if it has a DeepEqual method
	return t.Methods["DeepEqual"] != nil
}

// Filter determines which types to process.
func (g *genDeepEqual) Filter(c *generator.Context, t *types.Type) bool {
	// Only exported types in this package are processed.
	return t.Name.Package == g.outPkg && unicode.IsUpper(rune(t.Name.Name[0]))
}

// Imports returns the imports needed by the generated code.
func (g *genDeepEqual) Imports(c *generator.Context) []string {
	return g.imports.ImportLines()
}

// GenerateType generates the DeepEqual method for a single type.
func (g *genDeepEqual) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	switch t.Kind {
	// Gengo handles type definitions inconsistently.
	case types.Alias, types.Struct:
		// Process this type
	default:
		return nil
	}

	sw := generator.NewSnippetWriter(w, c, "$", "$")

	targs := generator.Args{
		"type": t,
	}

	// Generate the main DeepEqual method. We have already filtered for types
	// in this package.
	sw.Do("// DeepEqual returns true if this object is equal to the other object.\n", targs)
	sw.Do("func (this $.type|raw$) DeepEqual(other *$.type|raw$) bool {\n", targs)
	sw.Do("    if other == nil { return false }\n\n", targs)

	g.emitForType(c, sw, t, "", "")

	sw.Do("\n", targs)
	sw.Do("    return true\n", targs)
	sw.Do("}\n\n", targs)

	return sw.Error()
}

func (g *genDeepEqual) emitForType(c *generator.Context, sw *generator.SnippetWriter, t *types.Type, field string, deref string) {
	targs := generator.Args{
		"ptr":   mkSymbolArgs(c, ptrPkgSymbols),
		"type":  t,
		"field": field,
	}
	if field == "" {
		targs["lhsDeref"] = ""
		targs["rhsDeref"] = "*"
	} else {
		targs["lhsDeref"] = deref
		targs["rhsDeref"] = deref
	}

	switch t.Kind {
	case types.Builtin:
		// For builtin types, use direct comparison.
		sw.Do("    if $.lhsDeref$this$.field$ != $.rhsDeref$other$.field$ { return false }\n", targs)
		return

	case types.Alias:
		// For type aliases, treat it like the underlying type.
		g.emitForType(c, sw, t.Underlying, field, "")
		return

	case types.Pointer:
		//FIXME: field
		g.emitForPointer(c, sw, t.Elem, field, "*") //FIXME: deref?
		return
		/*
			elemType := t.Elem
			if g.hasDeepEqual(elemType) {
				sw.Do("    return (*this).DeepEqual(*other)\n", targs)
			} else if elemType.IsComparable() {
				sw.Do("    return $.ptr.Equal|raw$(this, other)\n", targs)
			} else {
				// For non-comparable types without DeepEqual, use helper
				helperName := g.getTypeHelperName(elemType)
				sw.Do("    return $.HelperName$(this, other)\n", map[string]interface{}{"HelperName": helperName})
			}
		*/
	case types.Struct:
		// Compare each exported field.
		for _, m := range t.Members {
			if !unicode.IsUpper(rune(m.Name[0])) {
				continue
			}
			g.emitForField(c, sw, m)
			/*
				g.generateFieldComparison(c, sw, m, "other."+m.Name)
			*/
		}

		/*
			case types.Map:
				// For maps, compare lengths and each key-value pair
				sw.Do("    if len(*this) != len(*other) { return false }\n", targs)
				sw.Do("    for k, v := range *this {\n", targs)
				sw.Do("        otherV, exists := (*other)[k]\n", targs)
				sw.Do("        if !exists { return false }\n", targs)
				if t.Elem.IsComparable() {
					sw.Do("        if v != otherV { return false }\n", targs)
				} else if t.Elem.Kind == types.Pointer {
					if g.hasDeepEqual(t.Elem.Elem) {
						sw.Do("        if !v.DeepEqual(otherV) { return false }\n", targs)
					} else if t.Elem.Elem.IsComparable() {
						sw.Do("        if !$.ptr.Equal|raw$(v, otherV) { return false }\n", targs)
					} else {
						helperName := g.getTypeHelperName(t.Elem.Elem)
						sw.Do("        if !$.HelperName$(v, otherV) { return false }\n", map[string]interface{}{"HelperName": helperName})
					}
				} else if g.needsTypeHelper(t.Elem) {
					helperName := g.getTypeHelperName(t.Elem)
					sw.Do("        if !$.HelperName$(&v, &otherV) { return false }\n", map[string]interface{}{"HelperName": helperName})
				}
				sw.Do("    }\n", targs)

			case types.Slice, types.Array:
				// For slices and arrays, compare lengths and each element
				sw.Do("    if len(*this) != len(*other) { return false }\n", targs)
				sw.Do("    for i := range *this {\n", targs)
				if t.Elem.IsComparable() {
					sw.Do("        if (*this)[i] != (*other)[i] { return false }\n", targs)
				} else if t.Elem.Kind == types.Pointer {
					if g.hasDeepEqual(t.Elem.Elem) {
						sw.Do("        if !(*this)[i].DeepEqual((*other)[i]) { return false }\n", targs)
					} else if t.Elem.Elem.IsComparable() {
						sw.Do("        if !$.ptr.Equal|raw$((*this)[i], (*other)[i]) { return false }\n", targs)
					} else {
						helperName := g.getTypeHelperName(t.Elem.Elem)
						sw.Do("        if !$.HelperName$((*this)[i], (*other)[i]) { return false }\n", map[string]interface{}{"HelperName": helperName})
					}
				} else if g.needsTypeHelper(t.Elem) {
					helperName := g.getTypeHelperName(t.Elem)
					sw.Do("        if !$.HelperName$(&(*this)[i], &(*other)[i]) { return false }\n", map[string]interface{}{"HelperName": helperName})
				}
				sw.Do("    }\n", targs)

		*/
	}
}

func (g *genDeepEqual) emitForPointerField(c *generator.Context, sw *generator.SnippetWriter, m types.Member) {
	targs := generator.Args{
		"field": m.Name,
	}
	sw.Do("if (this$.field$ == nil || other$.field$ == nil) && this$.field$ != other$.field$ {\n", targs)
	sw.Do("    return false\n", targs)
	sw.Do("} else if this$.field$ != nil {\n", targs)
	g.emitForField(c, sw, m) //FIXME: deref?
	sw.Do("}\n", targs)
}

func (g *genDeepEqual) emitForField(c *generator.Context, sw *generator.SnippetWriter, m types.Member) {
	targs := generator.Args{
		"ptr":   mkSymbolArgs(c, ptrPkgSymbols),
		"field": m.Name,
	}

	// If the type has a DeepEqual method, use it
	if g.hasDeepEqual(m.Type) {
		//FIXME: alias type
		if m.Type.Kind == types.Pointer {
			sw.Do("if !this.$.field$.DeepEqual(other.$.field$) { return false }\n", targs)
		} else {
			sw.Do("if !this.$.field$.DeepEqual(&other.$.field$) { return false }\n", targs)
		}
		return
	}
	g.emitForType(c, sw, m.Type, "."+m.Name, "")

	/*
		// For pointer type, we need to handle nil checks and element comparison
		if m.Type.Kind == types.Pointer {
			elemType := m.Type.Elem
			if g.hasDeepEqual(elemType) {
				sw.Do("    if !this.$.Name$.DeepEqual($.otherField$) { return false }\n", targs)
			} else if elemType.IsComparable() {
				sw.Do("    if !$.ptr.Equal|raw$(this.$.Name$, $.otherField$) { return false }\n", targs)
			} else {
				// For non-comparable types without DeepEqual, use helper
				helperName := g.getTypeHelperName(elemType)
				sw.Do("    if !$.HelperName$(this.$.Name$, $.otherField$) { return false }\n",
					map[string]interface{}{
						"HelperName": helperName,
						"Name":       m.Name,
						"otherField": otherField,
					})
			}
			return
		}

		if m.Type.IsComparable() {
			// For comparable types, use direct comparison
			sw.Do("    if this.$.Name$ != $.otherField$ { return false }\n", targs)
			return
		}

		switch m.Type.Kind {
		case types.Map:
			// For maps, compare lengths and each key-value pair
			sw.Do("    if len(this.$.Name$) != len($.otherField$) { return false }\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			if m.Type.Elem.IsComparable() {
				sw.Do("    for k, v := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        otherV, exists := $.otherField$[k]\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        if !exists { return false }\n", nil)
				sw.Do("        if v != otherV { return false }\n", nil)
				sw.Do("    }\n", nil)
			} else if m.Type.Elem.Kind == types.Pointer {
				sw.Do("    for k, v := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        otherV, exists := $.otherField$[k]\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        if !exists { return false }\n", nil)
				if g.hasDeepEqual(m.Type.Elem.Elem) {
					sw.Do("        if !v.DeepEqual(otherV) { return false }\n", nil)
				} else if m.Type.Elem.Elem.IsComparable() {
					sw.Do("        if !$.ptr.Equal|raw$(v, otherV) { return false }\n", targs)
				} else {
					helperName := g.getTypeHelperName(m.Type.Elem.Elem)
					sw.Do("        if !$.HelperName$(v, otherV) { return false }\n",
						map[string]interface{}{"HelperName": helperName})
				}
				sw.Do("    }\n", nil)
			} else if g.needsTypeHelper(m.Type.Elem) {
				helperName := g.getTypeHelperName(m.Type.Elem)
				sw.Do("    for k, v := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        otherV, exists := $.otherField$[k]\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        if !exists { return false }\n", nil)
				sw.Do("        if !$.HelperName$(&v, &otherV) { return false }\n",
					map[string]interface{}{"HelperName": helperName})
				sw.Do("    }\n", nil)
			}

		case types.Slice, types.Array:
			// For slices and arrays, compare lengths and each element
			sw.Do("    if len(this.$.Name$) != len($.otherField$) { return false }\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			if m.Type.Elem.IsComparable() {
				sw.Do("    for i := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        if this.$.Name$[i] != $.otherField$[i] { return false }\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("    }\n", nil)
			} else if m.Type.Elem.Kind == types.Pointer {
				sw.Do("    for i := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				if g.hasDeepEqual(m.Type.Elem.Elem) {
					sw.Do("        if !this.$.Name$[i].DeepEqual($.otherField$[i]) { return false }\n",
						map[string]interface{}{"Name": m.Name, "otherField": otherField})
				} else if m.Type.Elem.Elem.IsComparable() {
					sw.Do("        if !$.ptr.Equal|raw$(this.$.Name$[i], $.otherField$[i]) { return false }\n",
						map[string]interface{}{
							"ptr":        mkSymbolArgs(c, ptrPkgSymbols),
							"Name":       m.Name,
							"otherField": otherField,
						})
				} else {
					helperName := g.getTypeHelperName(m.Type.Elem.Elem)
					sw.Do("        if !$.HelperName$(this.$.Name$[i], $.otherField$[i]) { return false }\n",
						map[string]interface{}{"HelperName": helperName})
				}
				sw.Do("    }\n", nil)
			} else if g.needsTypeHelper(m.Type.Elem) {
				helperName := g.getTypeHelperName(m.Type.Elem)
				sw.Do("    for i := range this.$.Name$ {\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
				sw.Do("        if !$.HelperName$(&this.$.Name$[i], &$.otherField$[i]) { return false }\n",
					map[string]interface{}{
						"HelperName": helperName,
						"Name":       m.Name,
						"otherField": otherField,
					})
				sw.Do("    }\n", nil)
			}

		case types.Struct:
			// For structs, use their DeepEqual method
			if g.hasDeepEqual(m.Type) {
				sw.Do("    if !this.$.Name$.DeepEqual(&$.otherField$) { return false }\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
			} else {
				helperName := g.getTypeHelperName(m.Type)
				sw.Do("    if !$.HelperName$(&this.$.Name$, &$.otherField$) { return false }\n",
					map[string]interface{}{
						"HelperName": helperName,
						"Name":       m.Name,
						"otherField": otherField,
					})
			}

		case types.Alias:
			// For type aliases, use the underlying type's comparison method
			if m.Type.Underlying.IsComparable() {
				sw.Do("    if this.$.Name$ != $.otherField$ { return false }\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
			} else {
				helperName := g.getTypeHelperName(m.Type)
				sw.Do("    if !$.HelperName$(&this.$.Name$, &$.otherField$) { return false }\n",
					map[string]interface{}{
						"HelperName": helperName,
						"Name":       m.Name,
						"otherField": otherField,
					})
			}
		}
	*/
}

/*
func (g *genDeepEqual) needsTypeHelper(t *types.Type) bool {
	//FIXME: only generate helpers for types in other packages
	// Only generate helpers for base Alias and Struct types
	return t.Kind == types.Alias || t.Kind == types.Struct
}

func (g *genDeepEqual) getTypeHelperName(t *types.Type) string {
	// For pointers, use the element type's name
	if t.Kind == types.Pointer {
		return "deepEqual" + t.Elem.Name.Name
	}
	return "deepEqual" + t.Name.Name
}

// generateTypeHelper generates a helper function for comparing complex types
func (g *genDeepEqual) generateTypeHelper(c *generator.Context, sw *generator.SnippetWriter, t *types.Type) {
	// For pointers, generate helper for the element type
	if t.Kind == types.Pointer {
		t = t.Elem
	}

	helperName := g.getTypeHelperName(t)

	targs := generator.Args{
		"ptr":        mkSymbolArgs(c, ptrPkgSymbols),
		"HelperName": helperName,
		"Type":       t,
	}

	// Generate one helper that takes pointers
	sw.Do("func $.HelperName$(this, other *$.Type|raw$) bool {\n", targs)
	sw.Do("    if this == other { return true }\n", targs)
	sw.Do("    if this == nil || other == nil { return false }\n", targs)

	if g.hasDeepEqual(t) {
		sw.Do("    return this.DeepEqual(other)\n", targs)
	} else if t.IsComparable() {
		sw.Do("    return *this == *other\n", targs)
	} else {
		switch t.Kind {
		case types.Struct:
			for _, m := range t.Members {
				if !unicode.IsUpper(rune(m.Name[0])) {
					continue
				}
				field := types.Member{
					Name: m.Name,
					Type: m.Type,
				}
				g.generateFieldComparison(c, sw, field, "other."+m.Name)
			}
			sw.Do("    return true\n", targs)
		case types.Alias:
			// For aliases, compare using the underlying type
			if t.Underlying.IsComparable() {
				sw.Do("    return *this == *other\n", targs)
			} else {
				g.generateFieldComparison(c, sw, types.Member{Name: "*this", Type: t.Underlying}, "*other")
				sw.Do("    return true\n", targs)
			}
		}
	}
	sw.Do("}\n\n", targs)
}

func (g *genDeepEqual) generateTypeComparison(c *generator.Context, sw *generator.SnippetWriter, t *types.Type, lhs, rhs string, isPtr bool) {
	targs := generator.Args{
		"ptr": mkSymbolArgs(c, ptrPkgSymbols),
		"lhs": lhs,
		"rhs": rhs,
	}

	if t.Methods["DeepEqual"] != nil {
		if isPtr {
			sw.Do("    return $.lhs$.DeepEqual($.rhs$)\n", targs)
		} else {
			sw.Do("    return (&$.lhs$).DeepEqual(&$.rhs$)\n", targs)
		}
		return
	}

	if t.IsComparable() {
		sw.Do("    return $.lhs$ == $.rhs$\n", targs)
		return
	}

	switch t.Kind {
	case types.Struct:
		for _, m := range t.Members {
			if !unicode.IsUpper(rune(m.Name[0])) {
				continue
			}
			field := types.Member{
				Name: lhs + "." + m.Name,
				Type: m.Type,
			}
			g.generateFieldComparison(c, sw, field, rhs)
		}

	case types.Alias:
		// For aliases, compare using the underlying type
		g.generateTypeComparison(c, sw, t.Underlying, lhs, rhs, isPtr)
	}
}

// generateFieldComparison generates comparison code for a struct field
func (g *genDeepEqual) generateFieldComparison(c *generator.Context, sw *generator.SnippetWriter, m types.Member, otherField string) {
	targs := generator.Args{
		"ptr":        mkSymbolArgs(c, ptrPkgSymbols),
		"Name":       m.Name,
		"otherField": otherField,
	}

	// If the type has a DeepEqual method, use it
	if g.hasDeepEqual(m.Type) {
		if m.Type.Kind == types.Pointer {
			sw.Do("    if !this.$.Name$.DeepEqual($.otherField$) { return false }\n", targs)
		} else {
			sw.Do("    if !this.$.Name$.DeepEqual(&$.otherField$) { return false }\n", targs)
		}
		return
	}

	// For pointer type, we need to handle nil checks and element comparison
	if m.Type.Kind == types.Pointer {
		elemType := m.Type.Elem
		if g.hasDeepEqual(elemType) {
			sw.Do("    if !this.$.Name$.DeepEqual($.otherField$) { return false }\n", targs)
		} else if elemType.IsComparable() {
			sw.Do("    if !$.ptr.Equal|raw$(this.$.Name$, $.otherField$) { return false }\n", targs)
		} else {
			// For non-comparable types without DeepEqual, use helper
			helperName := g.getTypeHelperName(elemType)
			sw.Do("    if !$.HelperName$(this.$.Name$, $.otherField$) { return false }\n",
				map[string]interface{}{
					"HelperName": helperName,
					"Name":       m.Name,
					"otherField": otherField,
				})
		}
		return
	}

	if m.Type.IsComparable() {
		// For comparable types, use direct comparison
		sw.Do("    if this.$.Name$ != $.otherField$ { return false }\n", targs)
		return
	}

	switch m.Type.Kind {
	case types.Map:
		// For maps, compare lengths and each key-value pair
		sw.Do("    if len(this.$.Name$) != len($.otherField$) { return false }\n",
			map[string]interface{}{"Name": m.Name, "otherField": otherField})
		if m.Type.Elem.IsComparable() {
			sw.Do("    for k, v := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        otherV, exists := $.otherField$[k]\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        if !exists { return false }\n", nil)
			sw.Do("        if v != otherV { return false }\n", nil)
			sw.Do("    }\n", nil)
		} else if m.Type.Elem.Kind == types.Pointer {
			sw.Do("    for k, v := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        otherV, exists := $.otherField$[k]\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        if !exists { return false }\n", nil)
			if g.hasDeepEqual(m.Type.Elem.Elem) {
				sw.Do("        if !v.DeepEqual(otherV) { return false }\n", nil)
			} else if m.Type.Elem.Elem.IsComparable() {
				sw.Do("        if !$.ptr.Equal|raw$(v, otherV) { return false }\n", targs)
			} else {
				helperName := g.getTypeHelperName(m.Type.Elem.Elem)
				sw.Do("        if !$.HelperName$(v, otherV) { return false }\n",
					map[string]interface{}{"HelperName": helperName})
			}
			sw.Do("    }\n", nil)
		} else if g.needsTypeHelper(m.Type.Elem) {
			helperName := g.getTypeHelperName(m.Type.Elem)
			sw.Do("    for k, v := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        otherV, exists := $.otherField$[k]\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        if !exists { return false }\n", nil)
			sw.Do("        if !$.HelperName$(&v, &otherV) { return false }\n",
				map[string]interface{}{"HelperName": helperName})
			sw.Do("    }\n", nil)
		}

	case types.Slice, types.Array:
		// For slices and arrays, compare lengths and each element
		sw.Do("    if len(this.$.Name$) != len($.otherField$) { return false }\n",
			map[string]interface{}{"Name": m.Name, "otherField": otherField})
		if m.Type.Elem.IsComparable() {
			sw.Do("    for i := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        if this.$.Name$[i] != $.otherField$[i] { return false }\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("    }\n", nil)
		} else if m.Type.Elem.Kind == types.Pointer {
			sw.Do("    for i := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			if g.hasDeepEqual(m.Type.Elem.Elem) {
				sw.Do("        if !this.$.Name$[i].DeepEqual($.otherField$[i]) { return false }\n",
					map[string]interface{}{"Name": m.Name, "otherField": otherField})
			} else if m.Type.Elem.Elem.IsComparable() {
				sw.Do("        if !$.ptr.Equal|raw$(this.$.Name$[i], $.otherField$[i]) { return false }\n",
					map[string]interface{}{
						"ptr":        mkSymbolArgs(c, ptrPkgSymbols),
						"Name":       m.Name,
						"otherField": otherField,
					})
			} else {
				helperName := g.getTypeHelperName(m.Type.Elem.Elem)
				sw.Do("        if !$.HelperName$(this.$.Name$[i], $.otherField$[i]) { return false }\n",
					map[string]interface{}{"HelperName": helperName})
			}
			sw.Do("    }\n", nil)
		} else if g.needsTypeHelper(m.Type.Elem) {
			helperName := g.getTypeHelperName(m.Type.Elem)
			sw.Do("    for i := range this.$.Name$ {\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
			sw.Do("        if !$.HelperName$(&this.$.Name$[i], &$.otherField$[i]) { return false }\n",
				map[string]interface{}{
					"HelperName": helperName,
					"Name":       m.Name,
					"otherField": otherField,
				})
			sw.Do("    }\n", nil)
		}

	case types.Struct:
		// For structs, use their DeepEqual method
		if g.hasDeepEqual(m.Type) {
			sw.Do("    if !this.$.Name$.DeepEqual(&$.otherField$) { return false }\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
		} else {
			helperName := g.getTypeHelperName(m.Type)
			sw.Do("    if !$.HelperName$(&this.$.Name$, &$.otherField$) { return false }\n",
				map[string]interface{}{
					"HelperName": helperName,
					"Name":       m.Name,
					"otherField": otherField,
				})
		}

	case types.Alias:
		// For type aliases, use the underlying type's comparison method
		if m.Type.Underlying.IsComparable() {
			sw.Do("    if this.$.Name$ != $.otherField$ { return false }\n",
				map[string]interface{}{"Name": m.Name, "otherField": otherField})
		} else {
			helperName := g.getTypeHelperName(m.Type)
			sw.Do("    if !$.HelperName$(&this.$.Name$, &$.otherField$) { return false }\n",
				map[string]interface{}{
					"HelperName": helperName,
					"Name":       m.Name,
					"otherField": otherField,
				})
		}
	}
}
*/
