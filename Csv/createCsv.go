package Csv

import (
	"log"
	"encoding/csv"
	"os"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func CsvReader(csvFile string) (records [][]string, err error) {
	// ファイルを開く
	file, err := os.Open(csvFile)
	failOnError(err)
	// プログラムが終わったらファイルを閉じる
	defer file.Close()
	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	failOnError(err)

	return records, err
}