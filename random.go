package main

import (
	"fmt"

	"./Score"
	"./Tile"
)

func main() {
	//上がり牌
	endTiles := Tile.NewEndTiles()
	for _, tile := range endTiles {
		fmt.Println(tile.JapaneseName)
	}
	data := Score.NewFinishData(endTiles)
	for _, mentsu := range data.Mentsu {
		for _, dataTile := range mentsu.Tiles {
			fmt.Printf("%s|",dataTile.JapaneseName)
		}
		i:=0
		for i<2{
			fmt.Println(data.HeadTile)
		}
	}

}
