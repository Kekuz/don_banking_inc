package pkg

import (
	"encoding/csv"
	"os"
)

func CreateFile(filename string, data [][]string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func RemoveFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func GetFileData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
