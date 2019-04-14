package Score

import (
	"../Tile"

)

type Mentsu struct{
	tiles [3]Tile.Tile
	isCalled bool
}

type Finisher struct{
	Tile Tile.Tile
	isTsumo bool
}

type FinishData struct {
	EndTiles []Tile.Tile
	Mentsu [4]Mentsu
	Finisher,Dora,HeadTile Tile.Tile
	PublicWind,PrivateWind Tile.Name
}

//鳴きの考慮なし＠要編集
func NewMentsu(tiles [3]Tile.Tile)(mentsu *Mentsu){
	mentsu.tiles=tiles
	return mentsu
}

//面前判定の考慮なし＠要編集
func NewFinisher(tile Tile.Tile)(finisher *Finisher){
	finisher.Tile = tile
	return finisher
}

func NewFinishData(endTiles []Tile.Tile)(finishData *FinishData){
	finishData.EndTiles = endTiles
	return finishData
}

func devideTiles(endTiles []Tile.Tile){
	
}

func calculateAScore(tiles []Tile.Tile)(score int){



	return score
}