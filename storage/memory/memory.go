package memory

import "github.com/robertpeteuil/minesweeper-api/types"

type DB struct {
	games map[string]*types.Game
}

func New() *DB {
	return &DB{
		games: make(map[string]*types.Game),
	}
}
