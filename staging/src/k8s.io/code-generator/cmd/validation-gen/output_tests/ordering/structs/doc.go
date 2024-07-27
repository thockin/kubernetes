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

// +k8s:validation-gen=TypeMeta

// This is a test package.
package structs

type Tother struct {
	// +validateTrue={"flags":[], "msg":"Tother, no flags"}
	OS string `json:"os"`
}

// Treat these as 4 bits, and ensure all combinations
//   bit 0: no flags
//   bit 1: PtrOK
//   bit 2: IsFatal
//   bit 3: PtrOK | IsFatal

// Note: No validations.
type T00 struct {
	TypeMeta int
	S        string  `json:"s"`
	PS       *string `json:"ps"`
	T        Tother  `json:"t"`
	PT       *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T01, no flags"}
type T01 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T01.S, no flags"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T01.PS, no flags"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T01.T, no flags"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T01.PT, no flags"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["PtrOK"], "msg":"T02, PtrOK"}
type T02 struct {
	TypeMeta int
	// +validateTrue={"flags":["PtrOK"], "msg":"T02.S, PtrOK"}
	S string `json:"s"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T02.PS, PtrOK"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T02.T, PtrOK"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T02.PT, PtrOK"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T03, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"T03, PtrOK"}
type T03 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T03.S, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T03.S, PtrOK"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T03.PS, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T03.PS, PtrOK"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T03.T, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T03.T, PtrOK"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T03.PT, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T03.PT, PtrOK"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["IsFatal"], "msg":"T04, IsFatal"}
type T04 struct {
	TypeMeta int
	// +validateTrue={"flags":["IsFatal"], "msg":"T04.S, IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T04.PS, IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T04.T, IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T04.PT, IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T05, no flags"}
// +validateTrue={"flags":["IsFatal"], "msg":"T05, IsFatal"}
type T05 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T05.S, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T05.S, IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T05.PS, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T05.PS, IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T05.T, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T05.T, IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T05.PT, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T05.PT, IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["PtrOK"], "msg":"T06, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"T06, IsFatal"}
type T06 struct {
	TypeMeta int
	// +validateTrue={"flags":["PtrOK"], "msg":"T06.S, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T06.S, IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T06.PS, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T06.PS, IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T06.T, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T06.T, IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T06.PT, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T06.PT, IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T07, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"T07, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"T07, IsFatal"}
type T07 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T07.S, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T07.S, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T07.S, IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T07.PS, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T07.PS, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T07.PS, IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T07.T, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T07.T, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T07.T, IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T07.PT, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T07.PT, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T07.PT, IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T08, PtrOK|IsFatal"}
type T08 struct {
	TypeMeta int
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T08.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T08.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T08.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T08.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T09, no flags"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T09, PtrOK|IsFatal"}
type T09 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T09.S, no flags"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T09.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T09.PS, no flags"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T09.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T09.T, no flags"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T09.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T09.PT, no flags"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T09.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T10, PtrOK|IsFatal"}
// +validateTrue={"flags":["PtrOK"], "msg":"T10, PtrOK"}
type T10 struct {
	TypeMeta int
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T10.S, PtrOK|IsFatal"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T10.S, PtrOK"}
	S string `json:"s"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T10.PS, PtrOK|IsFatal"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T10.PS, PtrOK"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T10.T, PtrOK|IsFatal"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T10.T, PtrOK"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T10.PT, PtrOK|IsFatal"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T10.PT, PtrOK"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T11, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"T11, PtrOK"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T11, PtrOK|IsFatal"}
type T11 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T11.S, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T11.S, PtrOK"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T11.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T11.PS, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T11.PS, PtrOK"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T11.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T11.T, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T11.T, PtrOK"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T11.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T11.PT, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T11.PT, PtrOK"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T11.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["IsFatal"], "msg":"T12, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T12, PtrOK|IsFatal"}
type T12 struct {
	TypeMeta int
	// +validateTrue={"flags":["IsFatal"], "msg":"T12.S, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T12.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T12.PS, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T12.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T12.T, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T12.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["IsFatal"], "msg":"T12.PT, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T12.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T13, no flags"}
// +validateTrue={"flags":["IsFatal"], "msg":"T13, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T13, PtrOK|IsFatal"}
type T13 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T13.S, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T13.S, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T13.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T13.PS, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T13.PS, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T13.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T13.T, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T13.T, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T13.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T13.PT, no flags"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T13.PT, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T13.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":["PtrOK"], "msg":"T14, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"T14, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T14, PtrOK|IsFatal"}
type T14 struct {
	TypeMeta int
	// +validateTrue={"flags":["PtrOK"], "msg":"T14.S, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T14.S, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T14.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T14.PS, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T14.PS, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T14.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T14.T, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T14.T, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T14.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":["PtrOK"], "msg":"T14.PT, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T14.PT, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T14.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// +validateTrue={"flags":[], "msg":"T15, no flags"}
// +validateTrue={"flags":["PtrOK"], "msg":"T15, PtrOK"}
// +validateTrue={"flags":["IsFatal"], "msg":"T15, IsFatal"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T15, PtrOK|IsFatal"}
type T15 struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"T15.S, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T15.S, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T15.S, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T15.S, PtrOK|IsFatal"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"T15.PS, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T15.PS, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T15.PS, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T15.PS, PtrOK|IsFatal"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"T15.T, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T15.T, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T15.T, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T15.T, PtrOK|IsFatal"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"T15.PT, no flags"}
	// +validateTrue={"flags":["PtrOK"], "msg":"T15.PT, PtrOK"}
	// +validateTrue={"flags":["IsFatal"], "msg":"T15.PT, IsFatal"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"T15.PT, PtrOK|IsFatal"}
	PT *Tother `json:"pt"`
}

