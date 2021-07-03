package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	"teachbase/4-algo/pkg/crawler"
	"teachbase/4-algo/pkg/crawler/spider"
	"teachbase/4-algo/pkg/index"
)

var (
	s        = flag.String("s", "", "Search string")
	store    []crawler.Document
	storeIdx []int
	sites    = []string{"https://go.dev/", "https://golang.org/"}
)

func main() {
	flag.Parse()
	indexDoc := make(index.IndexReverse)
	if *s == "" {
		return
	}
	var i = 1
	var scanner = spider.New()
	for _, site := range sites {
		res, err := scanner.Scan(site, 2)
		if err != nil {
			fmt.Print(err)
			continue
		}
		for _, item := range res {
			item.ID = i
			item.Title = strings.ToLower(item.Title)
			store = append(store, item)
			storeIdx = append(storeIdx, i)
			i++

		}
	}
	indexDoc.Add(store)
	foundDoc := indexDoc.Search(*s)
	fmt.Println("Найдено:")
	for _, docSlice := range foundDoc {
		for _, doc := range docSlice {
			// Не нашел как кастовать store, решил завести отдельный slice
			dIdx := sort.SearchInts(storeIdx, doc)
			fmt.Println(store[dIdx].URL, store[dIdx].Title)
		}
	}
}
