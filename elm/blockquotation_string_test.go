package elm_test

import (
	"github.com/reiver/go-gmi/elm"

	"testing"
)

func TestBlockQuotation_String(t *testing.T) {

	tests := []struct{
		Value    string
		EOL      string
		Expected string
	}{
		{
			Value:     "",
			EOL:       "\u000A", // line feed
			Expected: ">\u000A",
		},
		{
			Value:     "",
			EOL:       "\u000B", // vertical tab
			Expected: ">\u000B",
		},
		{
			Value:     "",
			EOL:       "\u000C", // form feed
			Expected: ">\u000C",
		},
		{
			Value:     "",
			EOL:       "\u000D", // carriage return
			Expected: ">\u000D",
		},
		{
			Value:     "",
			EOL:       "\u000D\u000A", // carriage return, line feed
			Expected: ">\u000D\u000A",
		},
		{
			Value:     "",
			EOL:       "\u0085", // next line
			Expected: ">\u0085",
		},
		{
			Value:     "",
			EOL:       "\u2028", // line separator
			Expected: ">\u2028",
		},
		{
			Value:     "",
			EOL:       "\u2029", // paragraph separator
			Expected: ">\u2029",
		},



		{
			Value:     "Hello world!",
			EOL:                   "\u000A", // line feed
			Expected: ">Hello world!\u000A",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u000B", // vertical tab
			Expected: ">Hello world!\u000B",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u000C", // form feed
			Expected: ">Hello world!\u000C",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u000D", // carriage return
			Expected: ">Hello world!\u000D",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u000D\u000A", // carriage return, line feed
			Expected: ">Hello world!\u000D\u000A",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u0085", // next line
			Expected: ">Hello world!\u0085",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u2028", // line separator
			Expected: ">Hello world!\u2028",
		},
		{
			Value:     "Hello world!",
			EOL:                   "\u2029", // paragraph separator
			Expected: ">Hello world!\u2029",
		},



		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u000A", // line feed
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u000A",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u000B", // vertical tab
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u000B",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u000C", // form feed
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u000C",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u000D", // carriage return
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u000D",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u000D\u000A", // carriage return, line feed
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u000D\u000A",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u0085", // next line
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u0085",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u2028", // line separator
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u2028",
		},
		{
			Value:     "apple 🍎 BANANA 🍌 Cherry 🍒",
			EOL:                                   "\u2029", // paragraph separator
			Expected: ">apple 🍎 BANANA 🍌 Cherry 🍒\u2029",
		},



		{
			Value:     " extra   spaces! ",
			EOL:                           "\u000A", // line feed
			Expected: "> extra   spaces! "+"\u000A",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u000B", // vertical tab
			Expected: "> extra   spaces! "+"\u000B",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u000C", // form feed
			Expected: "> extra   spaces! "+"\u000C",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u000D", // carriage return
			Expected: "> extra   spaces! "+"\u000D",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u000D\u000A", // carriage return, line feed
			Expected: "> extra   spaces! "+"\u000D\u000A",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u0085", // next line
			Expected: "> extra   spaces! "+"\u0085",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u2028", // line separator
			Expected: "> extra   spaces! "+"\u2028",
		},
		{
			Value:     " extra   spaces! ",
			EOL:                           "\u2029", // paragraph separator
			Expected: "> extra   spaces! "+"\u2029",
		},



		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u000A", // line feed
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u000A",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u000B", // vertical tab
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u000B",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u000C", // form feed
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u000C",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u000D", // carriage return
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u000D",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u000D\u000A", // carriage return, line feed
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u000D\u000A",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u0085", // next line
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u0085",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u2028", // line separator
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u2028",
		},
		{
			Value:     "۰۱۲۳۴۵۶۷۸۹",
			EOL:               "\u2029", // paragraph separator
			Expected: ">۰۱۲۳۴۵۶۷۸۹\u2029",
		},
	}

	for testNumber, test := range tests {

		var element elm.BlockQuotation
		element.Set(test.Value, test.EOL)

		actual   := element.String()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value from the .String() method was not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
