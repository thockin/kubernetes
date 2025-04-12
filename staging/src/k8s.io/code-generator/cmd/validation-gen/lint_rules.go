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

package main

import "k8s.io/gengo/v2/types"

var ruleOptionalAndRequired = conflictingTagsRule(
	"fields cannot be both optional and required",
	"+k8s:optional", "+k8s:required")

var ruleRequiredAndDefault = conflictingTagsRule(
	"fields with default values are always optional",
	"+k8s:required", "+default")

// FIXME: when lint recurses it triggers this on every field of (e.g.) pod.
// FIXME: hasValiations() is for the whole type, so still triggers on PodTemplate
// FIXME: need to check if this FIELD has validations
// FIXME: could check for +optional OR +k8s:optional?
// FIXME: Also need a nolint tag for output_tests?
var rulePointersOptional = kindRequiresTagRule(
	"pointer fields must be optional",
	types.Pointer, "+optional")

var defaultLintRules = []lintRule{
	ruleOptionalAndRequired,
	ruleRequiredAndDefault,
	rulePointersOptional,
}
