package Score

import (
	"fmt"

	"../Tile"
)

type Mentsu struct {
	Tiles    *[3]Tile.Tile
	IsCalled bool
}

type Finisher struct {
	Tile    *Tile.Tile
	IsTsumo bool
}

// 上がりハイをここに格納し、点数計算できる状態にする
// <TODO　特殊系の考慮>
type FinishData struct {
	EndTiles                 []Tile.Tile
	Mentsu                   *[4]Mentsu
	Finisher, Dora, HeadTile Tile.Tile
	PublicWind, PrivateWind  Tile.Name
}

//鳴きの考慮なし＠要編集
func NewMentsu(tiles *[3]Tile.Tile) *Mentsu {
	mentsu := &Mentsu{Tiles : tiles, IsCalled: true} 
	return mentsu
}

//面前判定の考慮なし＠要編集
func NewFinisher(tile *Tile.Tile) *Finisher {
	finisher := &Finisher{Tile:tile,IsTsumo:true}
	return finisher
}

func backtrack(i int, t []Tile.Tile, d *FinishData) {
	u := Tile.Tile{}
	k := 0
	for j := i + 1; j < 14 && j-i < 3; j++ {
		fmt.Printf("i = %d,j = %d  ", i, j)
		if j == i+1 && t[j].ID == t[i].ID && d.HeadTile == u {
			d.HeadTile = t[i]
			fmt.Println(d.HeadTile)
			backtrack(j+1, t, d)
		} else if j == i+2 && t[j].Kind == t[i].Kind {
			slice := t[i:3]
			array := [3]Tile.Tile{}
			for i,s :=range slice{
				array[i]=s
			}
			fmt.Printf("array = ")
			fmt.Println(array)
			shuntsu := NewMentsu(&array)
			fmt.Printf("shuntsu = ")
			fmt.Println(shuntsu.Tiles)
			d.Mentsu[k] = *shuntsu
			k++
			fmt.Println(d.Mentsu[k].IsCalled)
			for _, tile := range d.Mentsu[k].Tiles {
				fmt.Println(tile)
			}
			backtrack(j+1, t, d)
			k--
		} else if j == i+2 && t[j].ID == t[i].ID {
			slice := t[i:3]
			array := [3]Tile.Tile{}
			for i,s :=range slice{
				array[i]=s
			}
			d.Mentsu[k] = *NewMentsu(&array)
			fmt.Println(d.Mentsu[k].IsCalled)
			for _, tile := range d.Mentsu[k].Tiles {
				fmt.Println(tile)
			}
			k++
			backtrack(j+1, t, d)
			k--
		}
	}
}

func NewFinishData(endTiles []Tile.Tile) *FinishData {
	finishData := new(FinishData)
	finishData.EndTiles = endTiles

	return finishData
}

func DevideTiles(endTiles []Tile.Tile) *FinishData{
	finishData:=NewFinishData(endTiles)
	backtrack(0, endTiles, finishData)

	return finishData
}

func CalculateAScore(tiles []Tile.Tile) (score int) {

	return score
}
