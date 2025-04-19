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

package main

import (
	"fmt"
	"sort"
	"strings"

	"k8s.io/gengo/v2/types"
	"k8s.io/klog/v2"
)

// Linter is a struct that holds the state of the linting process.
// It contains a map of types that have been linted, a list of linting rules,
// and a list of errors that occurred during the linting process.
type Linter struct {
	linted map[*types.Type]bool
	rules  []lintRule
	// Errors is all the errors, grouped by type, that occurred during the
	// linting process.
	Errors map[*types.Type][]error
}

// lintRule is a function that validates a slice of comments.
// It returns a string as an error message if the comments are invalid,
// and an error there is an error happened during the linting process.
type lintRule func(comments []string) (string, error)

func (l *Linter) AddError(t *types.Type, field, msg string) {
	var err error
	if field == "" {
		err = fmt.Errorf("%s", msg)
	} else {
		err = fmt.Errorf("field %s: %s", field, msg)
	}
	l.Errors[t] = append(l.Errors[t], err)
}

func NewLinter(rules ...lintRule) *Linter {
	if len(rules) == 0 {
		rules = defaultLintRules
	}
	return &Linter{
		linted: make(map[*types.Type]bool),
		rules:  rules,
		Errors: map[*types.Type][]error{},
	}
}

func (l *Linter) LintType(t *types.Type) error {
	if _, ok := l.linted[t]; ok {
		return nil
	}
	l.linted[t] = true

	if t.CommentLines != nil {
		klog.V(5).Infof("linting type %s", t.Name.String())
		lintErrs, err := l.lintComments(t.CommentLines)
		if err != nil {
			return err
		}
		for _, lintErr := range lintErrs {
			l.AddError(t, "", lintErr)
		}
	}
	switch t.Kind {
	case types.Alias:
		// Recursively lint the underlying type of the alias.
		if err := l.LintType(t.Underlying); err != nil {
			return err
		}
	case types.Struct:
		// Recursively lint each member of the struct.
		for _, member := range t.Members {
			klog.V(5).Infof("linting comments for field %s of type %s", member.String(), t.Name.String())
			lintErrs, err := l.lintComments(member.CommentLines)
			if err != nil {
				return err
			}
			for _, lintErr := range lintErrs {
				l.AddError(t, member.Name, lintErr)
			}
			if err := l.LintType(member.Type); err != nil {
				return err
			}
		}
	case types.Slice, types.Array, types.Pointer:
		// Recursively lint the element type of the slice or array.
		if err := l.LintType(t.Elem); err != nil {
			return err
		}
	case types.Map:
		// Recursively lint the key and element types of the map.
		if err := l.LintType(t.Key); err != nil {
			return err
		}
		if err := l.LintType(t.Elem); err != nil {
			return err
		}
	}
	return nil
}

// lintComments runs all registered rules on a slice of comments.
func (l *Linter) lintComments(comments []string) ([]string, error) {
	var lintErrs []string
	for _, rule := range l.rules {
		if msg, err := rule(comments); err != nil {
			return nil, err
		} else if msg != "" {
			lintErrs = append(lintErrs, msg)
		}
	}

	return lintErrs, nil
}

// conflictingTagsRule creates a lintRule which checks for conflicting tags.
func conflictingTagsRule(msg string, tags ...string) lintRule {
	if len(tags) < 2 {
		panic("conflictingTagsRule: at least 2 tags must be specified")
	}

	return func(comments []string) (string, error) {
		found := make(map[string]bool)
		for _, comment := range comments {
			for _, tag := range tags {
				if strings.HasPrefix(comment, tag) {
					found[tag] = true
				}
			}
		}
		if len(found) > 1 {
			keys := make([]string, 0, len(found))
			for k := range found {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			return fmt.Sprintf("conflicting tags: {%s}: %s", strings.Join(keys, ", "), msg), nil
		}
		return "", nil
	}
}
