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
)

const qnameCharFmt string = "[A-Za-z0-9]"
const qnameExtCharFmt string = "[-A-Za-z0-9_.]"
const qualifiedNameFmt string = "(" + qnameCharFmt + qnameExtCharFmt + "*)?" + qnameCharFmt
const qualifiedNameMaxLength int = 63

var qualifiedNameRegexp = regexp.MustCompile("^" + qualifiedNameFmt + "$")

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
	default:
		return false, append(errs, RegexError(qualifiedNameFmt, "MyName", "my.name", "123-abc")+
			" with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName'")
	}

	if len(name) == 0 {
		errs = append(errs, "name part "+EmptyError())
	} else if len(name) > qualifiedNameMaxLength {
		errs = append(errs, "name part "+MaxLenError(qualifiedNameMaxLength))
	}
	if !qualifiedNameRegexp.MatchString(name) {
		errs = append(errs, "name part "+RegexError(qualifiedNameFmt, "MyName", "my.name", "123-abc"))
	}
	return len(errs) == 0, errs
}

const labelValueFmt string = "(" + qualifiedNameFmt + ")?"
const LabelValueMaxLength int = 63

var labelValueRegexp = regexp.MustCompile("^" + labelValueFmt + "$")

func IsLabelValue(value string) (bool, []string) {
	var errs []string
	if len(value) > LabelValueMaxLength {
		errs = append(errs, MaxLenError(LabelValueMaxLength))
	}
	if !labelValueRegexp.MatchString(value) {
		errs = append(errs, RegexError(labelValueFmt, "MyValue", "my_value", "12345"))
	}
	return len(errs) == 0, errs
}

const DNS1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const DNS1123LabelMaxLength int = 63

var dns1123LabelRegexp = regexp.MustCompile("^" + DNS1123LabelFmt + "$")

// IsDNS1123Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1123).
func IsDNS1123Label(value string) (bool, []string) {
	var errs []string
	if len(value) > DNS1123LabelMaxLength {
		errs = append(errs, MaxLenError(DNS1123LabelMaxLength))
	}
	if !dns1123LabelRegexp.MatchString(value) {
		errs = append(errs, RegexError(DNS1123LabelFmt, "my-name", "123-abc"))
	}
	return len(errs) == 0, errs
}

const DNS1123SubdomainFmt string = DNS1123LabelFmt + "(\\." + DNS1123LabelFmt + ")*"
const DNS1123SubdomainMaxLength int = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + DNS1123SubdomainFmt + "$")

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) (bool, []string) {
	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !dns1123SubdomainRegexp.MatchString(value) {
		errs = append(errs, RegexError(DNS1123SubdomainFmt, "example.com"))
	}
	return len(errs) == 0, errs
}

const DNS952LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"
const DNS952LabelMaxLength int = 24

var dns952LabelRegexp = regexp.MustCompile("^" + DNS952LabelFmt + "$")

// IsDNS952Label tests for a string that conforms to the definition of a label in
// DNS (RFC 952).
func IsDNS952Label(value string) (bool, []string) {
	var errs []string
	if len(value) > DNS952LabelMaxLength {
		errs = append(errs, MaxLenError(DNS952LabelMaxLength))
	}
	if !dns952LabelRegexp.MatchString(value) {
		errs = append(errs, RegexError(DNS952LabelFmt, "my-name", "abc-123"))
	}
	return len(errs) == 0, errs
}

const CIdentifierFmt string = "[A-Za-z_][A-Za-z0-9_]*"

var cIdentifierRegexp = regexp.MustCompile("^" + CIdentifierFmt + "$")

// IsCIdentifier tests for a string that conforms the definition of an identifier
// in C. This checks the format, but not the length.
func IsCIdentifier(value string) (bool, []string) {
	if !cIdentifierRegexp.MatchString(value) {
		return false, []string{RegexError(CIdentifierFmt, "my_name", "MY_NAME", "MyName")}
	}
	return true, nil
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

var portNameCharsetRegex = regexp.MustCompile("^[-a-z0-9]+$")
var portNameOneLetterRegexp = regexp.MustCompile("[a-z]")

// IsPortName check that the argument is valid syntax. It must be
// non-empty and no more than 15 characters long. It may contain only [-a-z0-9]
// and must contain at least one letter [a-z]. It must not start or end with a
// hyphen, nor contain adjacent hyphens.
//
// Note: We only allow lower-case characters, even though RFC 6335 is case
// insensitive.
func IsPortName(port string) (bool, []string) {
	var errs []string
	if len(port) > 15 {
		errs = append(errs, MaxLenError(15))
	}
	if !portNameCharsetRegex.MatchString(port) {
		errs = append(errs, "must contain only alpha-numeric characters (a-z, 0-9), and hyphens (-)")
	}
	if !portNameOneLetterRegexp.MatchString(port) {
		errs = append(errs, "must contain at least one letter (a-z)")
	}
	if strings.Contains(port, "--") {
		errs = append(errs, "must not contain consecutive hyphens")
	}
	if len(port) > 0 && (port[0] == '-' || port[len(port)-1] == '-') {
		errs = append(errs, "must not begin or end with a hyphen")
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

const percentFmt string = "[0-9]+%"

var percentRegexp = regexp.MustCompile("^" + percentFmt + "$")

func IsPercent(percent string) (bool, []string) {
	if !percentRegexp.MatchString(percent) {
		return false, []string{RegexError(percentFmt, "1%", "93%")}
	}
	return true, nil
}

const httpHeaderNameFmt string = "[-A-Za-z0-9]+"

var httpHeaderNameRegexp = regexp.MustCompile("^" + httpHeaderNameFmt + "$")

// IsHTTPHeaderName checks that a string conforms to the Go HTTP library's
// definition of a valid header field name (a stricter subset than RFC7230).
func IsHTTPHeaderName(value string) (bool, []string) {
	if !httpHeaderNameRegexp.MatchString(value) {
		return false, []string{RegexError(httpHeaderNameFmt, "X-Header-Name")}
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

// MaxLenError returns a string explanation of a "string too long" validation
// failure.
func MaxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

// RegexError returns a string explanation of a regex validation failure.
func RegexError(fmt string, examples ...string) string {
	s := "must match the regex " + fmt
	if len(examples) == 0 {
		return s
	}
	s += " (e.g. "
	for i := range examples {
		if i > 0 {
			s += " or "
		}
		s += "'" + examples[i] + "'"
	}
	return s + ")"
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