// Note: these are intentionally in the wrong final order.
// +validateTrue={"flags":[], "msg":"TMultiple, no flags 1"}
// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple, PtrOK 1"}
// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple, IsFatal 1"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple, PtrOK|IsFatal 1"}
// +validateTrue="T0, string payload"
// +validateTrue={"flags":[], "msg":"TMultiple, no flags 2"}
// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple, PtrOK 2"}
// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple, IsFatal 2"}
// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple, PtrOK|IsFatal 2"}
type TMultiple struct {
	TypeMeta int
	// +validateTrue={"flags":[], "msg":"TMultiple.S, no flags 1"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.S, PtrOK 1"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.S, IsFatal 1"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.S, PtrOK|IsFatal 1"}
	// +validateTrue="T0, string payload"
	// +validateTrue={"flags":[], "msg":"TMultiple.S, no flags 2"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.S, PtrOK 2"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.S, IsFatal 2"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.S, PtrOK|IsFatal 2"}
	S string `json:"s"`
	// +validateTrue={"flags":[], "msg":"TMultiple.PS, no flags 1"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.PS, PtrOK 1"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.PS, IsFatal 1"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.PS, PtrOK|IsFatal 1"}
	// +validateTrue="T0, string payload"
	// +validateTrue={"flags":[], "msg":"TMultiple.PS, no flags 2"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.PS, PtrOK 2"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.PS, IsFatal 2"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.PS, PtrOK|IsFatal 2"}
	PS *string `json:"ps"`
	// +validateTrue={"flags":[], "msg":"TMultiple.T, no flags 1"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.T, PtrOK 1"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.T, IsFatal 1"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.T, PtrOK|IsFatal 1"}
	// +validateTrue="T0, string payload"
	// +validateTrue={"flags":[], "msg":"TMultiple.T, no flags 2"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.T, PtrOK 2"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.T, IsFatal 2"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.T, PtrOK|IsFatal 2"}
	T Tother `json:"t"`
	// +validateTrue={"flags":[], "msg":"TMultiple.PT, no flags 1"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.PT, PtrOK 1"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.PT, IsFatal 1"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.PT, PtrOK|IsFatal 1"}
	// +validateTrue="T0, string payload"
	// +validateTrue={"flags":[], "msg":"TMultiple.PT, no flags 2"}
	// +validateTrue={"flags":["PtrOK"], "msg":"TMultiple.PT, PtrOK 2"}
	// +validateTrue={"flags":["IsFatal"], "msg":"TMultiple.PT, IsFatal 2"}
	// +validateTrue={"flags":["PtrOK", "IsFatal"], "msg":"TMultiple.PT, PtrOK|IsFatal 2"}
	PT *Tother `json:"pt"`
}
