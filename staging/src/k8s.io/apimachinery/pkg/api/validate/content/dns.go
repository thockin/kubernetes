/*
Copyright 2014 The Kubernetes Authors.

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

package content

import (
	"regexp"
	"unicode"
)

const dns1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const dns1123LabelMaxLength int = 63

// DNS1123LabelMaxLength is a DNS label's max length (RFC 1123).
const DNS1123LabelMaxLength int = dns1123LabelMaxLength

var dnsLabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")

// IsDNS1123Label returns error messages if the specified value does not
// parse as per the definition of a label in DNS (approximately RFC 1123).
func IsDNS1123Label(value string) []string {
	var errs []string
	if len(value) > dns1123LabelMaxLength {
		errs = append(errs, MaxLenError(dns1123LabelMaxLength))
	}
	isAlNum := func(r rune) bool {
		if r > unicode.MaxASCII {
			return false
		}
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return true
		}
		if unicode.IsDigit(r) {
			return true
		}
		return false
	}
	if runes := []rune(value); len(runes) == 0 || !isAlNum(runes[0]) || !isAlNum(runes[len(runes)-1]) {
		errs = append(errs, "must start and end with lower-case alphanumeric characters")
	}
	if !dnsLabelRegexp.MatchString(value) {
		errs = append(errs, "must consist of lower-case alphanumeric characters or '-'")
	}
	return errs
}
