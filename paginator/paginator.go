package paginator

import (
	"strconv"
	"strings"
)

// TableOfContentsAnchor is a string that holds the default HTML tag for the table of contents link
const TableOfContentsAnchor = `<h3 id="TOC">TOC</h3>`

// Book has many chapters
type Book struct {
	Chapters []string `json:"chapters"`
}

// generateHTMLChapters
func generateHTMLChapters(text string) (Book, error) {
	b := Book{Chapters: []string{}}
	// TODO: proper parser to break into chunks
	chapterSplit := "CHAPTER "
	parts := strings.Split(text, chapterSplit)
	for i, original := range parts {
		var chapterHeader string
		if i == 0 {
			chapterHeader = TableOfContentsAnchor
			// TODO: full TOC links outside?
		} else {
			chapterNumber := strconv.Itoa(i)
			// TODO: use original chapter numbers
			chapterHeader = `<h3 id="` + strings.TrimSpace(chapterSplit) + chapterNumber + `">` + chapterSplit + chapterNumber + `</h3>`
		}
		converted := strings.Replace(original, "\\n", "<br />", -1)
		b.Chapters = append(b.Chapters, chapterHeader+converted)
	}
	return b, nil
}
