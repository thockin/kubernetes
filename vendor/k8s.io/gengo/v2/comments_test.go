package gengo

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractExtendedCommentTags(t *testing.T) {
	mktags := func(t ...Tag) []Tag { return t }
	mkstrs := func(s ...string) []string { return s }

	cases := []struct {
		name     string
		comments []string
		prefixes []string
		expect   map[string][]Tag
	}{{
		name: "no args",
		comments: []string{
			"Human comment that is ignored",
			"+simpleNoVal",
			"+simpleWithVal=val",
			"+duplicateNoVal",
			"+duplicateNoVal",
			"+duplicateWithVal=val1",
			"+duplicateWithVal=val2",
		},
		expect: map[string][]Tag{
			"simpleNoVal":   mktags(Tag{"simpleNoVal", nil, ""}),
			"simpleWithVal": mktags(Tag{"simpleWithVal", nil, "val"}),
			"duplicateNoVal": mktags(
				Tag{"duplicateNoVal", nil, ""},
				Tag{"duplicateNoVal", nil, ""}),
			"duplicateWithVal": mktags(
				Tag{"duplicateWithVal", nil, "val1"},
				Tag{"duplicateWithVal", nil, "val2"}),
		},
	}, {
		name: "empty parens",
		comments: []string{
			"Human comment that is ignored",
			"+simpleNoVal()",
			"+simpleWithVal()=val",
			"+duplicateNoVal()",
			"+duplicateNoVal()",
			"+duplicateWithVal()=val1",
			"+duplicateWithVal()=val2",
		},
		expect: map[string][]Tag{
			"simpleNoVal":   mktags(Tag{"simpleNoVal", nil, ""}),
			"simpleWithVal": mktags(Tag{"simpleWithVal", nil, "val"}),
			"duplicateNoVal": mktags(
				Tag{"duplicateNoVal", nil, ""},
				Tag{"duplicateNoVal", nil, ""}),
			"duplicateWithVal": mktags(
				Tag{"duplicateWithVal", nil, "val1"},
				Tag{"duplicateWithVal", nil, "val2"}),
		},
	}, {
		name: "mixed no args and empty parens",
		comments: []string{
			"Human comment that is ignored",
			"+noVal",
			"+withVal=val1",
			"+noVal()",
			"+withVal()=val2",
		},
		expect: map[string][]Tag{
			"noVal": mktags(
				Tag{"noVal", nil, ""},
				Tag{"noVal", nil, ""}),
			"withVal": mktags(
				Tag{"withVal", nil, "val1"},
				Tag{"withVal", nil, "val2"}),
		},
	}, {
		name: "with args",
		comments: []string{
			"Human comment that is ignored",
			"+simpleNoVal(arg)",
			"+simpleWithVal(arg)=val",
			"+duplicateNoVal(arg1)",
			"+duplicateNoVal(arg2)",
			"+duplicateWithVal(arg1)=val1",
			"+duplicateWithVal(arg2)=val2",
		},
		expect: map[string][]Tag{
			"simpleNoVal":   mktags(Tag{"simpleNoVal", mkstrs("arg"), ""}),
			"simpleWithVal": mktags(Tag{"simpleWithVal", mkstrs("arg"), "val"}),
			"duplicateNoVal": mktags(
				Tag{"duplicateNoVal", mkstrs("arg1"), ""},
				Tag{"duplicateNoVal", mkstrs("arg2"), ""}),
			"duplicateWithVal": mktags(
				Tag{"duplicateWithVal", mkstrs("arg1"), "val1"},
				Tag{"duplicateWithVal", mkstrs("arg2"), "val2"}),
		},
	}, {
		name: "mixed no args and empty parens",
		comments: []string{
			"Human comment that is ignored",
			"+noVal",
			"+withVal=val1",
			"+noVal(arg)",
			"+withVal(arg)=val2",
		},
		expect: map[string][]Tag{
			"noVal": mktags(
				Tag{"noVal", nil, ""},
				Tag{"noVal", mkstrs("arg"), ""}),
			"withVal": mktags(
				Tag{"withVal", nil, "val1"},
				Tag{"withVal", mkstrs("arg"), "val2"}),
		},
	}, {
		name: "prefixes",
		comments: []string{
			"Human comment that is ignored",
			"+pfx1Foo",
			"+pfx2Foo=val1",
			"+pfx3Bar",
			"+pfx4Bar=val",
			"+pfx1Foo(arg)",
			"+pfx2Foo(arg)=val2",
			"+pfx3Bar(arg)",
			"+pfx4Bar(arg)=val",
		},
		prefixes: []string{"pfx1", "pfx2"},
		expect: map[string][]Tag{
			"pfx1Foo": mktags(
				Tag{"pfx1Foo", nil, ""},
				Tag{"pfx1Foo", mkstrs("arg"), ""}),
			"pfx2Foo": mktags(
				Tag{"pfx2Foo", nil, "val1"},
				Tag{"pfx2Foo", mkstrs("arg"), "val2"}),
		},
	}}

	for _, tc := range cases {
		result, _ := ExtractFunctionStyleCommentTags("+", tc.prefixes, tc.comments)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("case %q: wrong result:\n%v", tc.name, cmp.Diff(tc.expect, result))
		}
	}
}

