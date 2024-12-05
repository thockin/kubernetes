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

package testscheme

import (
	"bytes"
	stdcmp "cmp"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	fuzz "github.com/google/gofuzz"

	"k8s.io/apimachinery/pkg/api/operation"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Scheme is similar to runtime.Scheme, but for validation testing purposes. Scheme only supports validation,
// supports registration of any type (not just runtime.Object) and implements Register directly, allowing it
// to also be used as a scheme builder.
// Must only be used with tests that perform all registration before calls to validate.
type Scheme struct {
	validationFuncs    map[reflect.Type]func(opCtx operation.Context, object, oldObject interface{}, subresources ...string) field.ErrorList
	registrationErrors field.ErrorList
}

// New creates a new Scheme.
func New() *Scheme {
	return &Scheme{validationFuncs: map[reflect.Type]func(opCtx operation.Context, object interface{}, oldObject interface{}, subresources ...string) field.ErrorList{}}
}

// AddValidationFunc registers a validation function.
// Last writer wins.
func (s *Scheme) AddValidationFunc(srcType any, fn func(opCtx operation.Context, object, oldObject interface{}, subresources ...string) field.ErrorList) {
	s.validationFuncs[reflect.TypeOf(srcType)] = fn
}

// Validate validates an object using the registered validation function.
func (s *Scheme) Validate(opts sets.Set[string], object any, subresources ...string) field.ErrorList {
	if len(s.registrationErrors) > 0 {
		return s.registrationErrors // short circuit with registration errors if any are present
	}
	if fn, ok := s.validationFuncs[reflect.TypeOf(object)]; ok {
		return fn(operation.Context{Operation: operation.Create, Options: opts}, object, nil, subresources...)
	}
	return nil
}

// ValidateUpdate validates an update to an object using the registered validation function.
func (s *Scheme) ValidateUpdate(opts sets.Set[string], object, oldObject any, subresources ...string) field.ErrorList {
	if len(s.registrationErrors) > 0 {
		return s.registrationErrors // short circuit with registration errors if any are present
	}
	if fn, ok := s.validationFuncs[reflect.TypeOf(object)]; ok {
		return fn(operation.Context{Operation: operation.Update}, oldObject, object, subresources...)
	}
	return nil
}

// Register adds a scheme setup function to the list.
func (s *Scheme) Register(funcs ...func(*Scheme) error) {
	for _, f := range funcs {
		err := f(s)
		if err != nil {
			s.registrationErrors = append(s.registrationErrors, toRegistrationError(err))
		}
	}
}

func toRegistrationError(err error) *field.Error {
	return field.InternalError(nil, fmt.Errorf("registration error: %w", err))
}

// Test returns a ValidationTestBuilder for this scheme.
func (s *Scheme) Test(t *testing.T) *ValidationTestBuilder {
	return &ValidationTestBuilder{t, s}
}

// ValidationTestBuilder provides convenience functions to build
// validation tests.
type ValidationTestBuilder struct {
	*testing.T
	s *Scheme
}

const fixtureEnvVar = "UPDATE_VALIDATION_GEN_FIXTURE_DATA"

// ValidateFixtures ensures that the validation errors of all registered types match what is expected by the test fixture files.
// For each registered type, a value is created for the type, and populated by fuzzing the value, before validating the type.
// See ValueFuzzed for details.
//
// If the UPDATE_VALIDATION_GEN_FIXTURE_DATA=true environment variable is set, test fixture files are created or overridden.
//
// Fixtures:
//   - validate-false.json: defines a map of registered type to a map of field path to  +validateFalse validations args
//     that are expected to be returned as errors when the type is validated.
func (s *ValidationTestBuilder) ValidateFixtures() {
	flag := os.Getenv(fixtureEnvVar)
	// Run validation
	got := map[string]map[string][]string{}
	for t := range s.s.validationFuncs {
		var v any
		if t.Kind() == reflect.Ptr {
			v = reflect.New(t.Elem()).Interface()
		} else {
			v = reflect.Indirect(reflect.New(t)).Interface()
		}
		if reflect.TypeOf(v).Kind() != reflect.Ptr {
			v = &v
		}
		s.ValueFuzzed(v)
		vt := &ValidationTester{ValidationTestBuilder: s, value: v}
		byPath := vt.ValidateFalseArgsByPath()
		got[t.String()] = byPath
	}

	testdataFilename := "testdata/validate-false.json"
	if flag == "true" {
		// Generate fixture file
		if err := os.MkdirAll(path.Dir(testdataFilename), os.FileMode(0755)); err != nil {
			s.Fatal("error making directory", err)
		}
		data, err := json.MarshalIndent(got, "  ", "  ")
		if err != nil {
			s.Fatal(err)
		}
		err = os.WriteFile(testdataFilename, data, os.FileMode(0644))
		if err != nil {
			s.Fatal(err)
		}
	} else {
		// Load fixture file
		testdataFile, err := os.Open(testdataFilename)
		if errors.Is(err, os.ErrNotExist) {
			s.Fatalf("%s test fixture data not found. Run go test with the environment variable %s=true to create test fixture data.",
				testdataFilename, fixtureEnvVar)
		} else if err != nil {
			s.Fatal(err)
		}
		defer testdataFile.Close()

		byteValue, err := io.ReadAll(testdataFile)
		testdata := map[string]map[string][]string{}
		err = json.Unmarshal(byteValue, &testdata)
		if err != nil {
			s.Fatal(err)
		}
		// Compare fixture with validation results
		expectedKeys := sets.New[string]()
		gotKeys := sets.New[string]()
		for k := range got {
			gotKeys.Insert(k)
		}
		hasErrors := false
		for k, expectedForType := range testdata {
			expectedKeys.Insert(k)
			gotForType, ok := got[k]
			s.T.Run(k, func(t *testing.T) {
				if !ok {
					t.Errorf("%q has expected validateFalse args in %s but got no validation errors.", k, testdataFilename)
					hasErrors = true
				} else {
					if !cmp.Equal(gotForType, expectedForType) {
						t.Errorf("validateFalse args, grouped by field path, differed from %s:\n%s\n",
							testdataFilename, cmp.Diff(gotForType, expectedForType, cmpopts.SortMaps(stdcmp.Less[string])))
						hasErrors = true
					}
				}
			})
		}
		for unexpectedType := range gotKeys.Difference(expectedKeys) {
			s.T.Run(unexpectedType, func(t *testing.T) {
				t.Errorf("%q got unexpected validateFalse args, grouped by field path:\n%s\n",
					unexpectedType, cmp.Diff(nil, got[unexpectedType], cmpopts.SortMaps(stdcmp.Less[string])))
				hasErrors = true
			})
		}
		if hasErrors {
			s.T.Logf("If the test expectations have changed, run go test with the environment variable %s=true", fixtureEnvVar)
		}
	}
}

func fuzzer() *fuzz.Fuzzer {
	return fuzz.New().NilChance(0.0).NumElements(1, 1).RandSource(rand.NewSource(0))
}

// ValueFuzzed automatically populates the given value using a deterministic fuzzer.
// The fuzzer sets pointers to values and always includes a two map keys and slice elements.
func (s *ValidationTestBuilder) ValueFuzzed(value any) *ValidationTester {
	fuzzer().Fuzz(value)
	return &ValidationTester{ValidationTestBuilder: s, value: value}
}

// Value returns a ValidationTester for the given value. The value
// must be a registered with the scheme for validation.
func (s *ValidationTestBuilder) Value(value any) *ValidationTester {
	return &ValidationTester{ValidationTestBuilder: s, value: value}
}

// ValidationTester provides convenience functions to define validation
// tests for a validatable value.
type ValidationTester struct {
	*ValidationTestBuilder
	value    any
	oldValue any
	opts     sets.Set[string]
}

// OldValue sets the oldValue for this ValidationTester. When oldValue is set to
// a non-nil value, update validation will be used to test validation.
// oldValue must be the same type as value.
// Returns ValidationTester to support call chaining.
func (v *ValidationTester) OldValue(oldValue any) *ValidationTester {
	v.oldValue = oldValue
	return v
}

// OldValueFuzzed automatically populates the given value using a deterministic fuzzer.
// The fuzzer sets pointers to values and always includes a two map keys and slice elements.
func (v *ValidationTester) OldValueFuzzed(oldValue any) *ValidationTester {
	fuzzer().Fuzz(oldValue)
	v.oldValue = oldValue
	return v
}

// Opts sets the ValidationOpts to use.
func (v *ValidationTester) Opts(opts sets.Set[string]) *ValidationTester {
	v.opts = opts
	return v
}

func multiline(errs field.ErrorList) string {
	if len(errs) == 0 {
		return "<no errors>"
	}
	if len(errs) == 1 {
		return errs[0].Error()
	}

	var buf bytes.Buffer
	for _, err := range errs {
		buf.WriteString("\n")
		buf.WriteString(err.Error())
	}
	return buf.String()
}

// ExpectValid validates the value and calls t.Errorf if any validation errors are returned.
// Returns ValidationTester to support call chaining.
func (v *ValidationTester) ExpectValid() *ValidationTester {
	v.T.Run(fmt.Sprintf("%T", v.value), func(t *testing.T) {
		errs := v.validate()
		if len(errs) > 0 {
			t.Errorf("want no errors, got: %v", multiline(errs))
		}
	})
	return v
}

// ExpectValidAt validates the value and calls t.Errorf for any validation errors at the given path.
// Returns ValidationTester to support call chaining.
func (v *ValidationTester) ExpectValidAt(fldPath *field.Path) *ValidationTester {
	v.T.Run(fmt.Sprintf("%T.%v", v.value, fldPath), func(t *testing.T) {
		var got field.ErrorList
		for _, e := range v.validate() {
			if e.Field == fldPath.String() {
				got = append(got, e)
			}
		}
		if len(got) > 0 {
			t.Errorf("want no errors at %v, got: %v", fldPath, got)
		}
	})
	return v
}

// ExpectInvalid validates the value and calls t.Errorf if want does not match the actual errors.
// Returns ValidationTester to support call chaining.
func (v *ValidationTester) ExpectInvalid(want ...*field.Error) *ValidationTester {
	return v.expectInvalid(byFullError, want...)
}

// ExpectValidateFalse validates the value and calls t.Errorf if the actual errors do not
// match the given validateFalseArgs.  For example, if the value to validate has a
// single `+validateFalse="type T1"` tag, ExpectValidateFalse("type T1") will pass.
// Returns ValidationTester to support call chaining.
func (v *ValidationTester) ExpectValidateFalse(validateFalseArgs ...string) *ValidationTester {
	var want []*field.Error
	for _, s := range validateFalseArgs {
		want = append(want, field.Invalid(nil, "", fmt.Sprintf("forced failure: %s", s)))
	}
	return v.expectInvalid(byDetail, want...)
}

func (v *ValidationTester) ExpectValidateFalseByPath(validateFalseArgsByField map[string][]string) *ValidationTester {
	v.T.Run(fmt.Sprintf("%T", v.value), func(t *testing.T) {
		byField := v.ValidateFalseArgsByPath()
		// ensure args are sorted
		for _, args := range validateFalseArgsByField {
			sort.Strings(args)
		}
		if !cmp.Equal(validateFalseArgsByField, byField) {
			t.Errorf("validateFalse args, grouped by field path, differed from expected:\n%s\n", cmp.Diff(validateFalseArgsByField, byField, cmpopts.SortMaps(stdcmp.Less[string])))
		}

	})
	return v
}

func (v *ValidationTester) ValidateFalseArgsByPath() map[string][]string {
	byField := map[string][]string{}
	errs := v.validate()
	for _, e := range errs {
		if strings.HasPrefix(e.Detail, "forced failure: ") {
			arg := strings.TrimPrefix(e.Detail, "forced failure: ")
			f := e.Field
			if f == "<nil>" {
				f = ""
			}
			byField[f] = append(byField[f], arg)
		}
	}
	// ensure args are sorted
	for _, args := range byField {
		sort.Strings(args)
	}
	return byField
}

func (v *ValidationTester) expectInvalid(matcher matcher, errs ...*field.Error) *ValidationTester {
	v.T.Run(fmt.Sprintf("%T", v.value), func(t *testing.T) {
		want := sets.New[string]()
		for _, e := range errs {
			want.Insert(matcher(e))
		}

		got := sets.New[string]()
		for _, e := range v.validate() {
			got.Insert(matcher(e))
		}
		if !got.Equal(want) {
			t.Errorf("validation errors differed from expected:\n%v\n", cmp.Diff(want, got, cmpopts.SortMaps(stdcmp.Less[string])))

			for x := range got.Difference(want) {
				fmt.Printf("%q,\n", strings.TrimPrefix(x, "forced failure: "))
			}
		}
	})
	return v
}

type matcher func(err *field.Error) string

func byDetail(err *field.Error) string {
	return err.Detail
}

func byFullError(err *field.Error) string {
	return err.Error()
}

func (v *ValidationTester) validate() field.ErrorList {
	var errs field.ErrorList
	if v.oldValue == nil {
		errs = v.s.Validate(v.opts, v.value)
	} else {
		errs = v.s.ValidateUpdate(v.opts, v.value, v.oldValue)
	}
	return errs
}
