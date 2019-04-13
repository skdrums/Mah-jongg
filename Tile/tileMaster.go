package Tile

import (
	"strconv"
)

type Tile struct {
	id           int
	name         string
	kind         string
	imageName    string
	japaneseName string
}

func NewTile(id int, name, kind, imageName, japaneseName string) *Tile {
	t := new(Tile)
	t.id = id
	t.name = name
	t.kind = kind
	t.imageName = imageName
	t.japaneseName = japaneseName
	return t
}

func CreateTiles(records [][]string) (tiles []Tile) {
	for key, record := range records {
		if key == 0 {
			continue
		}
		id, _ := strconv.Atoi(record[0])
		tiles = append(tiles, *NewTile(id, record[1], record[2], record[3], record[4]))
	}
	return tiles
}