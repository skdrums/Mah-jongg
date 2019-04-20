package Tile

import (
	"sort"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"../Csv"
)

type ID int
type Name string
type Kind string
type ImageName string
type JapaneseName string

type Tile struct {
	ID           ID
	Name         Name
	Kind         Kind
	ImageName    ImageName
	JapaneseName JapaneseName
}

// Tileの生成
func NewTile(id ID, name Name, kind Kind, imageName ImageName, japaneseName JapaneseName) *Tile {
	t := new(Tile)
	t.ID = id
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

var AllTiles = CreateAllTiles()


type NewTileOption func(*Tile)

//Tileの生成オプション
func WithId(id ID) NewTileOption {
	return func(tile *Tile) {
		tile.ID = id
	}
}

//Tileの生成オプション
func WithNameAndKind(name Name, kind Kind) NewTileOption {
	return func(tile *Tile) {
		tile.Name, tile.Kind = name, kind
	}
}

//オプション付Tileの生成
func NewTileWithOption(options ...NewTileOption) (*Tile, error) {
	tile := new(Tile)
	for _, option := range options {
		option(tile)
	}

	//WithIdの場合
	if tile.ID != 0 {
		for _, defaultTile := range AllTiles {
			if tile.ID == defaultTile.ID {
				return &defaultTile, nil
			}
		}
	}
	//WithNameAndKindの場合
	if tile.Name != "" {
		for _, defaultTile := range AllTiles {
			if tile.Name == defaultTile.Name && tile.Kind == defaultTile.Kind {
				return &defaultTile, nil
			}
		}
	}
	return tile, errors.New("couldn't make tile.")
}

func NewHaipai() (haipai []Tile) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 13; i++ {
		tile, _ := NewTileWithOption(WithId(ID(rand.Intn(len(AllTiles)) + 1)))
		haipai = append(haipai, *tile)
	}
	sort.Slice(haipai,func(i,j int) bool {return haipai[i].ID < haipai[j].ID})
	return haipai
}


