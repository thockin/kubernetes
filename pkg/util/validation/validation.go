/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package validation

import (
	"fmt"
	"math"
	"net"
	"regexp"
	"strings"
	"unicode"
)

const qualifiedNameMaxLength int = 63

func IsQualifiedName(value string) (bool, []string) {
	var errs []string
	parts := strings.Split(value, "/")
	var name string
	switch len(parts) {
	case 1:
		name = parts[0]
	case 2:
		var prefix string
		prefix, name = parts[0], parts[1]
		if len(prefix) == 0 {
			errs = append(errs, "prefix part "+EmptyError())
		} else if ok, msgs := IsDNS1123Subdomain(prefix); !ok {
			errs = append(errs, prefixEach(msgs, "prefix part ")...)
		}
	case 0:
		return false, append(errs, EmptyError())
	default:
		return false, append(errs, "must not contain more than 1 slash (/)")
	}

	errs = append(errs, prefixEach(checkQualifiedNamePart(name), "name part ")...)

	return len(errs) == 0, errs
}

func checkQualifiedNamePart(name string) []string {
	var errs []string
	runes := []rune(name)
	if len(runes) == 0 {
		return append(errs, EmptyError())
	}
	if len(runes) > qualifiedNameMaxLength {
		errs = append(errs, MaxLenError(qualifiedNameMaxLength))
	}
	if !isAlphaNum(runes[0]) || !isAlphaNum(runes[len(runes)-1]) {
		errs = append(errs, "must begin and end with an alpha-numeric character (A-Z, a-z, 0-9)")
	}
	if len(runes) > 2 { // don't repeat errors for first and last chars
		for _, r := range runes[1 : len(runes)-1] {
			if !isAlphaNum(r) && r != '.' && r != '-' && r != '_' {
				errs = append(errs, "must contain only alpha-numeric characters (A-Z, a-z, 0-9), hyphens (-), underscores (_), and dots (.)")
				break
			}
		}
	}
	return errs
}

func IsLabelValue(value string) (bool, []string) {
	var errs []string
	if len(value) == 0 {
		return true, nil
	}
	errs = append(errs, checkQualifiedNamePart(value)...)
	return len(errs) == 0, errs
}

const DNS1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const DNS1123LabelMaxLength int = 63

// IsDNS1123Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1123).
func IsDNS1123Label(value string) (bool, []string) {
	var errs []string
	runes := []rune(value)
	if len(runes) == 0 {
		return false, append(errs, EmptyError())
	}
	if len(runes) > DNS1123LabelMaxLength {
		errs = append(errs, MaxLenError(DNS1123LabelMaxLength))
	}
	if !isLowerAlphaNum(runes[0]) || !isLowerAlphaNum(runes[len(runes)-1]) {
		errs = append(errs, "must begin and end with an alpha-numeric character (a-z, 0-9)")
	}
	if len(runes) > 2 { // don't repeat errors for first and last chars
		for _, r := range runes[1 : len(runes)-1] {
			if !isDNSInnerChar(r) {
				errs = append(errs, "must contain only alpha-numeric characters (a-z, 0-9) and hyphens (-)")
				break
			}
		}
	}
	return len(errs) == 0, errs
}

func isDNSInnerChar(r rune) bool {
	return isLowerAlphaNum(r) || r == '-'
}

const DNS1123SubdomainFmt string = DNS1123LabelFmt + "(\\." + DNS1123LabelFmt + ")*"
const DNS1123SubdomainMaxLength int = 253

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) (bool, []string) {
	var errs []string
	runes := []rune(value)
	if len(runes) == 0 {
		return false, append(errs, EmptyError())
	}
	if len(runes) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	parts := strings.Split(value, ".")
	emptyErr, beginEndErr, contentErr := false, false, false
	for _, part := range parts {
		runes := []rune(part)
		if len(runes) == 0 {
			if !emptyErr {
				errs = append(errs, "parts "+EmptyError())
				emptyErr = true
			}
			continue
		}
		if !isLowerAlphaNum(runes[0]) || !isLowerAlphaNum(runes[len(runes)-1]) {
			if !beginEndErr {
				errs = append(errs, "each part must begin and end with an alpha-numeric character (a-z, 0-9)")
				beginEndErr = true
			}
		}
		if len(runes) > 2 { // don't repeat errors for first and last chars
			for _, r := range runes[1 : len(runes)-1] {
				if !isDNSInnerChar(r) {
					if !contentErr {
						errs = append(errs, "each part must contain only alpha-numeric characters (a-z, 0-9) and hyphens (-); parts must be joined by dots (.)")
						contentErr = true
					}
					break
				}
			}
		}
	}
	return len(errs) == 0, errs
}

const DNS952LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"
const DNS952LabelMaxLength int = 24

