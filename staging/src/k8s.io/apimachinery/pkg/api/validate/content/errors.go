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
	"fmt"
)

// TooLongError returns a string explanation of a "too long" validation
// failure.
func TooLongError(maxChars int) string {
	return fmt.Sprintf("must be no more than %d characters", maxChars)
}

// TooManyError returns a string explanation of a "too many" validation
// failure.
func TooManyError(maxItems int) string {
	return fmt.Sprintf("must contain no more than %d items", maxItems)
}

// RegexError returns a string explanation of a regex validation failure.
func RegexError(msg string, re string, examples ...string) string {
	if len(examples) == 0 {
		msg += " ("
	} else {
		msg += " (e.g. "
		for i := range examples {
			msg += "'" + examples[i] + "'"
			if i < len(examples)-1 {
				msg += ", "
			}
		}
		msg += "; "
	}
	msg += "regex: '" + re + "')"
	return msg
}

// InclusiveRangeError returns a string explanation of a numeric "must be
// between" validation failure.
func InclusiveRangeError(lo, hi int) string {
	return fmt.Sprintf(`must be between %d and %d, inclusive`, lo, hi)
}
