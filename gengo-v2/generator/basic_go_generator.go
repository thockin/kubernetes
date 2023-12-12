/*
Copyright 2015 The Kubernetes Authors.

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

package generator

import (
	"io"

	"k8s.io/gengo/v2/namer"
	"k8s.io/gengo/v2/types"
)

const (
	GolangFileType = "golang"
)

// BasicGoGenerator implements a nearly do-nothing Generator for .go files.
//
// It can be used to implement static content files, or serve as a base for
// more specialized Generator implementations, by embedding and overriding
// methods.
type BasicGoGenerator struct {
	// OptionalName, if present, will be used for the generator's name, and
	// the filename (with ".go" appended).  If this is not specified, the
	// Filename() method must be overridden.
	OptionalName string

	// OptionalBody, if present, will be used as the return from the "Init"
	// method. This causes it to be static content for the entire file if
	// no other generator touches the file.
	OptionalBody []byte
}

func (d BasicGoGenerator) Name() string                                        { return d.OptionalName }
func (d BasicGoGenerator) Filter(*Context, *types.Type) bool                   { return true }
func (d BasicGoGenerator) Namers(*Context) map[string]namer.Namer              { return nil }
func (d BasicGoGenerator) Imports(*Context) []string                           { return []string{} }
func (d BasicGoGenerator) PackageVars(*Context) []string                       { return []string{} }
func (d BasicGoGenerator) PackageConsts(*Context) []string                     { return []string{} }
func (d BasicGoGenerator) GenerateType(*Context, *types.Type, io.Writer) error { return nil }
func (d BasicGoGenerator) FileType() string                                    { return GolangFileType }
func (d BasicGoGenerator) Finalize(*Context, io.Writer) error                  { return nil }

func (d BasicGoGenerator) Filename() string {
	if len(d.OptionalName) == 0 {
		panic("BasicGoGenerator.Filename was called, but OptionalName was not specified")
	}
	return d.OptionalName + ".go"
}

func (d BasicGoGenerator) Init(c *Context, w io.Writer) error {
	_, err := w.Write(d.OptionalBody)
	return err
}

var (
	_ Generator = BasicGoGenerator{}
)
