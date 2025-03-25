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

// Package output_test contains test types for deepequal-gen
package output_test

import "k8s.io/utils/ptr"

// TypedefString is an "alias" type.
type TypedefString string

// Struct is a test type that includes OtherStruct
type Struct struct {
	StringField    string
	StringPtrField *string
	IntField       int
	IntPtrField    *int
	BoolField      bool
	BoolPtrField   *bool
	FloatField     float64
	FloatPtrField  *float64
	StructField    OtherStruct
	StructPtrField *OtherStruct
	RecursiveField *Struct

	/*
	   OtherPkgField          other.StructWithoutMethod
	   OtherPkgFieldPtr       *other.StructWithoutMethod
	   OtherPkgMethodField    other.StructWithMethod
	   OtherPkgMethodFieldPtr *other.StructWithMethod

	   SliceStringField    []string
	   SliceStringPtrField []*string
	   SliceIntField       []int
	   SliceIntPtrField    []*int
	   SliceBoolField      []bool
	   SliceBoolPtrField   []*bool
	   SliceFloatField     []float64
	   SliceFloatPtrField  []*float64
	   SliceStructField    []OtherStruct
	   SliceStructPtrField []*OtherStruct
	   SliceRecursiveField []*Struct

	   MapStringField    map[string]string
	   MapStringPtrField map[string]*string
	   MapIntField       map[string]int
	   MapIntPtrField    map[string]*int
	   MapBoolField      map[string]bool
	   MapBoolPtrField   map[string]*bool
	   MapFloatField     map[string]float64
	   MapFloatPtrField  map[string]*float64
	   MapStructField    map[string]OtherStruct
	   MapStructPtrField map[string]*OtherStruct
	   MapRecursiveField map[string]*Struct

	   SliceSliceStringField    [][]string
	   SliceSliceStringPtrField [][]*string
	   SliceMapStringField      []map[string]string
	   SliceMapStringPtrField   []map[string]*string
	   MapSliceStringField      map[string][]string
	   MapSliceStringPtrField   map[string][]*string
	   MapMapStringField        map[string]map[string]string
	   MapMapStringPtrField     map[string]map[string]*string

	   SliceSliceStructField    [][]OtherStruct
	   SliceSliceStructPtrField [][]*OtherStruct
	   SliceMapStructField      []map[string]OtherStruct
	   SliceMapStructPtrField   []map[string]*OtherStruct
	   MapSliceStructField      map[string][]OtherStruct
	   MapSliceStructPtrField   map[string][]*OtherStruct
	   MapMapStructField        map[string]map[string]OtherStruct
	   MapMapStructPtrField     map[string]map[string]*OtherStruct
	*/
}

// OtherStruct is a test type with basic fields
type OtherStruct struct {
	StringField    string
	StringPtrField *string
	IntField       int
	IntPtrField    *int
	BoolField      bool
	BoolPtrField   *bool
	FloatField     float64
	FloatPtrField  *float64

	/*
	   SliceStringField    []string
	   SliceStringPtrField []*string
	   SliceIntField       []int
	   SliceIntPtrField    []*int
	   SliceBoolField      []bool
	   SliceBoolPtrField   []*bool
	   SliceFloatField     []float64
	   SliceFloatPtrField  []*float64

	   MapStringField    map[string]string
	   MapStringPtrField map[string]*string
	   MapIntField       map[string]int
	   MapIntPtrField    map[string]*int
	   MapBoolField      map[string]bool
	   MapBoolPtrField   map[string]*bool
	   MapFloatField     map[string]float64
	   MapFloatPtrField  map[string]*float64

	   SliceSliceStringField    [][]string
	   SliceSliceStringPtrField [][]*string
	   SliceMapStringField      []map[string]string
	   SliceMapStringPtrField   []map[string]*string
	   MapSliceStringField      map[string][]string
	   MapSliceStringPtrField   map[string][]*string
	   MapMapStringField        map[string]map[string]string
	   MapMapStringPtrField     map[string]map[string]*string
	*/
}

