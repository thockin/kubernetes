/*
Copyright 2021 The Kubernetes Authors.

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

package validators

import (
	"encoding/json"
	"fmt"
	"slices"

	"k8s.io/gengo/v2"
	"k8s.io/gengo/v2/parser/tags"
	"k8s.io/gengo/v2/types"
)

var discriminatedUnionValidator = types.Name{Package: libValidationPkg, Name: "DiscriminatedUnion"}
var unionValidator = types.Name{Package: libValidationPkg, Name: "Union"}

var newDiscriminatedUnionMembership = types.Name{Package: libValidationPkg, Name: "NewDiscriminatedUnionMembership"}
var newUnionMembership = types.Name{Package: libValidationPkg, Name: "NewUnionMembership"}

func init() {
	AddToRegistry(InitUnionDeclarativeValidator)
}

func InitUnionDeclarativeValidator(cfg *ValidatorConfig) DeclarativeValidator {
	return &unionDeclarativeValidator{
		universe: cfg.GeneratorContext.Universe,
	}
}

type unionDeclarativeValidator struct {
	universe types.Universe
}

const (
	// +union and +unionDiscriminator tag are used by openapi-gen to publish x-kubernetes-union and x-kubernetes-discriminator
	// extensions into Kubernetes published OpenAPI.
	discriminatorTagName = "k8s:unionDiscriminator"
	memberTagName        = "k8s:unionMember"
)

// discriminatorParams defines JSON the parameter value for the +unionDiscriminator tag.
type discriminatorParams struct {
	// Union sets the name of the union this discriminator belongs to.
	// This is only needed for go structs that contain more than one union.
	// Optional.
	Union string `json:"union,omitempty"`
}

// memberParams defines the JSON parameter value for the +unionMember tag.
type memberParams struct {
	// Union sets the name of the union this member belongs to.
	// This is only needed for go structs that contain more than one union.
	// Optional.
	Union string `json:"union,omitempty"`
	// MemberName provides a name for a union member. If the union has a
	// discriminator, the member name must match the value the discriminator
	// is set to when this member is specified.
	// Optional.
	// Defaults to the go field name.
	MemberName string `json:"memberName,omitempty"`
}

// union defines how a union validation will be generated, based
// on +unionMember and +unionDiscriminator tags found in a go struct.
type union struct {
	// fields provides field information about all the members of the union.
	// Each slice element is a [2]string to provide a fieldName and memberName pair, where
	// [0] identifies the field name and [1] identifies the union member Name.
	// fields is index aligned with fieldMembers.
	// If member name is not set, it defaults to the go struct field name.
	fields []any
	// fieldMembers is a list of types.Member for all the members of the union.
	fieldMembers []any

	// discriminator is the name of the discriminator field
	discriminator *string
	// discriminatorMember is the types.Member of the discriminator field.
	discriminatorMember any
}

// unions represents all the unions for a go struct.
type unions map[string]*union

// getOrCreate gets a union by name, or initializes a new union by the given name.
func (us unions) getOrCreate(name string) *union {
	var u *union
	var ok bool
	if u, ok = us[name]; !ok {
		u = &union{}
		us[name] = u
	}
	return u
}

func (c *unionDeclarativeValidator) ExtractValidations(t *types.Type, comments []string) (Validations, error) {
	result := Validations{}
	unions := unions{}
	for _, member := range t.Members {
		commentTags := gengo.ExtractCommentTags("+", member.CommentLines)
		if commentTag, ok := commentTags[memberTagName]; ok {
			if len(commentTag) != 1 {
				return result, fmt.Errorf("must have one %q tag", memberTagName)
			}
			tag := commentTag[0]
			var fieldName string
			jsonTag, ok := tags.LookupJSON(member)
			if !ok {
				return result, fmt.Errorf("field %q is a union member but has no JSON struct field tag", member)
			}
			fieldName = jsonTag.Name
			if len(fieldName) == 0 {
				return result, fmt.Errorf("field %q is a union member but has no JSON name", member)
			}

			p := &memberParams{MemberName: member.Name}
			if len(tag) > 0 {
				// Name may optionally be overridden by tag's memberName field.
				if err := json.Unmarshal([]byte(tag), &p); err != nil {
					return result, fmt.Errorf("error parsing JSON value: %v (%q)", err, tag)
				}
			}
			u := unions.getOrCreate(p.Union)
			u.fields = append(u.fields, [2]string{fieldName, p.MemberName})
			u.fieldMembers = append(u.fieldMembers, member)
		}

		if commentTag, ok := commentTags[discriminatorTagName]; ok {
			if len(commentTag) != 1 {
				return result, fmt.Errorf("must have one %q tag", memberTagName)
			}
			tag := commentTag[0]

			p := &discriminatorParams{}
			if len(tag) > 0 {
				if err := json.Unmarshal([]byte(tag), &p); err != nil {
					return result, fmt.Errorf("error parsing JSON value: %v (%q)", err, tag)
				}
			}
			u := unions.getOrCreate(p.Union)

			var discriminatorFieldName string
			if jsonAnnotation, ok := tags.LookupJSON(member); ok {
				discriminatorFieldName = jsonAnnotation.Name
				u.discriminator = &discriminatorFieldName
				u.discriminatorMember = member
			}
		}
	}

	// Sort the keys for stable output.
	keys := make([]string, 0, len(unions))
	for k := range unions {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, unionName := range keys {
		u := unions[unionName]
		if len(u.fieldMembers) > 0 || u.discriminator != nil {
			// TODO: Avoid the "local" here. This was added to to avoid errors caused when the package is an empty string.
			//       The correct package would be the output package but is not known here. This does not show up in generated code.
			// TODO: Append a consistent hash suffix to avoid generated name conflicts?
			supportVarName := PrivateVar{Name: "UnionMembershipFor" + t.Name.Name + unionName, Package: "local"}
			if u.discriminator != nil {
				supportVar := Variable(supportVarName, Function(memberTagName, DefaultFlags, newDiscriminatedUnionMembership, append([]any{*u.discriminator}, u.fields...)...))
				result.Variables = append(result.Variables, supportVar)
				fn := Function(memberTagName, DefaultFlags, discriminatedUnionValidator, append([]any{supportVarName, u.discriminatorMember}, u.fieldMembers...)...)
				result.Functions = append(result.Functions, fn)
			} else {
				supportVar := Variable(supportVarName, Function(memberTagName, DefaultFlags, newUnionMembership, u.fields...))
				result.Variables = append(result.Variables, supportVar)
				fn := Function(memberTagName, DefaultFlags, unionValidator, append([]any{supportVarName}, u.fieldMembers...)...)
				result.Functions = append(result.Functions, fn)
			}
		}
	}

	return result, nil
}

func (unionDeclarativeValidator) Docs() []TagDoc {
	return []TagDoc{{
		Tag:         discriminatorTagName,
		Description: "Indicates that this field is the discriminator for a union.",
		Contexts:    []TagContext{TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<json-object>",
			Docs:        "",
			Schema: []TagPayloadSchema{{
				Key:   "union",
				Value: "<string>",
				Docs:  "the name of the union, if more than one exists",
			}},
		}},
	}, {
		Tag:         memberTagName,
		Description: "Indicates that this field is a member of a union.",
		Contexts:    []TagContext{TagContextField},
		Payloads: []TagPayloadDoc{{
			Description: "<json-object>",
			Docs:        "",
			Schema: []TagPayloadSchema{{
				Key:   "union",
				Value: "<string>",
				Docs:  "the name of the union, if more than one exists",
			}, {
				Key:     "memberName",
				Value:   "<string>",
				Docs:    "the discriminator value for this member",
				Default: "the field's name",
			}},
		}},
	}}
}