// IsDNS952Label tests for a string that conforms to the definition of a label in
// DNS (RFC 952).
func IsDNS952Label(value string) (bool, []string) {
	var errs []string
	runes := []rune(value)
	if len(runes) == 0 {
		return false, append(errs, EmptyError())
	}
	if len(runes) > DNS952LabelMaxLength {
		errs = append(errs, MaxLenError(DNS952LabelMaxLength))
	}
	if !isLowerAlpha(runes[0]) {
		errs = append(errs, "must begin with an alphabetic character (a-z)")
	}
	if !isLowerAlphaNum(runes[len(runes)-1]) {
		errs = append(errs, "must end with an alpha-numeric character (a-z, 0-9)")
	}
	if len(runes) > 2 { // don't repeat errors for first and last chars
		for _, r := range runes[1 : len(runes)-1] {
			if !isDNSInnerChar(r) {
				errs = append(errs, "must contain only alpha-numeric characters (a-z, 0-9) and hyphens (-)")
				break
			}
		}
	}
	return len(errs) == 0, errs
}

// IsCIdentifier tests for a string that conforms the definition of an identifier
// in C. This checks the format, but not the length.
func IsCIdentifier(value string) (bool, []string) {
	var errs []string
	runes := []rune(value)
	if len(runes) == 0 {
		return false, append(errs, EmptyError())
	}
	if !isASCIILetter(runes[0]) && runes[0] != '_' {
		errs = append(errs, "must start with an alphabetic character (a-z, A-Z) or underscore (_)")
	}
	if len(runes) > 1 { // don't repeat errors for first char
		for _, r := range runes[1:len(runes)] {
			if !isAlphaNum(r) && r != '_' {
				errs = append(errs, "must contain only alpha-numeric characters (a-z, A-Z, 0-9) and underscores (_)")
				break
			}
		}
	}
	return len(errs) == 0, errs
}

// IsPortNum tests that the argument is a valid, non-zero port number.
func IsPortNum(port int) (bool, []string) {
	return IsBetweenInclusive(1, 65535, int64(port))
}

// Now in libcontainer UID/GID limits is 0 ~ 1<<31 - 1
// TODO: once we have a type for UID/GID we should make these that type.
const (
	minUserID  = 0
	maxUserID  = math.MaxInt32
	minGroupID = 0
	maxGroupID = math.MaxInt32
)

// IsGroupID tests that the argument is a valid Unix GID.
func IsGroupID(gid int64) (bool, []string) {
	return IsBetweenInclusive(minGroupID, maxGroupID, int64(gid))
}

// IsUserID tests that the argument is a valid Unix UID.
func IsUserID(uid int64) (bool, []string) {
	return IsBetweenInclusive(minUserID, maxUserID, int64(uid))
}

// IsPortName check that the argument is valid syntax. It must be
// non-empty and no more than 15 characters long. It may contain only [-a-z0-9]
// and must contain at least one letter [a-z]. It must not start or end with a
// hyphen, nor contain adjacent hyphens.
//
// Note: We only allow lower-case characters, even though RFC 6335 is case
// insensitive.
func IsPortName(port string) (bool, []string) {
	var errs []string
	runes := []rune(port)
	if len(runes) == 0 {
		return false, append(errs, EmptyError())
	}
	if len(runes) > 15 {
		errs = append(errs, MaxLenError(15))
	}
	if runes[0] == '-' || runes[len(runes)-1] == '-' {
		errs = append(errs, "must not begin or end with a hyphen")
	}
	foundLetter := false // must contain at least one letter
	if isASCIILetter(runes[0]) || isASCIILetter(runes[len(runes)-1]) {
		foundLetter = true
	}
	if len(runes) > 2 { // don't repeat errors for first and last chars
		for _, r := range runes[1 : len(runes)-1] {
			if !isLowerAlphaNum(r) && r != '-' {
				errs = append(errs, "must contain only alpha-numeric characters (a-z, 0-9), and hyphens (-)")
				break
			}
			if isASCIILetter(r) {
				foundLetter = true
			}
		}
	}
	if !foundLetter {
		errs = append(errs, "must contain at least one letter (a-z)")
	}
	if strings.Contains(port, "--") {
		errs = append(errs, "must not contain consecutive hyphens")
	}
	return len(errs) == 0, errs
}

func isIPv4(value string) (net.IP, []string) {
	ip := net.ParseIP(value)
	if ip == nil || ip.To4() == nil {
		return nil, []string{"must be a valid IPv4 address, (e.g. 10.9.8.7)"}
	}
	return ip, nil
}

// IsIPv4 tests that the argument is a valid IPv4 address.
func IsIPv4(value string) (bool, []string) {
	ip, msgs := isIPv4(value)
	if ip == nil {
		return false, msgs
	}
	return true, nil
}

