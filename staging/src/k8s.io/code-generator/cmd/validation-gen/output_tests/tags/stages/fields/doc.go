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

// +k8s:validation-gen=TypeMeta
// +k8s:validation-gen-scheme-registry=k8s.io/code-generator/cmd/validation-gen/testscheme.Scheme

package fields

import "k8s.io/code-generator/cmd/validation-gen/testscheme"

var localSchemeBuilder = testscheme.New()

type Struct struct {
	TypeMeta int

	// +k8s:alpha=+k8s:minimum=10
	IntField int `json:"intField"`

	// +k8s:beta=+k8s:minimum=10
	IntFieldBeta int `json:"intFieldBeta"`

	// +k8s:alpha=+k8s:maxLength=5
	StringField string `json:"stringField"`

	// +k8s:beta=+k8s:maxLength=5
	StringFieldBeta string `json:"stringFieldBeta"`

	// +k8s:alpha=+k8s:maxItems=2
	SliceField []string `json:"sliceField"`

	// +k8s:beta=+k8s:maxItems=2
	SliceFieldBeta []string `json:"sliceFieldBeta"`

	// +k8s:alpha=+k8s:format=k8s-uuid
	UUIDField string `json:"uuidField"`

	// +k8s:beta=+k8s:format=k8s-uuid
	UUIDFieldBeta string `json:"uuidFieldBeta"`

	// +k8s:alpha=+k8s:subfield(inner)=+k8s:minimum=5
	Subfield SubStruct `json:"subfield"`

	// +k8s:beta=+k8s:subfield(inner)=+k8s:minimum=5
	SubfieldBeta SubStruct `json:"subfieldBeta"`

	// +k8s:alpha=+k8s:required
	RequiredField *string `json:"requiredField"`

	// +k8s:beta=+k8s:required
	RequiredFieldBeta *string `json:"requiredFieldBeta"`

	EnumField Enum `json:"enumField"`

	EnumFieldBeta BetaEnum `json:"enumFieldBeta"`

	// +k8s:alpha=+k8s:unionDiscriminator
	D D `json:"d"`

	// +k8s:alpha=+k8s:unionMember
	// +k8s:optional
	M1 *M1 `json:"m1"`

	// +k8s:alpha(since:"1.35")=+k8s:unionMember
	// +k8s:optional
	M2 *M2 `json:"m2"`

	// +k8s:alpha=+k8s:immutable
	ImmutableField string `json:"immutableField"`

	// +k8s:beta=+k8s:immutable
	ImmutableFieldBeta string `json:"immutableFieldBeta"`
}

type UnionStructBeta struct {
	TypeMeta int

	// +k8s:beta=+k8s:unionDiscriminator
	DBeta BetaD `json:"dBeta"`

	// +k8s:beta=+k8s:unionMember
	// +k8s:optional
	M1Beta *BetaM1 `json:"m1Beta"`

	// +k8s:beta(since:"1.35")=+k8s:unionMember
	// +k8s:optional
	M2Beta *BetaM2 `json:"m2Beta"`
}

type SubStruct struct {
	Inner int `json:"inner"`
}

// +k8s:alpha=+k8s:enum
type Enum string

const (
	EnumA Enum = "A"
)

// +k8s:beta=+k8s:enum
type BetaEnum string

const (
	BetaEnumA BetaEnum = "A"
)

type MapItem struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type D string

const (
	DM1 D = "M1"
	DM2 D = "M2"
)

type M1 struct{}
type M2 struct{}

type BetaD string

const (
	BetaDM1 BetaD = "M1Beta"
	BetaDM2 BetaD = "M2Beta"
)

type BetaM1 struct{}
type BetaM2 struct{}

type MixedStruct struct {
	TypeMeta int

	// +k8s:minimum=5
	// +k8s:alpha=+k8s:minimum=10
	IntField int `json:"intField"`

	// +k8s:minimum=5
	// +k8s:beta=+k8s:minimum=10
	IntFieldBeta int `json:"intFieldBeta"`

	// +k8s:maxItems=5
	// +k8s:alpha=+k8s:maxItems=3
	ListField []string `json:"listField"`

	// +k8s:maxItems=5
	// +k8s:beta=+k8s:maxItems=3
	ListFieldBeta []string `json:"listFieldBeta"`
}

type MyStruct struct {
	TypeMeta int

	// +k8s:alpha=+k8s:neq=5
	NEQField int `json:"neqField"`

	// +k8s:beta=+k8s:neq=5
	NEQFieldBeta int `json:"neqFieldBeta"`

	// +k8s:alpha=+k8s:forbidden
	ForbiddenField *string `json:"forbiddenField"`

	// +k8s:beta=+k8s:forbidden
	ForbiddenFieldBeta *string `json:"forbiddenFieldBeta"`

	// +k8s:alpha=+k8s:update=NoModify
	UpdateField string `json:"updateField"`

	// +k8s:beta=+k8s:update=NoModify
	UpdateFieldBeta string `json:"updateFieldBeta"`

	// +k8s:alpha=+k8s:zeroOrOneOfMember
	// +k8s:optional
	Z1 *Z1 `json:"z1"`

	// +k8s:alpha=+k8s:zeroOrOneOfMember
	// +k8s:optional
	Z2 *Z2 `json:"z2"`

	// +k8s:alpha=+k8s:ifEnabled(MyFeature)=+k8s:minimum=10
	ConditionalField int `json:"conditionalField"`

	// +k8s:beta=+k8s:ifEnabled(MyFeature)=+k8s:minimum=10
	ConditionalFieldBeta int `json:"conditionalFieldBeta"`

	// +k8s:alpha=+k8s:alpha=+k8s:minimum=20
	RecursiveAlpha int `json:"recursiveAlpha"`

	// +k8s:beta=+k8s:beta=+k8s:minimum=20
	RecursiveBeta int `json:"recursiveBeta"`
}

type MyStructBeta struct {
	TypeMeta int

	// +k8s:beta=+k8s:zeroOrOneOfMember
	// +k8s:optional
	Z1Beta *BetaZ1 `json:"z1Beta"`

	// +k8s:beta=+k8s:zeroOrOneOfMember
	// +k8s:optional
	Z2Beta *BetaZ2 `json:"z2Beta"`
}

type StructWithValidateFalse struct {
	TypeMeta int

	// +k8s:alpha=+k8s:validateFalse="always fails"
	ValidateFalse *string `json:"validateFalse"`

	// +k8s:beta=+k8s:validateFalse="always fails"
	ValidateFalseBeta *string `json:"validateFalseBeta"`
}

type Z1 struct{}
type Z2 struct{}

type BetaZ1 struct{}
type BetaZ2 struct{}
