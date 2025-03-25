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

// Package output_test verifies the functionality of the generated DeepEqual methods
package output_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

// TestAgainstSelf verifies that a struct equals itself.
// This is a basic sanity check that should always pass.
func TestAgainstSelf(t *testing.T) {
	// Create a fuzzer with:
	// - 50% chance of nil pointers
	// - 0-5 elements in slices/maps
	// - Maximum recursion depth of 5
	f := fuzz.New().NilChance(0.5).NumElements(0, 5).MaxDepth(5)

	for i := 0; i < 1000; i++ {
		var s1 Struct
		f.Fuzz(&s1)

		// Create a copy and change its pointers
		s2 := s1
		s2.ChangePointers()

		// They should still be equal
		if !s1.DeepEqual(&s2) {
			t.Fatalf("iteration %d: object not equal after ChangePointers: %#v", i, s1)
		}
	}
}

// TestAgainstDifferent verifies that two independently fuzzed structs are not equal.
// This helps verify that the DeepEqual method can detect differences.
func TestAgainstDifferent(t *testing.T) {
	f := fuzz.New().NilChance(0.5).NumElements(0, 5).MaxDepth(5)

	for i := 0; i < 1000; i++ {
		var s1, s2 Struct
		f.Fuzz(&s1)
		f.Fuzz(&s2)

		// Two independently fuzzed structs should not be equal
		if s1.DeepEqual(&s2) {
			t.Fatalf("iteration %d: different objects unexpectedly equal:\nfirst: %#v\nsecond: %#v", i, s1, s2)
		}
	}
}

// TestAgainstEmpty verifies that a fuzzed struct is not equal to an empty struct.
// This helps verify that the DeepEqual method properly compares all fields.
func TestAgainstEmpty(t *testing.T) {
	f := fuzz.New().NilChance(0.5).NumElements(0, 5).MaxDepth(5)

	for i := 0; i < 1000; i++ {
		var s Struct
		f.Fuzz(&s)

		// A fuzzed struct should not equal an empty struct
		empty := Struct{}
		if s.DeepEqual(&empty) {
			t.Fatalf("iteration %d: object unexpectedly equal to empty struct: %#v", i, s)
		}
	}
}