// Is NonSpecialIPv4 tests that the argument is a "normal" IP.  Specifically:
//   * valid IPv4 address
//   * not 0.0.0.0
//   * not a loopback address
//   * not a link-local address
func IsNonSpecialIPv4(value string) (bool, []string) {
	ip, msgs := isIPv4(value)
	if ip == nil {
		return false, msgs
	}
	if ip.IsUnspecified() {
		return false, []string{"may not be unspecified (0.0.0.0)"}
	}
	if ip.IsLoopback() {
		return false, []string{"may not be in the loopback range (127.0.0.0/8)"}
	}
	if ip.IsLinkLocalUnicast() {
		return false, []string{"may not be in the link-local range (169.254.0.0/16)"}
	}
	if ip.IsLinkLocalMulticast() {
		return false, []string{"may not be in the link-local multicast range (224.0.0.0/24)"}
	}
	return true, nil
}

var percentRegexp = regexp.MustCompile("^[0-9]+%$")

func IsPercent(percent string) (bool, []string) {
	if len(percent) == 0 {
		return false, []string{EmptyError()}
	}
	if !percentRegexp.MatchString(percent) {
		return false, []string{"must be one or more digits (0-9) followed by a percent (%)"}
	}
	return true, nil
}

const httpHeaderNameFmt string = "[-A-Za-z0-9]+"

var httpHeaderNameRegexp = regexp.MustCompile("^" + httpHeaderNameFmt + "$")

// IsHTTPHeaderName checks that a string conforms to the Go HTTP library's
// definition of a valid header field name (a stricter subset than RFC7230).
func IsHTTPHeaderName(value string) (bool, []string) {
	if len(value) == 0 {
		return false, []string{EmptyError()}
	}
	if !httpHeaderNameRegexp.MatchString(value) {
		return false, []string{"must contain only alpha-numeric characters (a-z, A-Z, 0-9) and dashes (-) (e.g. X-Header-Name)"}
	}
	return true, nil
}

func IsLessThan(max, value int64) (bool, []string) {
	if value < max {
		return true, nil
	}
	return false, []string{fmt.Sprintf("must be less than %d", max)}
}

func IsLessThanOrEqual(max, value int64) (bool, []string) {
	if value <= max {
		return true, nil
	}
	return false, []string{fmt.Sprintf("must be less than or equal to %d", max)}
}

func IsGreaterThan(min, value int64) (bool, []string) {
	if value > min {
		return true, nil
	}
	return false, []string{fmt.Sprintf("must be greater than %d", min)}
}

func IsGreaterThanOrEqual(min, value int64) (bool, []string) {
	if value >= min {
		return true, nil
	}
	return false, []string{fmt.Sprintf("must be greater than or equal to %d", min)}
}

func IsBetweenInclusive(lo, hi, value int64) (bool, []string) {
	if lo <= value && value <= hi {
		return true, nil
	}
	return false, []string{fmt.Sprintf("must be between %d and %d, inclusive", lo, hi)}
}

// PathNameMayNotBe specifies strings that cannot be used as names specified as path segments (like the REST API or etcd store)
var PathNameMayNotBe = []string{".", ".."}

// PathNameMayNotContain specifies substrings that cannot be used in names specified as path segments (like the REST API or etcd store)
var PathNameMayNotContain = []string{"/", "%"}

// IsValidPathSegmentName validates the name can be safely encoded as a path segment
func IsValidPathSegmentName(name string) (bool, []string) {
	for _, illegalName := range PathNameMayNotBe {
		if name == illegalName {
			return false, []string{fmt.Sprintf("may not be '%s'", illegalName)}
		}
	}

	return IsValidPathSegmentPrefix(name)
}

// IsValidPathSegmentPrefix validates the name can be used as a prefix for a name which will be encoded as a path segment
// It does not check for exact matches with disallowed names, since an arbitrary suffix might make the name valid
func IsValidPathSegmentPrefix(name string) (bool, []string) {
	for _, illegalContent := range PathNameMayNotContain {
		if strings.Contains(name, illegalContent) {
			return false, []string{fmt.Sprintf("may not contain '%s'", illegalContent)}
		}
	}

	return true, nil
}

// MaxLenError returns a string explanation of a "string too long" validation
// failure.
func MaxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

// EmptyError returns a string explanation of a "must not be empty" validation
// failure.
func EmptyError() string {
	return "must be non-empty"
}

func prefixEach(msgs []string, prefix string) []string {
	for i := range msgs {
		msgs[i] = prefix + msgs[i]
	}
	return msgs
}

func isASCIILetter(r rune) bool {
	return unicode.IsLetter(r) && r <= unicode.MaxASCII
}

func isAlphaNum(r rune) bool {
	return isASCIILetter(r) || unicode.IsDigit(r)
}

func isLowerAlpha(r rune) bool {
	return isASCIILetter(r) && unicode.IsLower(r)
}

func isLowerAlphaNum(r rune) bool {
	return isLowerAlpha(r) || unicode.IsDigit(r)
}
