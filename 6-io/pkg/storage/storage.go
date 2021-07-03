package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"teachbase/6-io/pkg/crawler"
	"teachbase/6-io/pkg/crawler/spider"
)

// Update обновим поисковый индекс из интернет
func Update(sites []string, store *[]crawler.Document, storeIdx *[]int) {
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
			*store = append(*store, item)
			*storeIdx = append(*storeIdx, i)
			i++

		}
	}
	storeJson, errStore := json.Marshal(store)
	if errStore != nil {
		log.Fatal(errStore)
	}
	errStore = saveFile(storeJson, "store.json")
	if errStore != nil {
		log.Fatal(errStore)
	}
	storeIdxJson, errStoreIdx := json.Marshal(storeIdx)
	if errStoreIdx != nil {
		log.Fatal(errStoreIdx)
	}
	errStoreIdx = saveFile(storeIdxJson, "store_idx.json")
	if errStoreIdx != nil {
		log.Fatal(errStoreIdx)
	}
}

// LoadCache загрузить из кеша
func LoadCache(store *[]crawler.Document, storeIdx *[]int) {
	// Загрузим индекс из файлов
	storeSaved, err := loadFile("store.json")
	if err != nil {
		log.Println("Файл кеша store.json не найден")
	}
	storeIdxSaved, err := loadFile("store_idx.json")
	if err != nil {
		log.Println("Файл кеша store_idx.json не найден")
	}
	err = json.Unmarshal(storeSaved, &store)
	if err != nil {
		log.Fatal(err)
	}
	errIdx := json.Unmarshal(storeIdxSaved, &storeIdx)
	if errIdx != nil {
		log.Fatal(errIdx)
	}
}

// ClearCache удаляет кешированные данные с URL
func ClearCache() {
	err := clear("store.json")
	if err != nil {
		log.Println("Файл кеша store.json не найден")
	}
	err = clear("store_idx.json")
	if err != nil {
		log.Println("Файл кеша store_idx.json не найден")
	}

}

// saveFile сохранение результатов в файл
func saveFile(b []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	err = save(f, b)
	if err != nil {
		return err
	}
	return nil
}

// save сохраняет через io.Writer,
func save(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

// Load загрузка из указанного файла
func loadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return load(f)
}

// load чтение через io.Reader
func load(r io.Reader) ([]byte, error) {
	var b []byte
	var buf = make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		b = append(b, buf...)
	}
	return b, nil
}

// clear удаление файла
func clear(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
