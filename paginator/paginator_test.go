package paginator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateChaptersEmpty(t *testing.T) {
	t.Parallel()
	result, err := generateChapters("")
	assertErrIsNil(t, err)
	assertStringSlicesEqual(t, []string{TableOfContentsAnchor}, result.Chapters)
}

func TestGenerateChapters(t *testing.T) {
	t.Parallel()
	result, err := generateChapters("CHAPTER 1")
	assertErrIsNil(t, err)
	assertStringSlicesEqual(t,
		[]string{TableOfContentsAnchor, `<h3 id="CHAPTER 1">CHAPTER 1`}, result.Chapters)
}

func TestChapters(t *testing.T) {
	var testCases = []struct {
		s        string
		expected []string
	}{
		{s: "", expected: []string{TableOfContentsAnchor}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s to chapters with anchors", tc.s), func(t *testing.T) {
			result, err := generateChapters(tc.s)
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
