package db

import (
	"errors"
	"sync"
)

type PokeTable struct {
	poke map[int32]*Poke
	mu   sync.Mutex
}

type Poke struct {
	Name string
}

var pokeTable *PokeTable

func NewPokeTable() *PokeTable {
	if pokeTable == nil {
		pokeTable = &PokeTable{
			poke: make(map[int32]*Poke),
		}
	}

	return pokeTable
}

func GetPoke(no int32) (Poke, error) {
	if pokeTable == nil {
		return Poke{}, errors.New("none index")
	}

	pokeTable.mu.Lock()
	defer pokeTable.mu.Unlock()

	for i, p := range pokeTable.poke {
		if i == no {
			return *p, nil
		}
	}

	return Poke{}, errors.New("not found")
}

func (pt *PokeTable) AddPoke(no int32, name string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	pt.poke[no] = &Poke{Name: name}
}
