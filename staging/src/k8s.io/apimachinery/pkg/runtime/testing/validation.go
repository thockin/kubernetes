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

package testing

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

type VersionValidationRunner func(t *testing.T, versionValidationErrors field.ErrorList)

// RunValidationForEachVersion runs f as a subtest of t for each version of the given unversioned object.
// Each subtest is named by GroupVersionKind.  Each call to f is provided the field.ErrorList results
// of converting the unversioned object to a version and validating it.
//
// Only autogenerated validation is run. To test both handwritten and autogenerated validation:
//
//		RunValidationForEachVersion(t, testCase.pod, func(t *testing.T, versionValidationErrors field.ErrorList) {
//		  errs := ValidatePod(testCase.obj) // hand written validation
//	      errs = append(errs, versionValidationErrors...) // generated declarative validation
//		  // Validate that the errors are what was expected for this test case.
//		})
func RunValidationForEachVersion(t *testing.T, unversioned runtime.Object, fn VersionValidationRunner) {
	runValidation(t, unversioned, fn)
}

// RunUpdateValidationForEachVersion is like RunValidationForEachVersion but for update validation.
func RunUpdateValidationForEachVersion(t *testing.T, unversioned, unversionedOld runtime.Object, fn VersionValidationRunner) {
	runUpdateValidation(t, unversioned, unversionedOld, fn)
}

// RunStatusValidationForEachVersion is like RunUpdateValidationForEachVersion but for status validation.
func RunStatusValidationForEachVersion(t *testing.T, unversioned, unversionedOld runtime.Object, fn VersionValidationRunner) {
	runUpdateValidation(t, unversioned, unversionedOld, fn, "status")
}

func runValidation(t *testing.T, unversioned runtime.Object, fn VersionValidationRunner, subresources ...string) {
	unversionedGVKs, _, err := legacyscheme.Scheme.ObjectKinds(unversioned)
	if err != nil {
		t.Fatal(err)
	}
	for _, unversionedGVK := range unversionedGVKs {
		gvs := legacyscheme.Scheme.VersionsForGroupKind(unversionedGVK.GroupKind())
		for _, gv := range gvs {
			gvk := gv.WithKind(unversionedGVK.Kind)
			t.Run(gvk.String(), func(t *testing.T) {
				if gvk.Version != runtime.APIVersionInternal { // skip internal
					versioned, err := legacyscheme.Scheme.New(gvk)
					if err != nil {
						t.Fatal(err)
					}
					err = legacyscheme.Scheme.Convert(unversioned, versioned, nil)
					if err != nil {
						t.Fatal(err)
					}
					fn(t, legacyscheme.Scheme.Validate(versioned, subresources...))
				}
			})
		}
	}
}

func runUpdateValidation(t *testing.T, unversionedNew, unversionedOld runtime.Object, fn VersionValidationRunner, subresources ...string) {
	unversionedGVKs, _, err := legacyscheme.Scheme.ObjectKinds(unversionedNew)
	if err != nil {
		t.Fatal(err)
	}
	for _, unversionedGVK := range unversionedGVKs {
		gvs := legacyscheme.Scheme.VersionsForGroupKind(unversionedGVK.GroupKind())
		for _, gv := range gvs {
			gvk := gv.WithKind(unversionedGVK.Kind)
			t.Run(gvk.String(), func(t *testing.T) {
				if gvk.Version != runtime.APIVersionInternal { // skip internal
					versionedNew, err := legacyscheme.Scheme.New(gvk)
					if err != nil {
						t.Fatal(err)
					}
					err = legacyscheme.Scheme.Convert(unversionedNew, versionedNew, nil)
					if err != nil {
						t.Fatal(err)
					}

					var versionedOld runtime.Object
					if unversionedOld != nil {
						versionedOld, err = legacyscheme.Scheme.New(gvk)
						if err != nil {
							t.Fatal(err)
						}

						err = legacyscheme.Scheme.Convert(unversionedOld, versionedOld, nil)
						if err != nil {
							t.Fatal(err)
						}
					}

					fn(t, legacyscheme.Scheme.ValidateUpdate(versionedNew, versionedOld, subresources...))
				}
			})
		}
	}
}
