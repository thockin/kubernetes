/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package field

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Path represents the path from some root to a particular field.
type Path struct {
	name   string // the name of this field or "" if this is an index
	index  string // if name == "", this is a subscript (index or map key) of the previous element
	parent *Path  // nil if this is the root element
}

// NewPath creates a root Path object.
func NewPath(name string, moreNames ...string) *Path {
	r := &Path{name: name, parent: nil}
	for _, anotherName := range moreNames {
		r = &Path{name: anotherName, parent: r}
	}
	return r
}

// Root returns the root element of this Path.
func (p *Path) Root() *Path {
	for ; p.parent != nil; p = p.parent {
		// Do nothing.
	}
	return p
}

// Child creates a new Path that is a child of the method receiver.
func (p *Path) Child(name string, moreNames ...string) *Path {
	r := NewPath(name, moreNames...)
	r.Root().parent = p
	return r
}

// Index indicates that the previous Path is to be subscripted by an int.
// This sets the same underlying value as Key.
func (p *Path) Index(index int) *Path {
	return &Path{index: strconv.Itoa(index), parent: p}
}

// Key indicates that the previous Path is to be subscripted by a string.
// This sets the same underlying value as Index.
func (p *Path) Key(key string) *Path {
	return &Path{index: key, parent: p}
}

// String produces a string representation of the Path.
func (p *Path) String() string {
	// make a slice to iterate
	elems := []*Path{}
	for ; p != nil; p = p.parent {
		elems = append(elems, p)
	}

	// iterate, but it has to be backwards
	buf := bytes.NewBuffer(nil)
	for i := range elems {
		p := elems[len(elems)-1-i]
		if p.parent != nil && len(p.name) > 0 {
			// This is either the root or it is a subscript.
			buf.WriteString(".")
		}
		if len(p.name) > 0 {
			buf.WriteString(p.name)
		} else {
			fmt.Fprintf(buf, "[%s]", p.index)
		}
	}
	return buf.String()
}

// NotFoundError returns a *Error indicating "value not found".  This is
// used to report failure to find a requested value (e.g. looking up an ID).
func (p *Path) NotFoundError(value interface{}) *Error {
	return &Error{ErrorTypeNotFound, p.String(), value, ""}
}

// RequiredError returns a *Error indicating "value required".  This is used
// to report required values that are not provided (e.g. empty strings, null
// values, or empty arrays).
func (p *Path) RequiredError(detail string) *Error {
	return &Error{ErrorTypeRequired, p.String(), "", detail}
}

// DuplicateError returns a *Error indicating "duplicate value".  This is
// used to report collisions of values that must be unique (e.g. names or IDs).
func (p *Path) DuplicateError(value interface{}) *Error {
	return &Error{ErrorTypeDuplicate, p.String(), value, ""}
}

// InvalidError returns a *Error indicating "invalid value".  This is used
// to report malformed values (e.g. failed regex match, too long, out of bounds).
func (p *Path) InvalidError(value interface{}, detail string) *Error {
	return &Error{ErrorTypeInvalid, p.String(), value, detail}
}

// NotSupportedError returns a *Error indicating "unsupported value".
// This is used to report unknown values for enumerated fields (e.g. a list of
// valid values).
func (p *Path) NotSupportedError(value interface{}, validValues []string) *Error {
	detail := ""
	if validValues != nil && len(validValues) > 0 {
		detail = "supported values: " + strings.Join(validValues, ", ")
	}
	return &Error{ErrorTypeNotSupported, p.String(), value, detail}
}

// ForbiddenError returns a *Error indicating "forbidden".  This is used to
// report valid (as per formatting rules) values which would be accepted under
// some conditions, but which are not permitted by current conditions (e.g.
// security policy).
func (p *Path) ForbiddenError(detail string) *Error {
	return &Error{ErrorTypeForbidden, p.String(), "", detail}
}

// TooLongError returns a *Error indicating "too long".  This is used to
// report that the given value is too long.  This is similar to
// Invalid, but the returned error will not include the too-long
// value.
func (p *Path) TooLongError(value interface{}, maxLength int) *Error {
	return &Error{ErrorTypeTooLong, p.String(), value, fmt.Sprintf("must have at most %d characters", maxLength)}
}

// InternalError returns a *Error indicating "internal error".  This is used
// to signal that an error was found that was not directly related to user
// input.  The err argument must be non-nil.
func (p *Path) InternalError(err error) *Error {
	return &Error{ErrorTypeInternal, p.String(), nil, err.Error()}
}
