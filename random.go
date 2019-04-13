package main

import (
	"fmt"
	"./Csv"
	"./Tile"
)



func main() {
	records, _ := Csv.CsvReader("./data/tile.csv")
	tiles := Tile.CreateTiles(records)
	for _, tile := range tiles {
		fmt.Println(tile)
		fmt.Printf("\n")
	}
}
