package Score

import (
	"fmt"

	"Mah-jongg/Tile"
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

var k int = 0
var m int = 0 // iに代入される!isUsedTileな数値を格納
var unusedTile = [14]Tile.Tile{}
var u = Tile.Tile{}
var isUsedTile = [14]bool{}

//var patterns = [3]FinishData{*NewFinishData(),*NewFinishData(t),*NewFinishData(t)}

func backtrack(i int, t []Tile.Tile, d *FinishData) {
	fmt.Printf("@i = %d \n", i)
	fmt.Println(d.HeadTile)

	// jの決定。(i以降の中で一番はじめにくる!isUsedTileな数字)
	var n int = 0 // jに代入される!isUsedTileな数値を格納
	for l := i + 1; l <= 14; l++ {
		if l == 14 || !isUsedTile[l] {
			n = l
			break
		}
	}

	unusedTile = [14]Tile.Tile{} //初期化
	// unusedTileの生成
	for key, value := range t {
		if !isUsedTile[key] {
			unusedTile[key] = value
		}
	}
	for _, value := range unusedTile {
		fmt.Printf("%s ", value.JapaneseName)
	}
	fmt.Printf("\n")

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

	//backtrackのメインロジック
	for j := n; j < 14 && i < 14; j++ {

		fmt.Printf("@i = %d,j = %d, n = %d  \n", i, j, n)
		if j == i+1 && t[j].ID == t[i].ID && *d.HeadTile == u { //雀頭
			d.HeadTile = &t[i]
			fmt.Println(bool(*d.HeadTile == u))
			fmt.Printf("ヘッドだよ")
			fmt.Println(d.HeadTile)
			isUsedTile[i], isUsedTile[j] = true, true
			for l := n + 1; l <= 14; l++ {
				if l == 14 || !isUsedTile[l] {
					m = l
					break
				}
			}
			backtrack(m, t, d)
			d.HeadTile = &u
			isUsedTile[i], isUsedTile[j] = false, false
		} else if j >= i+2 && t[i].ID == t[n].ID { // 暗子 or 一盃口
			flag := false
			if t[i].Kind == t[j].Kind && t[i].ID+Tile.ID(1) == t[j].ID {
				n = j
				for key, value := range unusedTile {
					if value.Kind == t[i].Kind && value.ID == t[i].ID+Tile.ID(2) {
						j = key
						flag = true
						break
					}
				}
			} else if t[i].ID == t[j].ID {
				flag = true
			}
			if flag {
				array := [3]Tile.Tile{}
				array[0], array[1], array[2] = t[i], t[n], t[j]
				for key, tile := range array {
					d.Mentsu[k].Tiles[key] = tile
				}
				k++
				isUsedTile[i], isUsedTile[n], isUsedTile[j] = true, true, true
				for l := i; l <= 14; l++ {
					if l == 14 || !isUsedTile[l] {
						m = l
						break
					}
				}
				backtrack(m, t, d)
				isUsedTile[i], isUsedTile[i+1], isUsedTile[j] = false, false, false
				k--
			}
		} else if j >= i+2 && t[i].Kind == t[n].Kind && t[n].ID == t[i].ID+Tile.ID(1) { // 順子
			array := [3]Tile.Tile{}
			if t[j].ID == t[i].ID+Tile.ID(2) {
				array[0], array[1], array[2] = t[i], t[n], t[j]
			} else {
				continue
			}
			for key, tile := range array {
				d.Mentsu[k].Tiles[key] = tile
			}
			k++
			isUsedTile[i], isUsedTile[n], isUsedTile[j] = true, true, true
			for l := n; l <= 14; l++ {
				if l == 14 || !isUsedTile[l] {
					m = l
					break
				}
			}
			backtrack(m, t, d)
			isUsedTile[i], isUsedTile[n], isUsedTile[j] = false, false, false
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
