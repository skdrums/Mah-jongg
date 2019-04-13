package Tile

import (
	"math/rand"
	"strconv"
	"errors"
	"../Csv"
)

type ID int
type Name string
type Kind string
type ImageName string
type JapaneseName string

type Tile struct {
	id           ID
	Name         Name
	Kind         Kind
	ImageName    ImageName
	JapaneseName JapaneseName
}

// Tileの生成
func NewTile(id ID, name Name, kind Kind, imageName ImageName, japaneseName JapaneseName) *Tile {
	t := new(Tile)
	t.id = id
	t.Name, t.Kind, t.ImageName, t.JapaneseName = name, kind, imageName, japaneseName

	return t
}

func CreateAllTiles() (tiles []Tile) {
	records, _ := Csv.CsvReader("./data/tile.csv")
	for key, record := range records {
		if key == 0 {
			continue
		}
		id, _ := strconv.Atoi(record[0])
		tiles = append(tiles, *NewTile(ID(id), Name(record[1]), Kind(record[2]), ImageName(record[3]), JapaneseName(record[4])))
	}
	return tiles
}

type NewTileOption func(*Tile)

//Tileの生成オプション
func WithId(id ID) NewTileOption {
	return func(tile *Tile) {
		tile.id = id
	}
}

//Tileの生成オプション
func WithNameAndKind(name Name, kind Kind) NewTileOption {
	return func(tile *Tile) {
		tile.Name, tile.Kind = name, kind
	}
}

//オプション付Tileの生成
func NewTileWithOption(options ...NewTileOption) (*Tile,error) {
	tile := new(Tile)
	for _, option := range options {
		option(tile)
	}
	defaultTiles := CreateAllTiles()

	//WithIdの場合
	if tile.id != 0 {
		for _, defaultTile := range defaultTiles {
			if tile.id == defaultTile.id {
				return &defaultTile,nil
			}
		}
	}
	//WithNameAndKindの場合
	if tile.Name != "" {
		for _, defaultTile := range defaultTiles {
			if tile.Name == defaultTile.Name && tile.Kind == defaultTile.Kind {
				return &defaultTile,nil
			}
		}
	}
	return tile,errors.New("couldn't make tile.")
}

func NewHaipai() (haipai []Tile) {
	tiles := CreateAllTiles()
	for i := 0; i < 13; i++ {
		tile,_ := NewTileWithOption(WithId(ID(rand.Intn(len(tiles))+1)))
		haipai = append(haipai, *tile)
	}
	return haipai
}
