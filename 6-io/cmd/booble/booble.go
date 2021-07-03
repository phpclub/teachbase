package main

import (
	"flag"
	"fmt"
	"sort"
	"teachbase/6-io/pkg/crawler"
	"teachbase/6-io/pkg/index"
	"teachbase/6-io/pkg/storage"
)

func main() {
	var (
		s        = flag.String("s", "", "Search string")
		rescan   = flag.Bool("rescan", false, "Rescan URL")
		store    []crawler.Document
		storeIdx []int
		sites    = []string{"https://go.dev/", "https://golang.org/"}
	)
	flag.Parse()
	indexDoc := make(index.Reverse)
	if *s == "" {
		return
	}
	if *rescan {
		storage.ClearCache()
		storage.Update(sites, &store, &storeIdx)
	}
	storage.LoadCache(&store, &storeIdx)
	if len(store) == 0 || len(storeIdx) == 0 {
		storage.Update(sites, &store, &storeIdx)
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
