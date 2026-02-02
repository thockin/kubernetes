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

package fields

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
)

func TestAlpha(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&Struct{
		IntField:           10,
		StringField:        "abc",
		SliceField:         []string{"a", "b"},
		UUIDField:          "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		Subfield:           SubStruct{Inner: 5},
		RequiredField:      ptr.To("val"),
		EnumField:          "A",
		D:                  DM1,
		M1:                 &M1{},
		ImmutableField:     "foo",
		IntFieldBeta:       10,
		StringFieldBeta:    "abc",
		SliceFieldBeta:     []string{"a", "b"},
		UUIDFieldBeta:      "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		SubfieldBeta:       SubStruct{Inner: 5},
		RequiredFieldBeta:  ptr.To("val"),
		EnumFieldBeta:      "A",
		ImmutableFieldBeta: "foo",
	}).ExpectValid()

	// Test failures marked as alpha
	st.Value(&Struct{
		IntField:           5,
		StringField:        "too-long",
		SliceField:         []string{"a", "b", "c"},
		UUIDField:          "not-a-uuid",
		Subfield:           SubStruct{Inner: 1},
		RequiredField:      nil,
		EnumField:          "B",
		D:                  DM1,
		M1:                 nil, // required by discriminator
		ImmutableField:     "bar",
		IntFieldBeta:       10,
		StringFieldBeta:    "abc",
		SliceFieldBeta:     []string{"a", "b"},
		UUIDFieldBeta:      "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		SubfieldBeta:       SubStruct{Inner: 5},
		RequiredFieldBeta:  ptr.To("val"),
		EnumFieldBeta:      "A",
		ImmutableFieldBeta: "foo",
	}).OldValue(&Struct{
		RequiredField:      ptr.To("old"),
		ImmutableField:     "foo",
		RequiredFieldBeta:  ptr.To("old"),
		ImmutableFieldBeta: "foo",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("intField"), 5, "").WithOrigin("minimum").MarkAlpha(),
		field.TooLong(field.NewPath("stringField"), "too-long", 5).WithOrigin("maxLength").MarkAlpha(),
		field.TooMany(field.NewPath("sliceField"), 3, 2).WithOrigin("maxItems").MarkAlpha(),
		field.Invalid(field.NewPath("uuidField"), "not-a-uuid", "").WithOrigin("format=k8s-uuid").MarkAlpha(),
		field.Invalid(field.NewPath("subfield", "inner"), 1, "").WithOrigin("minimum").MarkAlpha(),
		field.Required(field.NewPath("requiredField"), "").MarkAlpha(),
		field.NotSupported(field.NewPath("enumField"), Enum("B"), []string{"A"}).MarkAlpha(),
		field.Invalid(field.NewPath("m1"), nil, "").WithOrigin("union").MarkAlpha(),
		field.Invalid(field.NewPath("immutableField"), "bar", "").WithOrigin("immutable").MarkAlpha(),
	})
}

