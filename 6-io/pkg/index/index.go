package index

import (
	"strings"
	"teachbase/6-io/pkg/crawler"
	"unicode"
)

type Reverse map[string][]int

// tokenize разобьем на фразы без пробелов и спецсимволов
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r) || r == '.' || r == ',' || r == '-' || r == '&'
	})
}

// Add добавление токенов в индекс
func (idx Reverse) Add(docs []crawler.Document) {
	for _, doc := range docs {
		for _, token := range tokenize(doc.Title) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// Search поиск по фразе вернем id документов
func (idx Reverse) Search(s string) [][]int {
	var res [][]int
	for _, token := range tokenize(s) {
		if item, add := idx[token]; add {
			res = append(res, item)
		}
	}
	return res
}
