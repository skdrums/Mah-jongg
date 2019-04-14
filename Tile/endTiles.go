package Tile

import (
	"math/rand"
	"sort"
	"time"
)

var ShuntsuOrAnko = [2]string{"Shuntsu", "Anko"}
var NumberOrChanese = [2]string{"Number", "Chanese"}
var NumberKind = [5]Kind{"circles", "bamboos", "characters"}
var NumberName = [16]Name{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var ChaneseKind = [2]Kind{"winds", "dragons"}
var WindName = [4]Name{"east", "south", "west", "north"}
var DragonName = [3]Name{"white", "green", "red"}

type TypeOfMentsu struct {
	Mentsu          string
	NumberOrChanese string
	Name            Name
	Kind            Kind
}

func CreateMentsu() (mentsu []Tile) {
	typeOfMentsu := new(TypeOfMentsu)
	rand.Seed(time.Now().UnixNano())
	typeOfMentsu.Mentsu = ShuntsuOrAnko[rand.Intn(2)]
	//順子の場合
	if typeOfMentsu.Mentsu == "Shuntsu" {
		typeOfMentsu.Kind = NumberKind[rand.Intn(3)]
		typeOfMentsu.Name = NumberName[rand.Intn(7)]
		tile, _ := NewTileWithOption(WithNameAndKind(typeOfMentsu.Name, typeOfMentsu.Kind))
		mentsu = append(mentsu, *tile)
		for i := 1; i < 3; i++ {
			tile, _ := NewTileWithOption(WithId(tile.id + ID(i)))
			mentsu = append(mentsu, *tile)
		}
	} else {
		//暗子の場合
		typeOfMentsu.NumberOrChanese = NumberOrChanese[rand.Intn(2)]
		if typeOfMentsu.NumberOrChanese == "Number" {
			typeOfMentsu.Kind = NumberKind[rand.Intn(3)]
			typeOfMentsu.Name = NumberName[rand.Intn(9)]
		} else {
			typeOfMentsu.Kind = ChaneseKind[rand.Intn(2)]
			if typeOfMentsu.Kind == "winds" {
				typeOfMentsu.Name = WindName[rand.Intn(4)]
			} else {
				typeOfMentsu.Name = DragonName[rand.Intn(3)]
			}
		}
		for i := 0; i < 3; i++ {
			tile, _ := NewTileWithOption(WithNameAndKind(typeOfMentsu.Name, typeOfMentsu.Kind))
			mentsu = append(mentsu, *tile)
		}
	}
	return mentsu
}

func tilesIsCorrect(tiles []Tile) (isCorrect bool) {
	isCorrect = false
	i := 0
	for i < len(tiles)-4 {
		if tiles[i].id == tiles[i+4].id {
			return isCorrect
		} else if tiles[i].id == tiles[i+3].id {
			i += 4
			continue
		} else if tiles[i].id == tiles[i+2].id {
			i += 3
			continue
		} else if tiles[i].id == tiles[i+1].id {
			i += 2
			continue
		}
		i++
	}
	isCorrect = true
	return isCorrect
}

func NewEndTiles()(endTiles []Tile){
	isCorrect := false
	for isCorrect == false {
		endTiles = []Tile{}
		//メンツの追加
		for i := 0; i < 4; i++ {
			endTiles = append(endTiles, CreateMentsu()...)
		}
		//雀頭の生成
		headTile, _ := NewTileWithOption(WithId(ID(rand.Intn(len(AllTiles)))))
		for i := 0; i < 2; i++ {
			endTiles = append(endTiles, *headTile)
		}
		sort.Slice(endTiles, func(i, j int) bool { return endTiles[i].id < endTiles[j].id })
		isCorrect = tilesIsCorrect(endTiles)
	}
	return endTiles
}
