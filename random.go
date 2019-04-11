package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

type Tile struct {
	id   int
	name string
	kind string
	imageName string
	japaneseName string
}

func newTile(id int, name, kind, imageName, japaneseName string) *Tile {
	t := new(Tile)
	t.id = id
	t.name = name
	t.kind = kind
	t.imageName = imageName
	t.japaneseName = japaneseName
	return t
}

func main() {
	// ファイルを開く
	file, err := os.Open("./data/tile.csv")
	failOnError(err)

	// プログラムが終わったらファイルを閉じる
	defer file.Close()

	// // 12byte格納可能なスライスを用意する
	// message := make([]byte, 12)

	// // ファイル内のデータをスライスに読み出す
	// _, err = file.Read(message)
	// if err != nil { // エラー処理
	//     log.Fatal(err)
	// }

	// // 文字列にして表示
	// fmt.Print(string(message))

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	failOnError(err)

	tiles := []Tile{}
	for key, record := range records {
		if key == 0 {
			continue
		}
		id, _ := strconv.Atoi(record[0])
		tiles = append(tiles, *newTile(id, record[1], record[2], record[3], record[4]))
	}
	for _, tile := range tiles {
		fmt.Println(tile)
		fmt.Printf("\n")
	}
	fmt.Println(t.japaneseName)
}
