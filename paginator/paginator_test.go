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
		input += `CHAPTER ` + strconv.Itoa(i) + `\nfoobar <this@example> \n\n`
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
			expected: []string{TableOfContentsAnchor}},
		// TODO: the extra 1 is wrong
		{testName: "chapter only", s: "CHAPTER 1",
			expected: []string{TableOfContentsAnchor, `<h3 id="CHAPTER1">CHAPTER 1</h3>1`}},
		// TODO: the extra 1 is wrong
		{testName: "single chapter", s: `CHAPTER 1\nfoobar`,
			expected: []string{TableOfContentsAnchor, `<h3 id="CHAPTER1">CHAPTER 1</h3>1<br />foobar`}},
		{testName: "single chapter needs escaping", s: `CHAPTER 1\n<foo&bar@>`,
			expected: []string{TableOfContentsAnchor, `<h3 id="CHAPTER1">CHAPTER 1</h3>1<br />&lt;foo&amp;bar@&gt;`}},
		// TODO: the extra 1 and 2 are wrong
		{testName: "two chapters", s: `CHAPTER 1\nfoobar\n\nCHAPTER 2`,
			expected: []string{
				TableOfContentsAnchor,
				`<h3 id="CHAPTER1">CHAPTER 1</h3>1<br />foobar<br /><br />`,
				`<h3 id="CHAPTER2">CHAPTER 2</h3>2`}},
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
