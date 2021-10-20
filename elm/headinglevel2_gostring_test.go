package elm_test

import (
	"github.com/reiver/go-gmi/elm"

	"testing"
)

func TestHeadingLevel2_GoString(t *testing.T) {

	tests := []struct{
		Value    string
		EOL      string
		Expected string
	}{
		{
			Value:                       "",
			EOL:                             "\u000A", // line feed
			Expected: `elm.Create("## ", "", "\n")`,
		},
		{
			Value:                       "",
			EOL:                             "\u000B", // vertical tab
			Expected: `elm.Create("## ", "", "\v")`,
		},
		{
			Value:                       "",
			EOL:                             "\u000C", // form feed
			Expected: `elm.Create("## ", "", "\f")`,
		},
		{
			Value:                       "",
			EOL:                             "\u000D", // carriage return
			Expected: `elm.Create("## ", "", "\r")`,
		},
		{
			Value:                       "",
			EOL:                             "\u000D\u000A", // carriage return, line feed
			Expected: `elm.Create("## ", "", "\r\n")`,
		},
		{
			Value:                       "",
			EOL:                             "\u0085", // next line
			Expected: `elm.Create("## ", "", "\u0085")`,
		},
		{
			Value:                       "",
			EOL:                             "\u2028", // line separator
			Expected: `elm.Create("## ", "", "\u2028")`,
		},
		{
			Value:                       "",
			EOL:                             "\u2029", // paragraph separator
			Expected: `elm.Create("## ", "", "\u2029")`,
		},



		{
			Value:                       "Hello world!",
			EOL:                                         "\u000A", // line feed
			Expected: `elm.Create("## ", "Hello world!", "\n")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u000B", // vertical tab
			Expected: `elm.Create("## ", "Hello world!", "\v")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u000C", // form feed
			Expected: `elm.Create("## ", "Hello world!", "\f")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u000D", // carriage return
			Expected: `elm.Create("## ", "Hello world!", "\r")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u000D\u000A", // carriage return, line feed
			Expected: `elm.Create("## ", "Hello world!", "\r\n")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u0085", // next line
			Expected: `elm.Create("## ", "Hello world!", "\u0085")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u2028", // line separator
			Expected: `elm.Create("## ", "Hello world!", "\u2028")`,
		},
		{
			Value:                       "Hello world!",
			EOL:                                         "\u2029", // paragraph separator
			Expected: `elm.Create("## ", "Hello world!", "\u2029")`,
		},



		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u000A", // line feed
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\n")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u000B", // vertical tab
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\v")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u000C", // form feed
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\f")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u000D", // carriage return
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\r")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u000D\u000A", // carriage return, line feed
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\r\n")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u0085", // next line
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\u0085")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u2028", // line separator
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\u2028")`,
		},
		{
			Value:                       "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                                         "\u2029", // paragraph separator
			Expected: `elm.Create("## ", "apple 🍎 BANANA 🍌 Cherry 🍒", "\u2029")`,
		},



		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u000A", // line feed
			Expected: `elm.Create("## ", " extra   spaces! ", "\n")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u000B", // vertical tab
			Expected: `elm.Create("## ", " extra   spaces! ", "\v")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u000C", // form feed
			Expected: `elm.Create("## ", " extra   spaces! ", "\f")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u000D", // carriage return
			Expected: `elm.Create("## ", " extra   spaces! ", "\r")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u000D\u000A", // carriage return, line feed
			Expected: `elm.Create("## ", " extra   spaces! ", "\r\n")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u0085", // next line
			Expected: `elm.Create("## ", " extra   spaces! ", "\u0085")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u2028", // line separator
			Expected: `elm.Create("## ", " extra   spaces! ", "\u2028")`,
		},
		{
			Value:                       " extra   spaces! ",
			EOL:                                              "\u2029", // paragraph separator
			Expected: `elm.Create("## ", " extra   spaces! ", "\u2029")`,
		},



		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u000A", // line feed
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\n")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u000B", // vertical tab
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\v")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u000C", // form feed
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\f")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u000D", // carriage return
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\r")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u000D\u000A", // carriage return, line feed
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\r\n")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u0085", // next line
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\u0085")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u2028", // line separator
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\u2028")`,
		},
		{
			Value:                       "۰۱۲۳۴۵۶۷۸۹",
			EOL:                                     "\u2029", // paragraph separator
			Expected: `elm.Create("## ", "۰۱۲۳۴۵۶۷۸۹", "\u2029")`,
		},
	}

	for testNumber, test := range tests {

		var element elm.HeadingLevel2
		element.Set(test.Value, test.EOL)

		actual   := element.GoString()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value from the .String() method was not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
