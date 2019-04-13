package main

import (
	"fmt"

	"./Tile"
)

func main() {
	// tiles := Tile.CreateAllTiles()

	// //for test.
	// for _, tile := range tiles {
	// 	fmt.Println(tile)
	// 	fmt.Printf("\n")
	// }

	haipai := Tile.NewHaipai()
	for _, tile := range haipai {
		fmt.Println(tile.JapaneseName)
	}
	fmt.Printf("\n")

	endTiles := Tile.NewEndTiles()
	for _, tile := range endTiles {
		fmt.Println(tile.JapaneseName)
	}

}
