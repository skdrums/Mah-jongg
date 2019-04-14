package main

import (
	"fmt"

	"./Tile"
)

func main() {
	//上がり牌
	endTiles := Tile.NewEndTiles()
	for _, tile := range endTiles {
		fmt.Println(tile.JapaneseName)
	}

}
