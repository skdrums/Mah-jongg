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
	Finisher, Dora, HeadTile *Tile.Tile
	PublicWind, PrivateWind  Tile.Name
}

//鳴きの考慮なし＠要編集
func NewMentsu(tiles *[3]Tile.Tile) *Mentsu {
	mentsu := &Mentsu{Tiles: tiles, IsCalled: true}
	return mentsu
}

//面前判定の考慮なし＠要編集
func NewFinisher(tile *Tile.Tile) *Finisher {
	finisher := &Finisher{Tile: tile, IsTsumo: true}
	return finisher
}
// kはtest用flag
var k int = 0
var u = Tile.Tile{}
var isUsedTile = [14]int{}

//var patterns = [3]FinishData{*NewFinishData(),*NewFinishData(t),*NewFinishData(t)}

func backtrack(i int, t []Tile.Tile, d *FinishData) {
	fmt.Printf("@i = %d \n", i)
	fmt.Println(bool(d.HeadTile == &u))
	fmt.Println(d.HeadTile)
	//できていれば表示
	if i == 14 {
		for _, mentsu := range d.Mentsu {
			for _, dataTile := range mentsu.Tiles {
				fmt.Printf("%s|", dataTile.JapaneseName)
			}
		}
		i := 0
		for i < 2 {
			fmt.Printf("%s|", d.HeadTile.JapaneseName)
			i++
		}
	}
	//backtrack中身
	for j := i + 1; j < 14 && j-i < 3; j++ {

		fmt.Printf("@i = %d,j = %d  \n", i, j)
		if j == i+1 && t[j].ID == t[i].ID && *d.HeadTile == u { //雀頭
			d.HeadTile = &t[i]
			fmt.Println(bool(*d.HeadTile == u))
			fmt.Printf("ヘッドだよ")
			fmt.Println(d.HeadTile)
			isUsedTile[i],isUsedTile[j]=1,1
			backtrack(j+1, t, d)
			d.HeadTile = &u
			isUsedTile[i],isUsedTile[j]=0,0
		} else if j == i+2 && t[j].Kind == t[i].Kind && t[j].ID == t[i].ID+Tile.ID(2) && t[j].ID == t[i+1].ID+Tile.ID(1) { // 順子
			slice := t[i : i+3]
			array := [3]Tile.Tile{}
			for key, s := range slice {
				array[key] = s
			}
			for key, tile := range array {
				d.Mentsu[k].Tiles[key] = tile
			}
			// fmt.Printf("Mentsu[%d].Tiles = ", k)
			// fmt.Println(d.Mentsu[k].Tiles)
			k++
			isUsedTile[i],isUsedTile[i+1],isUsedTile[j]=1,1,1
			backtrack(j+1, t, d)
			isUsedTile[i],isUsedTile[i+1],isUsedTile[j]=0,0,0
			k--
		} else if j == i+2 && t[j].ID == t[i].ID { // 暗子
			slice := t[i : i+3]
			array := [3]Tile.Tile{}
			for i, s := range slice {
				array[i] = s
			}
			for key, tile := range array {
				d.Mentsu[k].Tiles[key] = tile
			}
			// fmt.Printf("Mentsu[%d].Tiles = ", k)
			// fmt.Println(d.Mentsu[k].Tiles)
			k++
			isUsedTile[i],isUsedTile[i+1],isUsedTile[j]=1,1,1
			backtrack(j+1, t, d)
			isUsedTile[i],isUsedTile[i+1],isUsedTile[j]=0,0,0
			k--
		}
	}
}

func NewFinishData(endTiles []Tile.Tile) *FinishData {
	tiles := [4][3]Tile.Tile{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			tiles[i][j] = *new(Tile.Tile)
		}
	}
	mentsu := [4]Mentsu{*NewMentsu(&tiles[0]), *NewMentsu(&tiles[1]), *NewMentsu(&tiles[2]), *NewMentsu(&tiles[3])}
	finishData := FinishData{EndTiles: endTiles, Mentsu: &mentsu, Finisher: new(Tile.Tile), Dora: new(Tile.Tile), HeadTile: new(Tile.Tile)}

	return &finishData
}

func DevideTiles(endTiles []Tile.Tile) *FinishData {
	finishData := NewFinishData(endTiles)
	backtrack(0, endTiles, finishData)

	return finishData
}

func CalculateAScore(tiles []Tile.Tile) (score int) {

	return score
}
