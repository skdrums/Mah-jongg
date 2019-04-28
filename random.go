package main

import (
	"fmt"

	"Mah-jongg/Score"
	"Mah-jongg/Tile"
)

func main() {
	//上がり牌
	endTiles := Tile.NewEndTiles()
	for i, tile := range endTiles {
		fmt.Printf("%.2d  %s\n",i,tile.JapaneseName)
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
