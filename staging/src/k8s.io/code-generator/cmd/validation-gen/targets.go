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
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/code-generator/cmd/validation-gen/validators"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"
)

// These are the comment tags that carry parameters for validation generation.
const (
	tagName         = "k8s:validation-gen"
	inputTagName    = "k8s:validation-gen-input"
	enabledTagName  = "k8s:validation-gen-enabled-tags"
	disabledTagName = "k8s:validation-gen-disabled-tags"
)

func extractTag(comments []string) ([]string, bool) {
	m := gengo.ExtractCommentTags("+", comments)
	v, found := m[tagName]
	return v, found
}

func extractInputTag(comments []string) []string {
	return gengo.ExtractCommentTags("+", comments)[inputTagName]
}

func extractFiltersTags(comments []string) (enabled, disabled []string) {
	return gengo.ExtractCommentTags("+", comments)[enabledTagName],
		gengo.ExtractCommentTags("+", comments)[disabledTagName]
}

func checkTag(comments []string, require ...string) bool {
	values := gengo.ExtractCommentTags("+", comments)[tagName]
	if len(require) == 0 {
		return len(values) == 1 && values[0] == ""
	}
	return reflect.DeepEqual(values, require)
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":             namer.NewPublicNamer(1),
		"raw":                namer.NewRawNamer("", nil),
		"objectvalidationfn": validationFnNamer(),
	}
}

func validationFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix: "Validate_",
		Join: func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func GetTargets(context *generator.Context, args *Args) []generator.Target {
	boilerplate, err := gengo.GoBoilerplate(args.GoHeaderFile, gengo.StdBuildTag, gengo.StdGeneratedBy)
	if err != nil {
		klog.Fatalf("Failed loading boilerplate: %v", err)
	}

	var targets []generator.Target

	// First load other "input" packages.  We do this as a single call because
	// it is MUCH faster.
	inputPkgs := make([]string, 0, len(context.Inputs))
	pkgToInput := map[string]string{}
	for _, input := range context.Inputs {
		klog.V(5).Infof("considering pkg %q", input)

		pkg := context.Universe[input]

		// if the types are not in the same package where the validation
		// functions are to be emitted
		inputTags := extractInputTag(pkg.Comments)
		if len(inputTags) > 1 {
			panic(fmt.Sprintf("there may only be one input tag, got %#v", inputTags))
		}
		if len(inputTags) == 1 {
			inputPath := inputTags[0]
			if strings.HasPrefix(inputPath, "./") || strings.HasPrefix(inputPath, "../") {
				// this is a relative dir, which will not work under gomodules.
				// join with the local package path, but warn
				klog.Fatalf("relative path (%s=%s) is not supported; use full package path (as used by 'import') instead", inputTagName, inputPath)
			}

			klog.V(5).Infof("  input pkg %v", inputPath)
			inputPkgs = append(inputPkgs, inputPath)
			pkgToInput[input] = inputPath
		} else {
			pkgToInput[input] = input
		}
	}

	// Make sure explicit peer-packages are added.
	var peerPkgs []string
	for _, pkg := range args.ExtraPeerDirs {
		// In case someone specifies a peer as a path into vendor, convert
		// it to its "real" package path.
		if i := strings.Index(pkg, "/vendor/"); i != -1 {
			pkg = pkg[i+len("/vendor/"):]
		}
		peerPkgs = append(peerPkgs, pkg)
	}
	if expanded, err := context.FindPackages(peerPkgs...); err != nil {
		klog.Fatalf("cannot find peer packages: %v", err)
	} else {
		peerPkgs = expanded // now in fully canonical form
	}
	inputPkgs = append(inputPkgs, peerPkgs...)

	if len(inputPkgs) > 0 {
		if _, err := context.LoadPackages(inputPkgs...); err != nil {
			klog.Fatalf("cannot load packages: %v", err)
		}
	}
	// update context.Order to the latest context.Universe
	orderer := namer.Orderer{Namer: namer.NewPublicNamer(1)}
	context.Order = orderer.OrderUniverse(context.Universe)

	// We also need the to be able to look up the packages of inputs
	inputToPkg := make(map[string]string, len(pkgToInput))
	for k, v := range pkgToInput {
		inputToPkg[v] = k
	}

	// Build a cache of type->callNode for every type we need.
	typeNodes := map[*types.Type]*typeNode{}
	for _, input := range context.Inputs {
		klog.V(2).InfoS("processing", "pkg", input)

		pkg := context.Universe[input]

		enabledTags, disabledTags := extractFiltersTags(pkg.Comments)
		declarativeValidator := validators.NewValidator(context, enabledTags, disabledTags)

		typesWith, found := extractTag(pkg.Comments)
		if !found {
			klog.V(2).InfoS("  did not find required tag", "tag", tagName)
			continue
		}
		if len(typesWith) == 0 {
			klog.Fatalf("found package tag %q with no value", tagName)
		}
		shouldCreateObjectValidationFn := func(t *types.Type) bool {
			// opt-out
			if checkTag(t.SecondClosestCommentLines, "false") {
				return false
			}
			// opt-in
			if checkTag(t.SecondClosestCommentLines, "true") {
				return true
			}
			// all types
			for _, v := range typesWith {
				if v == "*" && !namer.IsPrivateGoName(t.Name.Name) {
					return true
				}
			}
			// For every k8s:validation-gen tag at the package level, interpret the value as a
			// field name (like TypeMeta, ListMeta, ObjectMeta) and trigger validation generation
			// for any type with any of the matching field names. Provides a more useful package
			// level validation than global (because we only need validations on a subset of objects -
			// usually those with TypeMeta).
			return isTypeWith(t, typesWith)
		}

		// Find the right input pkg, which might not be this one.
		inputPath := pkgToInput[input]
		// typesPkg is where the types that need validation are defined.
		// Sometimes it is different from pkg. For example, kubernetes core/v1
		// types are defined in k8s.io/api/core/v1, while the pkg which holds
		// defaulter code is at k/k/pkg/api/v1.
		typesPkg := context.Universe[inputPath]

		// Figure out which types we should be considering further.
		var rootTypes []*types.Type
		for _, t := range typesPkg.Types {
			if shouldCreateObjectValidationFn(t) {
				rootTypes = append(rootTypes, t)
			} else {
				klog.V(6).InfoS("skipping type", "type", t)
			}
		}
		// Deterministic ordering helps in logs and debugging.
		slices.SortFunc(rootTypes, func(a, b *types.Type) int {
			return cmp.Compare(a.Name.String(), b.Name.String())
		})

		//FIXME: new
		//FIXME: make a ctor
		td := &typeDiscoverer{
			validator:   declarativeValidator,
			inputToPkg:  inputToPkg,
			typeNodes:   map[*types.Type]*typeNodeNew{},
			typeTrees:   map[*types.Type]*nodeNew{},
			nodesByPath: map[string]*nodeNew{},
		}
		for _, t := range rootTypes {
			klog.V(4).InfoS("pre-processing", "type", t)
			if node, err := td.discoverTypeNew(t, field.NewPath(t.Name.String())); err != nil {
				klog.Fatalf("failed to generate validations: %v", err)
			} else {
				td.typeTrees[t] = node
				fmt.Println(node.String()) //FIXME:
			}
		}

		//FIXME: old
		for _, t := range rootTypes {
			if _, ok := typeNodes[t]; ok {
				continue // already did this one
			}
			klog.V(4).InfoS("pre-processing", "type", t)

			if err := discoverTypes(declarativeValidator, inputToPkg, t, typeNodes); err != nil {
				klog.Fatalf("OLD: failed to generate validations: %v", err)
			}
		}

		targets = append(targets,
			&generator.SimpleTarget{
				PkgName:       pkg.Name,
				PkgPath:       pkg.Path,
				PkgDir:        pkg.Dir, // output pkg is the same as the input
				HeaderComment: boilerplate,

				FilterFunc: func(c *generator.Context, t *types.Type) bool {
					return t.Name.Package == typesPkg.Path
				},

				GeneratorsFunc: func(c *generator.Context) (generators []generator.Generator) {
					return []generator.Generator{
						NewGenValidations(args.OutputFile, pkg.Path, rootTypes, typeNodes, inputToPkg, declarativeValidator),
					}
				},
			})
	}
	return targets
}

func isTypeWith(t *types.Type, typesWith []string) bool {
	if t.Kind == types.Struct && len(typesWith) > 0 {
		for _, field := range t.Members {
			for _, s := range typesWith {
				if field.Name == s {
					return true
				}
			}
		}
	}
	return false
}
