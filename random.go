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
	data := Score.DevideTiles(endTiles)
	fmt.Printf("headtile")
	fmt.Println(data.HeadTile)
	for _, mentsu := range data.Mentsu {
		for _, dataTile := range mentsu.Tiles {
			fmt.Printf("%s|",dataTile.JapaneseName)
		}		
	}
	i:=0
	for i<2{
		fmt.Printf("%s|",data.HeadTile.JapaneseName)
		i++
	}
}