func TestParseTagKey(t *testing.T) {
	mkss := func(s ...string) []string { return s }

	cases := []struct {
		input      string
		expectKey  string
		expectArgs []string
		err        bool
	}{
		{"simple", "simple", nil, false},
		{"parens()", "parens", nil, false},
		{"withArgLower(arg)", "withArgLower", mkss("arg"), false},
		{"withArgUpper(ARG)", "withArgUpper", mkss("ARG"), false},
		{"withArgMixed(ArG)", "withArgMixed", mkss("ArG"), false},
		{"withArgs(arg1, arg2)", "", nil, true},
		{"trailingParen(arg))", "", nil, true},
		{"trailingSpace(arg) ", "", nil, true},
		{"argWithDash(arg-name) ", "", nil, true},
		{"argWithUnder(arg_name) ", "", nil, true},
	}
	for _, tc := range cases {
		key, args, err := parseTagKey(tc.input)
		if err != nil && tc.err == false {
			t.Errorf("[%q]: expected success, got: %v", tc.input, err)
			continue
		}
		if err == nil {
			if tc.err == true {
				t.Errorf("[%q]: expected failure, got: %v(%v)", tc.input, key, args)
				continue
			}
			if key != tc.expectKey {
				t.Errorf("[%q]\nexpected key: %q, got: %q", tc.input, tc.expectKey, key)
			}
			if len(args) != len(tc.expectArgs) {
				t.Errorf("[%q]: expected %d args, got: %q", tc.input, len(tc.expectArgs), args)
				continue
			}
			for i := range tc.expectArgs {
				if want, got := tc.expectArgs[i], args[i]; got != want {
					t.Errorf("[%q]\nexpected %q, got %q", tc.input, want, got)
				}
			}
		}
	}
}

func TestParseTagArgs(t *testing.T) {
	mkss := func(s ...string) []string { return s }

	cases := []struct {
		input  string
		expect []string
		err    bool
	}{
		{")", nil, false},
		{"lower)", mkss("lower"), false},
		{"CAPITAL)", mkss("CAPITAL"), false},
		{"MiXeD)", mkss("MiXeD"), false},
		{"mIxEd)", mkss("mIxEd"), false},
		{"_under)", nil, true},
		{"has space", nil, true},
		{"has-dash", nil, true},
		{`"hasQuotes"`, nil, true},
		{"multiple, args)", nil, true},
		{"noClosingParen", nil, true},
		{"extraParen))", nil, true},
		{"trailingSpace) ", nil, true},
	}
	for _, tc := range cases {
		ret, err := parseTagArgs(tc.input)
		if err != nil && tc.err == false {
			t.Errorf("[%q]: expected success, got: %v", tc.input, err)
			continue
		}
		if err == nil {
			if tc.err == true {
				t.Errorf("[%q]: expected failure, got: %q", tc.input, ret)
				continue
			}
			if len(ret) != len(tc.expect) {
				t.Errorf("[%q]: expected %d results, got: %q", tc.input, len(tc.expect), ret)
				continue
			}
			for i := range tc.expect {
				if want, got := tc.expect[i], ret[i]; got != want {
					t.Errorf("[%q]\nexpected %q, got %q", tc.input, want, got)
				}
			}
		}
	}
}
