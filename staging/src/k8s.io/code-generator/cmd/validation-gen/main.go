/*
Copyright 20124The Kubernetes Authors.

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

// validation-gen is a tool for auto-generating Validation functions.
package main

import (
	"cmp"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/pflag"

	"k8s.io/code-generator/cmd/validation-gen/validators"
	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	args := &Args{}

	args.AddFlags(pflag.CommandLine)
	flag.Set("logtostderr", "true")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := args.Validate(); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	if args.PrintDocs {
		printDocs()
		os.Exit(0)
	}

	myTargets := func(context *generator.Context) []generator.Target {
		return GetTargets(context, args)
	}

	// Run it.
	if err := gengo.Execute(
		NameSystems(),
		DefaultNameSystem(),
		myTargets,
		gengo.StdBuildTag,
		pflag.Args(),
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}
	klog.V(2).Info("Completed successfully.")
}

type Args struct {
	OutputFile   string
	ExtraPkgs    []string // Always consider these as last-ditch possibilities for validations.
	GoHeaderFile string
	PrintDocs    bool
}

// AddFlags add the generator flags to the flag set.
func (args *Args) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&args.OutputFile, "output-file", "generated.validations.go",
		"the name of the file to be generated")
	fs.StringSliceVar(&args.ExtraPkgs, "extra-pkg", args.ExtraPkgs,
		"the import path of a package whose validation can be used by generated code, but is not being generated for")
	fs.StringVar(&args.GoHeaderFile, "go-header-file", "",
		"the path to a file containing boilerplate header text; the string \"YEAR\" will be replaced with the current 4-digit year")
	fs.BoolVar(&args.PrintDocs, "docs", false,
		"print documentation for supported declarative validations, and then exit")
}

// Validate checks the given arguments.
func (args *Args) Validate() error {
	if len(args.OutputFile) == 0 {
		return fmt.Errorf("--output-file must be specified")
	}

	return nil
}

func printDocs() {
	// We need a fake context to init the validator plugins.
	c := &generator.Context{
		Namers:    namer.NameSystems{},
		Universe:  types.Universe{},
		FileTypes: map[string]generator.FileType{},
	}

	// This gets a composite validator which aggregates the many plugins.
	validator := validators.NewValidator(c, nil, nil)

	docs := builtinTagDocs()
	docs = append(docs, validator.Docs()...)
	slices.SortFunc(docs, func(a, b validators.TagDoc) int {
		return cmp.Compare(a.Tag, b.Tag)
	})

	if jb, err := json.MarshalIndent(docs, "", "    "); err != nil {
		klog.Fatalf("failed to marshal docs: %v", err)
	} else {
		fmt.Println(string(jb))
	}
}
