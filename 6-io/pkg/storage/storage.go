package storage

import (
	"errors"
	"io"
	"os"
)

// Save сохранение результатов
func Save(s string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		return err
	}
	return nil
}

// Load загрузка из указанного файла
func Load(filename string) ([]byte, error) {
	var (
		ErrUnsupportedFile = errors.New("unsupported file")
	)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fStat, err := f.Stat()
	if err != nil {
		return nil, ErrUnsupportedFile
	}
	if fStat.Size() == 0 {
		return nil, ErrUnsupportedFile
	}

	data := make([]byte, fStat.Size())
	for {
		_, err := f.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
	}
	return data, nil
}

// Clear удаление кеша
func Clear(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
