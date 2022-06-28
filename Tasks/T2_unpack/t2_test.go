package unpack

import "testing"

var testCases = []struct {
	input  string
	output string
}{
	{
		"a4bc2d5e",
		"aaaabccddddde",
	},
	{
		"abcd",
		"abcd",
	},
	{
		"45",
		"",
	},
	{
		"",
		"",
	},
	{
		`qwe\4\5`,
		`qwe45`,
	},
	{
		`qwe\45`,
		`qwe44444`,
	},
	{
		`qwe\\5`,
		`qwe\\\\\`,
	},
	{
		`\\2qwe\\5`,
		`\\qwe\\\\\`,
	},
}

func TestUnpack(t *testing.T) {
	for _, c := range testCases {
		output, _ := Unpack(c.input)
		if output != c.output {
			t.Error("test failed")
		}
	}
}
