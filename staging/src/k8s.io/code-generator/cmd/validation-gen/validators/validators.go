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

import (
	"k8s.io/gengo/v2/types"
)

// DeclarativeValidator is able to extract validation function generators from
// types.go files.
type DeclarativeValidator interface {
	// ExtractValidations returns a Validations for the validation this DeclarativeValidator
	// supports for the given go type, and it's corresponding comment strings.
	ExtractValidations(t *types.Type, comments []string) (Validations, error)

	// Docs returns user-friendly documentation for all of the tags that this
	// validator supports.
	Docs() []TagDoc
}

// Validations defines the function calls and variables to generate to perform validation.
type Validations struct {
	Functions []FunctionGen
	Variables []VariableGen
}

func (v *Validations) Empty() bool {
	return len(v.Functions) == 0 && len(v.Variables) == 0
}

func (v *Validations) Len() int {
	return len(v.Functions) + len(v.Variables)
}

func (v *Validations) AddFunction(f FunctionGen) {
	v.Functions = append(v.Functions, f)
}

func (v *Validations) AddVariable(variable VariableGen) {
	v.Variables = append(v.Variables, variable)
}

func (v *Validations) Add(o Validations) {
	v.Functions = append(v.Functions, o.Functions...)
	v.Variables = append(v.Variables, o.Variables...)
}

// FunctionFlags define optional properties of a validator.  Most validators
// can just use DefaultFlags.
type FunctionFlags uint32

// IsSet returns true if all of the wanted flags are set.
func (ff FunctionFlags) IsSet(wanted FunctionFlags) bool {
	return (ff & wanted) == wanted
}

const (
	// DefaultFlags is defined for clarity.
	DefaultFlags FunctionFlags = 0

	// ShortCircuit indicates that further validations should be skipped if
	// this validator fails. Most validators are not fatal.
	ShortCircuit FunctionFlags = 1 << iota

	// NonError indicates that a failure of this validator should not be
	// accumulated as an error, but should trigger other aspects of the failure
	// path (e.g. early return when combined with ShortCircuit).
	NonError
)

// FunctionGen provides validation-gen with the information needed to generate a
// validation function invocation.
type FunctionGen interface {
	// TagName returns the tag which triggers this validator.
	TagName() string

	// SignatureAndArgs returns the function name and all extraArg value literals that are passed when the function
	// invocation is generated.
	//
	// The function signature must be of the form:
	//   func(opCtx operation.Context,
	//        fldPath field.Path,
	//        value, oldValue <ValueType>,     // always nilable
	//        extraArgs[0] <extraArgs[0]Type>, // optional
	//        ...,
	//        extraArgs[N] <extraArgs[N]Type>)
	//
	// extraArgs may contain:
	// - data literals comprised of maps, slices, strings, ints, floats and bools
	// - references, represented by types.Type (to reference any type in the universe), and types.Member (to reference members of the current value)
	//
	// If validation function to be called does not have a signature of this form, please introduce
	// a function that does and use that function to call the validation function.
	SignatureAndArgs() (function types.Name, extraArgs []any)

	// TypeArgs assigns types to the type parameters of the function, for invocation.
	TypeArgs() []types.Name

	// Flags returns the options for this validator function.
	Flags() FunctionFlags
}

// Identifier is a name that the generator will output as an identifier.
// Identifiers are generated using the RawNamer strategy.
type Identifier types.Name

// PrivateVar is a variable name that the generator will output as a private identifier.
// PrivateVars are generated using the PrivateNamer strategy.
type PrivateVar types.Name

// VariableGen provides validation-gen with the information needed to generate variable.
// Variables typically support generated functions by providing static information such
// as the list of supported symbols for an enum.
type VariableGen interface {
	// TagName returns the tag which triggers this validator.
	TagName() string

	// Var returns the variable identifier.
	Var() PrivateVar

	// Init generates the function call that the variable is assigned to.
	Init() FunctionGen
}

// Function creates a FunctionGen for a given function name and extraArgs.
func Function(tagName string, flags FunctionFlags, function types.Name, extraArgs ...any) FunctionGen {
	return GenericFunction(tagName, flags, function, nil, extraArgs...)
}

func GenericFunction(tagName string, flags FunctionFlags, function types.Name, typeArgs []types.Name, extraArgs ...any) FunctionGen {
	// Callers of Signature don't care if the args are all of a known type, it just
	// makes it easier to declare validators.
	var anyArgs []any
	if len(extraArgs) > 0 {
		anyArgs = make([]any, len(extraArgs))
		for i, arg := range extraArgs {
			anyArgs[i] = arg
		}
	}
	return &functionGen{tagName: tagName, flags: flags, function: function, extraArgs: anyArgs, typeArgs: typeArgs}
}

type functionGen struct {
	tagName   string
	function  types.Name
	extraArgs []any
	typeArgs  []types.Name
	flags     FunctionFlags
}

func (v *functionGen) TagName() string {
	return v.tagName
}

func (v *functionGen) SignatureAndArgs() (function types.Name, args []any) {
	return v.function, v.extraArgs
}

func (v *functionGen) TypeArgs() []types.Name { return v.typeArgs }

func (v *functionGen) Flags() FunctionFlags {
	return v.flags
}

// Variable creates a VariableGen for a given function name and extraArgs.
func Variable(variable PrivateVar, init FunctionGen) VariableGen {
	return &variableGen{
		variable: variable,
		init:     init,
	}
}

type variableGen struct {
	variable PrivateVar
	init     FunctionGen
}

func (v variableGen) TagName() string {
	return v.init.TagName()
}

func (v variableGen) Var() PrivateVar {
	return v.variable
}

func (v variableGen) Init() FunctionGen {
	return v.init
}

// TagDoc describes a comment-tag and its usage.
type TagDoc struct {
	// Tag is the tag name, without the leading '+'.
	Tag string
	// Description is a short description of this tag's purpose.
	Description string
	// Contexts lists the place or places this tag may be used.  Tags used in
	// the wrong context may or may not cause errors.
	Contexts []TagContext
	// Payloads lists zero or more varieties of value for this tag. If this tag
	// never has a payload, this list should be empty, but if the payload is
	// optional, this list should include an entry for "<none>".
	Payloads []TagPayloadDoc
}

// TagContext describes where a tag may be attached.
type TagContext string

const (
	// TagContextType indicates that a tag may be attached to a type
	// definition.
	TagContextType TagContext = "Type definition"
	// TagContextField indicates that a tag may be attached to a struct
	// field, the keys of a map, or the values of a map or slice.
	TagContextField TagContext = "Field definition, map key, map/slice value"
)

// TagPayloadDoc describes a value for a tag (e.g `+tagName=tagValue`).  Some
// tags upport multiple payloads, including <none> (e.g. `+tagName`).
type TagPayloadDoc struct {
	Description string
	Docs        string             `json:",omitempty"`
	Schema      []TagPayloadSchema `json:",omitempty"`
}

// TagPayloadSchema describes a JSON tag payload.
type TagPayloadSchema struct {
	Key     string // required
	Value   string // required
	Docs    string `json:",omitempty"`
	Default string `json:",omitempty"`
}