func TestBeta(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&Struct{
		IntField:           10,
		StringField:        "abc",
		SliceField:         []string{"a", "b"},
		UUIDField:          "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		Subfield:           SubStruct{Inner: 5},
		RequiredField:      ptr.To("val"),
		EnumField:          "A",
		D:                  DM1,
		M1:                 &M1{},
		ImmutableField:     "foo",
		IntFieldBeta:       10,
		StringFieldBeta:    "abc",
		SliceFieldBeta:     []string{"a", "b"},
		UUIDFieldBeta:      "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		SubfieldBeta:       SubStruct{Inner: 5},
		RequiredFieldBeta:  ptr.To("val"),
		EnumFieldBeta:      "A",
		ImmutableFieldBeta: "foo",
	}).ExpectValid()

	// Test failures marked as beta
	st.Value(&Struct{
		IntField:           10,
		StringField:        "abc",
		SliceField:         []string{"a", "b"},
		UUIDField:          "a0a2a2d2-0b87-4964-a123-78d00a8787a6",
		Subfield:           SubStruct{Inner: 5},
		RequiredField:      ptr.To("val"),
		EnumField:          "A",
		D:                  DM1,
		M1:                 &M1{},
		ImmutableField:     "foo",
		IntFieldBeta:       5,
		StringFieldBeta:    "too-long",
		SliceFieldBeta:     []string{"a", "b", "c"},
		UUIDFieldBeta:      "not-a-uuid",
		SubfieldBeta:       SubStruct{Inner: 1},
		RequiredFieldBeta:  nil,
		EnumFieldBeta:      "B",
		ImmutableFieldBeta: "bar",
	}).OldValue(&Struct{
		RequiredField:      ptr.To("old"),
		ImmutableField:     "foo",
		RequiredFieldBeta:  ptr.To("old"),
		ImmutableFieldBeta: "foo",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("intFieldBeta"), 5, "").WithOrigin("minimum").MarkBeta(),
		field.TooLong(field.NewPath("stringFieldBeta"), "too-long", 5).WithOrigin("maxLength").MarkBeta(),
		field.TooMany(field.NewPath("sliceFieldBeta"), 3, 2).WithOrigin("maxItems").MarkBeta(),
		field.Invalid(field.NewPath("uuidFieldBeta"), "not-a-uuid", "").WithOrigin("format=k8s-uuid").MarkBeta(),
		field.Invalid(field.NewPath("subfieldBeta", "inner"), 1, "").WithOrigin("minimum").MarkBeta(),
		field.Required(field.NewPath("requiredFieldBeta"), "").MarkBeta(),
		field.NotSupported(field.NewPath("enumFieldBeta"), BetaEnum("B"), []string{"A"}).MarkBeta(),
		field.Invalid(field.NewPath("immutableFieldBeta"), "bar", "").WithOrigin("immutable").MarkBeta(),
	})
}

func TestUnionBeta(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&UnionStructBeta{
		DBeta:  BetaDM1,
		M1Beta: &BetaM1{},
	}).ExpectValid()

	st.Value(&UnionStructBeta{
		DBeta:  BetaDM1,
		M1Beta: nil, // required by discriminator
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("m1Beta"), nil, "").WithOrigin("union").MarkBeta(),
	})
}

func TestMixed(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	// Valid case (meets both normal and alpha requirements)
	st.Value(&MixedStruct{
		IntField:      15,
		IntFieldBeta:  15,
		ListField:     []string{"a", "b", "c"},
		ListFieldBeta: []string{"a", "b", "c"},
	}).ExpectValid()

	// Fails alpha validation but passes normal validation
	// IntField: 5 <= 8 < 10 (alpha fails)
	// ListField: 3 < 4 <= 5 (alpha fails)
	st.Value(&MixedStruct{
		IntField:      8,
		IntFieldBeta:  8,
		ListField:     []string{"a", "b", "c", "d"},
		ListFieldBeta: []string{"a", "b", "c", "d"},
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("intField"), 8, "").WithOrigin("minimum").MarkAlpha(),
		field.Invalid(field.NewPath("intFieldBeta"), 8, "").WithOrigin("minimum").MarkBeta(),
		field.TooMany(field.NewPath("listField"), 4, 3).WithOrigin("maxItems").MarkAlpha(),
		field.TooMany(field.NewPath("listFieldBeta"), 4, 3).WithOrigin("maxItems").MarkBeta(),
	})

	// Fails both normal and alpha validation
	// IntField: 4 < 5 (normal fails) AND 4 < 10 (alpha fails)
	// ListField: 6 > 5 (normal fails) AND 6 > 3 (alpha fails)
	st.Value(&MixedStruct{
		IntField:      4,
		IntFieldBeta:  4,
		ListField:     []string{"a", "b", "c", "d", "e", "f"},
		ListFieldBeta: []string{"a", "b", "c", "d", "e", "f"},
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("intField"), 4, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("intField"), 4, "").WithOrigin("minimum").MarkAlpha(),
		field.Invalid(field.NewPath("intFieldBeta"), 4, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("intFieldBeta"), 4, "").WithOrigin("minimum").MarkBeta(),
		field.TooMany(field.NewPath("listField"), 6, 5).WithOrigin("maxItems"),
		field.TooMany(field.NewPath("listField"), 6, 3).WithOrigin("maxItems").MarkAlpha(),
		field.TooMany(field.NewPath("listFieldBeta"), 6, 5).WithOrigin("maxItems"),
		field.TooMany(field.NewPath("listFieldBeta"), 6, 3).WithOrigin("maxItems").MarkBeta(),
	})
}

