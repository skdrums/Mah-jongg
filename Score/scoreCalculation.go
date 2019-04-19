package Score

import (
	"fmt"
	"../Tile"
)

type Mentsu struct {
	Tiles    []Tile.Tile
	IsCalled bool
}

type Finisher struct {
	Tile    Tile.Tile
	IsTsumo bool
}

// 上がりハイをここに格納し、点数計算できる状態にする
// <TODO　特殊系の考慮>
type FinishData struct {
	EndTiles                 []Tile.Tile
	Mentsu                   []Mentsu
	Finisher, Dora, HeadTile Tile.Tile
	PublicWind, PrivateWind  Tile.Name
}

//鳴きの考慮なし＠要編集
func NewMentsu(tiles []Tile.Tile) *Mentsu {
	mentsu := new(Mentsu)
	mentsu.Tiles = append(mentsu.Tiles, tiles...)
	mentsu.IsCalled = false
	return mentsu
}

//面前判定の考慮なし＠要編集
func NewFinisher(tile Tile.Tile) *Finisher {
	finisher := new(Finisher)
	finisher.Tile = tile
	finisher.IsTsumo = true
	return finisher
}

func backtrack(i int, t []Tile.Tile, d *FinishData) {
	u := Tile.Tile{}
	k := 0
	for j := i + 1; j < 14; j++ {
		switch {
		case j == i+1 && t[j].ID == t[i].ID && d.HeadTile == u:
			d.HeadTile = t[i]
			fmt.Println(d.HeadTile)
			backtrack(j+1, t, d)
		case j == i+2 && t[j].ID == t[i].ID+2:
			d.Mentsu[k] = *NewMentsu(t[i:3])
			k++
			fmt.Println(d.Mentsu[k])
			backtrack(j+1, t, d)
			k--
		case j == i+2 && t[j].ID == t[i].ID:
			d.Mentsu[k] = *NewMentsu(t[i:3])
			fmt.Println(d.Mentsu[k])
			k++
			backtrack(j+1, t, d)
			k--
		}
	}
}

func NewFinishData(endTiles []Tile.Tile) *FinishData {
	finishData:=new(FinishData)
	finishData.EndTiles = endTiles
	backtrack(0, endTiles, finishData)

	return finishData
}

func devideTiles(endTiles []Tile.Tile) {

}

func calculateAScore(tiles []Tile.Tile) (score int) {

	return score
}
