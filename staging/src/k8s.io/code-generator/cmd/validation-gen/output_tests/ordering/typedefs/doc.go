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

// +k8s:validation-gen=*

// This is a test package.
package typedefs

// Treat these as 4 bits, and ensure all combinations
//   bit 0: no flags
//   bit 1: PtrOK
//   bit 2: IsFatal
//   bit 3: PtrOK | IsFatal

// Note: No validations.
type E00 string

// +validateTrue={"flags":[], "msg":"E01, no flags"}
type E01 string

// +validateTrue={"flags":["PtrOK"], "msg":"E02, PtrOK"}
type E02 string

// +validateTrue={"flags":[], "msg":"E03, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"E03, PtrOK"}
type E03 string

// +validateTrue={"flags":["IsFatal"], "msg":"E04, IsFatal"}
type E04 string

// +validateTrue={"flags":[], "msg":"E05, no flags"}
// +validateTrue={"flags":["IsFatal"], "msg":"E05, IsFatal"}
type E05 string

// +validateTrue={"flags":["PtrOK"], "msg":"E06, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"E06, IsFatal"}
type E06 string

// +validateTrue={"flags":[], "msg":"E07, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"E07, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"E07, IsFatal"}
type E07 string

// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E08, PtrOK|IsFatal"}
type E08 string

// +validateTrue={"flags":[], "msg":"E09, no flags"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E09, PtrOK|IsFatal"}
type E09 string

// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E10, PtrOK|IsFatal"}
// +validateTrue={"flags":["PtrOK"], "msg":"E10, PtrOK"}
type E10 string

// +validateTrue={"flags":[], "msg":"E11, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"E11, PtrOK"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E11, PtrOK|IsFatal"}
type E11 string

// +validateTrue={"flags":["IsFatal"], "msg":"E12, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E12, PtrOK|IsFatal"}
type E12 string

// +validateTrue={"flags":[], "msg":"E13, no flags"}
// +validateTrue={"flags":["IsFatal"], "msg":"E13, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E13, PtrOK|IsFatal"}
type E13 string

// +validateTrue={"flags":["PtrOK"], "msg":"E14, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"E14, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E14, PtrOK|IsFatal"}
type E14 string

// +validateTrue={"flags":[], "msg":"E15, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"E15, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"E15, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"E15, PtrOK|IsFatal"}
type E15 string

// Note: these are intentionally in the wrong final order.
// +validateTrue={"flags":[], "msg":"EMultiple, no flags 1"}
// +validateTrue={"flags":["PtrOK"], "msg":"EMultiple, PtrOK 1"}
// +validateTrue={"flags":["IsFatal"], "msg":"EMultiple, IsFatal 1"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"EMultiple, PtrOK|IsFatal 1"}
// +validateTrue="E0, string payload"
// +validateTrue={"flags":[], "msg":"EMultiple, no flags 2"}
// +validateTrue={"flags":["PtrOK"], "msg":"EMultiple, PtrOK 2"}
// +validateTrue={"flags":["IsFatal"], "msg":"EMultiple, IsFatal 2"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"EMultiple, PtrOK|IsFatal 2"}
type EMultiple string
