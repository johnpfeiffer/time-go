package paginator

import (
	"strconv"
	"strings"
)

// TableOfContentsAnchor is a string that holds the default HTML tag for the table of contents link
const TableOfContentsAnchor = `<h3 id="TOC">`

// Book has many chapters
type Book struct {
	Chapters []string `json:"chapters"`
}

// generateChapters
func generateChapters(text string) (Book, error) {
	b := Book{Chapters: []string{}}
	// TODO: proper parser to break into chunks
	chapterSplit := "CHAPTER "
	parts := strings.Split(text, chapterSplit)
	for i, p := range parts {
		var current string
		if i == 0 {
			current = TableOfContentsAnchor
			// TODO: full TOC links outside?
		} else {
			current = `<h3 id="` + chapterSplit + strconv.Itoa(i) + `">` + chapterSplit
		}
		b.Chapters = append(b.Chapters, current+p)
	}
	return b, nil
}