func TestMyStruct(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&MyStruct{
		NEQField:             10,
		NEQFieldBeta:         10,
		ConditionalField:     15,
		ConditionalFieldBeta: 15,
		RecursiveAlpha:       25,
		RecursiveBeta:        25,
	}).Opts([]string{"MyFeature"}).ExpectValid()

	st.Value(&MyStruct{
		NEQField:             5,
		NEQFieldBeta:         5,
		ForbiddenField:       ptr.To("val"),
		ForbiddenFieldBeta:   ptr.To("val"),
		UpdateField:          "new",
		UpdateFieldBeta:      "new",
		Z1:                   &Z1{},
		Z2:                   &Z2{},
		ConditionalField:     5,
		ConditionalFieldBeta: 5,
		RecursiveAlpha:       10,
		RecursiveBeta:        10,
	}).Opts([]string{"MyFeature"}).OldValue(&MyStruct{
		UpdateField:     "old",
		UpdateFieldBeta: "old",
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("neqField"), 5, "must not be 5").WithOrigin("neq").MarkAlpha(),
		field.Invalid(field.NewPath("neqFieldBeta"), 5, "must not be 5").WithOrigin("neq").MarkBeta(),
		field.Forbidden(field.NewPath("forbiddenField"), "").MarkAlpha(),
		field.Forbidden(field.NewPath("forbiddenFieldBeta"), "").MarkBeta(),
		field.Invalid(field.NewPath("updateField"), "new", "field cannot be modified once set").WithOrigin("update").MarkAlpha(),
		field.Invalid(field.NewPath("updateFieldBeta"), "new", "field cannot be modified once set").WithOrigin("update").MarkBeta(),
		field.Invalid(nil, &MyStruct{
			NEQField: 5, NEQFieldBeta: 5,
			ForbiddenField: ptr.To("val"), ForbiddenFieldBeta: ptr.To("val"),
			UpdateField: "new", UpdateFieldBeta: "new",
			Z1: &Z1{}, Z2: &Z2{},
			ConditionalField: 5, ConditionalFieldBeta: 5,
			RecursiveAlpha: 10, RecursiveBeta: 10,
		}, "only one of z1, z2 may be specified").WithOrigin("zeroOrOneOf").MarkAlpha(),
		field.Invalid(field.NewPath("conditionalField"), 5, "").WithOrigin("minimum").MarkAlpha(),
		field.Invalid(field.NewPath("conditionalFieldBeta"), 5, "").WithOrigin("minimum").MarkBeta(),
		field.Invalid(field.NewPath("recursiveAlpha"), 10, "").WithOrigin("minimum").MarkAlpha(),
		field.Invalid(field.NewPath("recursiveBeta"), 10, "").WithOrigin("minimum").MarkBeta(),
	})

	st.Value(&StructWithValidateFalse{
		ValidateFalse:     ptr.To("val"),
		ValidateFalseBeta: ptr.To("val"),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(field.NewPath("validateFalse"), "val", "always fails").MarkAlpha(),
		field.Invalid(field.NewPath("validateFalseBeta"), "val", "always fails").MarkBeta(),
	})
}

func TestMyStructBeta(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&MyStructBeta{
		Z1Beta: &BetaZ1{},
		Z2Beta: &BetaZ2{},
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByOrigin().ByValidationStabilityLevel(), field.ErrorList{
		field.Invalid(nil, &MyStructBeta{
			Z1Beta: &BetaZ1{}, Z2Beta: &BetaZ2{},
		}, "only one of z1Beta, z2Beta may be specified").WithOrigin("zeroOrOneOf").MarkBeta(),
	})
}
