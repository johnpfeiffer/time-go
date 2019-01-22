package paginator

import (
	"html"
	"strings"
	"unicode"
)

// TableOfContentsAnchor is a string that holds the default HTML tag for the table of contents link
const TableOfContentsAnchor = `<h3 id="TOC">TOC</h3>`

// Book has many chapters
type Book struct {
	Chapters []string `json:"chapters"`
}

// Chapter encapsulates the identifier and the text
type Chapter struct {
	Title string
	Text  string
}

// generateHTMLChapters
func generateHTMLChapters(text string) (Book, error) {
	b := Book{Chapters: []string{}}

	// TODO: TableOfContentsAnchor
	for _, c := range findChapters(text) {
		chapterHeader := `<h3 id="` + strings.Replace(c.Title, " ", "", -1) + `">` + c.Title + `</h3>`

		// TODO: modularize with a function that converts the chapter text
		converted := html.EscapeString(c.Text)
		// platform agnostic newline splitter
		textLines := strings.Split(strings.Replace(converted, "\r\n", "\n", -1), "\n")
		convertedAndBreaks := ""
		for _, t := range textLines {
			convertedAndBreaks += t + "<br />\n"
		}

		b.Chapters = append(b.Chapters, chapterHeader+convertedAndBreaks)
	}
	return b, nil
}

// findChapters relies on the heuristic of CHAPTER in all caps and returns the chapter titles and contents
func findChapters(text string) []Chapter {
	chapters := []Chapter{}
	chapterSplit := "CHAPTER "
	parts := strings.Split(text, chapterSplit)
	for _, p := range parts {
		if len(p) < 1 {
			continue
		}
		chapterIdentifier := discoverChapterIdentifier(p)
		textWithoutChapterIdentifier := p[len(chapterIdentifier):]
		c := Chapter{
			Title: chapterSplit + chapterIdentifier,
			// remove trailing space (and newlines)
			Text: strings.TrimRightFunc(textWithoutChapterIdentifier, unicode.IsSpace),
		}
		chapters = append(chapters, c)
	}
	return chapters
}

// discoverChapterIdentifier finds the thing after the chapter title, usually a number like 1 or 2
func discoverChapterIdentifier(s string) string {
	leftTrimmed := strings.TrimLeftFunc(s, unicode.IsSpace)
	parts := strings.Fields(leftTrimmed)
	return parts[0]
}

func toHTMLPage(b Book) string {
	s := `<html>` + strings.Join(b.Chapters, "<hr>\n") + `</html>`
	return s
}
