package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"
	"teachbase/6-io/pkg/crawler"
	"teachbase/6-io/pkg/crawler/spider"
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
		err := storage.Clear("store.json")
		if err != nil {
			log.Fatal(err)
		}
		err = storage.Clear("store_idx.json")
		if err != nil {
			log.Fatal(err)
		}
	}
	// Загрузим индекс из файлов
	storeSaved, errStore := storage.Load("store.json")
	storeIdxSaved, errStoreIdx := storage.Load("store_idx.json")
	// Если вернуло пустоту или произошла ошибка запускам scanner
	if len(storeSaved) == 0 || errStore != nil || len(storeIdxSaved) == 0 || errStoreIdx != nil {
		fmt.Println("Cохраненный индекс не найден")
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
		storeJson, errStore := json.Marshal(store)
		if errStore != nil {
			log.Fatal(errStore)
		}
		errStore = storage.Save(string(storeJson), "store.json")
		if errStore != nil {
			log.Fatal(errStore)
		}
		storeIdxJson, errStoreIdx := json.Marshal(storeIdx)
		if errStoreIdx != nil {
			log.Fatal(errStoreIdx)
		}
		errStoreIdx = storage.Save(string(storeIdxJson), "store_idx.json")
		if errStoreIdx != nil {
			log.Fatal(errStoreIdx)
		}
	} else {
		fmt.Println("Используется сохраненный индекс")
		err := json.Unmarshal(storeSaved, &store)
		if err != nil {
			log.Fatal(err)
		}
		errIdx := json.Unmarshal(storeIdxSaved, &storeIdx)
		if errIdx != nil {
			log.Fatal(errIdx)
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
