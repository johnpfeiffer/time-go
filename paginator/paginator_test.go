package paginator

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestManual(t *testing.T) {
	input := ""
	for i := 1; i < 12; i++ {
		current := strconv.Itoa(i)
		input += `CHAPTER ` + current + `
foobar <this@example>
and more unique stuff` + current + current + `&haha \n\n`
	}
	b, _ := generateHTMLChapters(input)
	fmt.Println(toHTMLPage(b))
}

func TestChapters(t *testing.T) {
	var testCases = []struct {
		testName string
		s        string
		expected []string
	}{
		{testName: "empty", s: "",
			expected: []string{}},
		{testName: "chapter only", s: "CHAPTER 1",
			expected: []string{`<h3 id="CHAPTER1">CHAPTER 1</h3><br />
`}},
		{testName: "single chapter no newline", s: `CHAPTER 1 foobar`,
			expected: []string{`<h3 id="CHAPTER1">CHAPTER 1</h3> foobar<br />
`}},
		{testName: "single chapter simple", s: `CHAPTER 1
foobar`,
			expected: []string{`<h3 id="CHAPTER1">CHAPTER 1</h3><br />
foobar<br />
`}},
		{testName: "single chapter trailing whitespace and newlines", s: `CHAPTER 1
foobar 	
`,
			expected: []string{`<h3 id="CHAPTER1">CHAPTER 1</h3><br />
foobar<br />
`}},
		{testName: "single chapter needs escaping", s: `CHAPTER 1
<foo&bar@>`,
			expected: []string{`<h3 id="CHAPTER1">CHAPTER 1</h3><br />
&lt;foo&amp;bar@&gt;<br />
`}},
		{testName: "two chapters", s: `CHAPTER 1
foo

CHAPTER 2
bar`,
			expected: []string{
				`<h3 id="CHAPTER1">CHAPTER 1</h3><br />
foo<br />
`,
				`<h3 id="CHAPTER2">CHAPTER 2</h3><br />
bar<br />`}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result, err := generateHTMLChapters(tc.s)
			assertErrIsNil(t, err)
			assertStringSlicesEqual(t, tc.expected, result.Chapters)
		})
	}
}

// Helper Functions
func assertErrIsNil(t *testing.T, err error) {
	if err != nil {
		t.Error("\nExpected nil error \nReceived: ", err)
	}
}

func assertStringSlicesEqual(t *testing.T, expected []string, result []string) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("\nExpected:", expected, "\nReceived:", result)
	}
}
