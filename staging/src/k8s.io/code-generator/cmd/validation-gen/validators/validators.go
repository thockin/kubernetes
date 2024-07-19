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

package validators

import "k8s.io/gengo/v2/types"

// DeclarativeValidator is able to extract validation function generators from
// types.go files.
type DeclarativeValidator interface {
	// ExtractValidations returns a FunctionGen for each validation this DeclarativeValidator
	// supports for the given go type, and it's corresponding comment strings.
	ExtractValidations(field string, t *types.Type, comments []string) ([]FunctionGen, error)
}

// FunctionGen provides validation-gen with the information needed to generate a
// validation function invocation.
type FunctionGen interface {
	TagName() string

	// SignatureAndArgs returns the function name and all extraArg value literals that are passed when the function
	// invocation is generated.
	//
	// The function signature must be of the form:
	//   func(field.Path, <valueType>, extraArgs[0] <extraArgs[0]Type>, ..., extraArgs[N] <extraArgs[N]Type>)
	//
	// extraArgs may contain strings, ints, floats and bools.
	//
	// If validation function to be called does not have a signature of this form, please introduce
	// a function that does and use that function to call the validation function.
	SignatureAndArgs() (function types.Name, extraArgs []any)

	// IsFatal indicates whether this particular validation function should be
	// considered fatal, or whether further validations may proceed.
	IsFatal() bool
}

// Function creates a FunctionGen for a given function name and extraArgs.
func Function(tagName string, function types.Name, extraArgs ...any) FunctionGen {
	return makeFunction(tagName, false, function, extraArgs...)
}

// FatalFunction creates a fatal-failure FunctionGen for a given function name
// and extraArgs.
func FatalFunction(tagName string, function types.Name, extraArgs ...any) FunctionGen {
	return makeFunction(tagName, true, function, extraArgs...)
}

func makeFunction(tagName string, fatal bool, function types.Name, extraArgs ...any) FunctionGen {
	// Callers of Signature don't care if the args are all of a known type, it just
	// makes it easier to declare validators.
	var anyArgs []any
	if len(extraArgs) > 0 {
		anyArgs = make([]any, len(extraArgs))
		for i, arg := range extraArgs {
			anyArgs[i] = arg
		}
	}
	return &functionGen{tagName: tagName, fatal: fatal, function: function, extraArgs: anyArgs}
}

type functionGen struct {
	tagName   string
	function  types.Name
	extraArgs []any
	fatal     bool
}

func (v *functionGen) TagName() string {
	return v.tagName
}

func (v *functionGen) SignatureAndArgs() (function types.Name, args []any) {
	return v.function, v.extraArgs
}

func (v *functionGen) IsFatal() bool {
	return v.fatal
}
