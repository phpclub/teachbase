package main

import (
	"flag"
	"fmt"
	"github.com/phpclub/teachbase/2-flag/pkg/crawler"
	"github.com/phpclub/teachbase/2-flag/pkg/crawler/spider"
	"strings"
)

var (
	s     string
	store []crawler.Document
	sites = []string{"https://go.dev/", "https://golang.org/"}
)

func init() {
	flag.StringVar(&s, "s", "", "Search string")
}

func main() {
	flag.Parse()
	if s == "" {
		return
	}
	var scanner = spider.New()
	for _, site := range sites {
		res, err := scanner.Scan(site, 2)
		if err != nil {
			fmt.Print(err)
			continue
		}
		store = append(store, res...)
	}
	// искать по s если она указана
	fmt.Println("Найдено:")
	for _, link := range find(strings.ToLower(s)) {
		fmt.Println(link.URL, link.Title)
	}
}

// _find возвращает слайс найденных документов
func find(value string) []crawler.Document {
	var foundDocs []crawler.Document
	for _, v := range store {
		if strings.Contains(strings.ToLower(v.Title), value) {
			foundDocs = append(foundDocs, v)
		}
	}
	return foundDocs
}
