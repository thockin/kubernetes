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

// deepequal-gen is a tool for auto-generating DeepEqual methods.
// It processes Go source files and generates DeepEqual methods for types
// that match certain criteria (exported types in the target package).
package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/generator"
	"k8s.io/gengo/v2/namer"
	"k8s.io/klog/v2"
)

func main() {
	// Initialize logging
	klog.InitFlags(nil)
	args := &Args{}

	// Set up command line flags
	args.AddFlags(pflag.CommandLine)
	flag.Set("logtostderr", "true")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Validate command line arguments
	if err := args.Validate(); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	// Create a function that will return our generator targets
	myTargets := func(context *generator.Context) []generator.Target {
		return GetTargets(context, args)
	}

	// Run the code generation
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

// Args holds the command-line arguments for the generator
type Args struct {
	// OutputFile is the name of the generated file
	OutputFile string
	// GoHeaderFile is the path to a file containing the header text to add to generated files
	GoHeaderFile string
}

// AddFlags adds the generator flags to the provided flag set
func (args *Args) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&args.OutputFile, "output-file", "zz_generated.deepequal.go", //FIXME: change default
		"the name of the file to be generated")
	fs.StringVar(&args.GoHeaderFile, "go-header-file", "",
		"the path to a file containing boilerplate header text; the string \"YEAR\" will be replaced with the current 4-digit year")
}

// Validate checks that the required arguments are provided
func (args *Args) Validate() error {
	if len(args.OutputFile) == 0 {
		return fmt.Errorf("--output-file must be specified")
	}
	return nil
}

// NameSystems returns the naming system for the generator
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": namer.NewPublicNamer(0),
	}
}

// DefaultNameSystem returns the default naming system
func DefaultNameSystem() string {
	return "public"
}

// GetTargets creates generator targets for each package that needs processing
func GetTargets(context *generator.Context, args *Args) []generator.Target {
	// Get the boilerplate header text
	boilerplate, err := gengo.GoBoilerplate(args.GoHeaderFile, gengo.StdBuildTag, gengo.StdGeneratedBy)
	if err != nil {
		klog.Fatalf("Failed to load boilerplate: %v", err)
	}

	targets := []generator.Target{}

	// Process each input package
	for _, inputPath := range context.Inputs {
		pkg := context.Universe[inputPath]
		if pkg == nil {
			panic(fmt.Sprintf("input package %q was not found in the universe", inputPath))
		}

		// Create a target for this package
		targets = append(targets, &generator.SimpleTarget{
			PkgName:       pkg.Name,
			PkgPath:       pkg.Path,
			PkgDir:        pkg.Dir, // output pkg is the same as the input
			HeaderComment: boilerplate,
			GeneratorsFunc: func(c *generator.Context) []generator.Generator {
				return []generator.Generator{
					NewDeepEqualGenerator(c, pkg, args.OutputFile),
				}
			},
		})
	}

	return targets
}