// ChangePointers creates new pointers for all pointer fields, copying the values
func (s *Struct) ChangePointers() {
	if s.StringPtrField != nil {
		s.StringPtrField = ptr.To(*s.StringPtrField)
	}
	if s.IntPtrField != nil {
		s.IntPtrField = ptr.To(*s.IntPtrField)
	}
	if s.BoolPtrField != nil {
		s.BoolPtrField = ptr.To(*s.BoolPtrField)
	}
	if s.FloatPtrField != nil {
		s.FloatPtrField = ptr.To(*s.FloatPtrField)
	}
	if s.StructPtrField != nil {
		s.StructPtrField = ptr.To(*s.StructPtrField)
	}
	if s.RecursiveField != nil {
		s.RecursiveField = ptr.To(*s.RecursiveField)
	}
	/*
		if s.OtherPkgFieldPtr != nil {
			s.OtherPkgFieldPtr = ptr.To(*s.OtherPkgFieldPtr)
		}
		if s.OtherPkgMethodFieldPtr != nil {
			s.OtherPkgMethodFieldPtr = ptr.To(*s.OtherPkgMethodFieldPtr)
		}

		// Handle slice pointer fields
		if s.SliceStringPtrField != nil {
			newSlice := make([]*string, len(s.SliceStringPtrField))
			for i, p := range s.SliceStringPtrField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceStringPtrField = newSlice
		}
		if s.SliceIntPtrField != nil {
			newSlice := make([]*int, len(s.SliceIntPtrField))
			for i, p := range s.SliceIntPtrField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceIntPtrField = newSlice
		}
		if s.SliceBoolPtrField != nil {
			newSlice := make([]*bool, len(s.SliceBoolPtrField))
			for i, p := range s.SliceBoolPtrField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceBoolPtrField = newSlice
		}
		if s.SliceFloatPtrField != nil {
			newSlice := make([]*float64, len(s.SliceFloatPtrField))
			for i, p := range s.SliceFloatPtrField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceFloatPtrField = newSlice
		}
		if s.SliceStructPtrField != nil {
			newSlice := make([]*OtherStruct, len(s.SliceStructPtrField))
			for i, p := range s.SliceStructPtrField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceStructPtrField = newSlice
		}
		if s.SliceRecursiveField != nil {
			newSlice := make([]*Struct, len(s.SliceRecursiveField))
			for i, p := range s.SliceRecursiveField {
				if p != nil {
					newSlice[i] = ptr.To(*p)
				}
			}
			s.SliceRecursiveField = newSlice
		}

		// Handle map pointer fields
		if s.MapStringPtrField != nil {
			newMap := make(map[string]*string, len(s.MapStringPtrField))
			for k, v := range s.MapStringPtrField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapStringPtrField = newMap
		}
		if s.MapIntPtrField != nil {
			newMap := make(map[string]*int, len(s.MapIntPtrField))
			for k, v := range s.MapIntPtrField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapIntPtrField = newMap
		}
		if s.MapBoolPtrField != nil {
			newMap := make(map[string]*bool, len(s.MapBoolPtrField))
			for k, v := range s.MapBoolPtrField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapBoolPtrField = newMap
		}
		if s.MapFloatPtrField != nil {
			newMap := make(map[string]*float64, len(s.MapFloatPtrField))
			for k, v := range s.MapFloatPtrField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapFloatPtrField = newMap
		}
		if s.MapStructPtrField != nil {
			newMap := make(map[string]*OtherStruct, len(s.MapStructPtrField))
			for k, v := range s.MapStructPtrField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapStructPtrField = newMap
		}
		if s.MapRecursiveField != nil {
			newMap := make(map[string]*Struct, len(s.MapRecursiveField))
			for k, v := range s.MapRecursiveField {
				if v != nil {
					newMap[k] = ptr.To(*v)
				}
			}
			s.MapRecursiveField = newMap
		}

		// Handle nested pointer fields
		if s.SliceSliceStringPtrField != nil {
			newSlice := make([][]*string, len(s.SliceSliceStringPtrField))
			for i, slice := range s.SliceSliceStringPtrField {
				if slice != nil {
					newInnerSlice := make([]*string, len(slice))
					for j, p := range slice {
						if p != nil {
							newInnerSlice[j] = ptr.To(*p)
						}
					}
					newSlice[i] = newInnerSlice
				}
			}
			s.SliceSliceStringPtrField = newSlice
		}
		if s.SliceMapStringPtrField != nil {
			newSlice := make([]map[string]*string, len(s.SliceMapStringPtrField))
			for i, m := range s.SliceMapStringPtrField {
				if m != nil {
					newMap := make(map[string]*string, len(m))
					for k, v := range m {
						if v != nil {
							newMap[k] = ptr.To(*v)
						}
					}
					newSlice[i] = newMap
				}
			}
			s.SliceMapStringPtrField = newSlice
		}
		if s.MapSliceStringPtrField != nil {
			newMap := make(map[string][]*string, len(s.MapSliceStringPtrField))
			for k, slice := range s.MapSliceStringPtrField {
				if slice != nil {
					newSlice := make([]*string, len(slice))
					for i, p := range slice {
						if p != nil {
							newSlice[i] = ptr.To(*p)
						}
					}
					newMap[k] = newSlice
				}
			}
			s.MapSliceStringPtrField = newMap
		}
		if s.MapMapStringPtrField != nil {
			newMap := make(map[string]map[string]*string, len(s.MapMapStringPtrField))
			for k, m := range s.MapMapStringPtrField {
				if m != nil {
					newInnerMap := make(map[string]*string, len(m))
					for k2, v := range m {
						if v != nil {
							newInnerMap[k2] = ptr.To(*v)
						}
					}
					newMap[k] = newInnerMap
				}
			}
			s.MapMapStringPtrField = newMap
		}

		// Handle nested struct pointer fields
		if s.SliceSliceStructPtrField != nil {
			newSlice := make([][]*OtherStruct, len(s.SliceSliceStructPtrField))
			for i, slice := range s.SliceSliceStructPtrField {
				if slice != nil {
					newInnerSlice := make([]*OtherStruct, len(slice))
					for j, p := range slice {
						if p != nil {
							newInnerSlice[j] = ptr.To(*p)
						}
					}
					newSlice[i] = newInnerSlice
				}
			}
			s.SliceSliceStructPtrField = newSlice
		}
		if s.SliceMapStructPtrField != nil {
			newSlice := make([]map[string]*OtherStruct, len(s.SliceMapStructPtrField))
			for i, m := range s.SliceMapStructPtrField {
				if m != nil {
					newMap := make(map[string]*OtherStruct, len(m))
					for k, v := range m {
						if v != nil {
							newMap[k] = ptr.To(*v)
						}
					}
					newSlice[i] = newMap
				}
			}
			s.SliceMapStructPtrField = newSlice
		}
		if s.MapSliceStructPtrField != nil {
			newMap := make(map[string][]*OtherStruct, len(s.MapSliceStructPtrField))
			for k, slice := range s.MapSliceStructPtrField {
				if slice != nil {
					newSlice := make([]*OtherStruct, len(slice))
					for i, p := range slice {
						if p != nil {
							newSlice[i] = ptr.To(*p)
						}
					}
					newMap[k] = newSlice
				}
			}
			s.MapSliceStructPtrField = newMap
		}
		if s.MapMapStructPtrField != nil {
			newMap := make(map[string]map[string]*OtherStruct, len(s.MapMapStructPtrField))
			for k, m := range s.MapMapStructPtrField {
				if m != nil {
					newInnerMap := make(map[string]*OtherStruct, len(m))
					for k2, v := range m {
						if v != nil {
							newInnerMap[k2] = ptr.To(*v)
						}
					}
					newMap[k] = newInnerMap
				}
			}
			s.MapMapStructPtrField = newMap
		}
	*/
}
