package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type Game interface {
	Run() *Plan
}

func readGameMaterials(lines []string) (*Field, []*StoneBase, error) {
	field, err := NewField(lines[0:32])
	if err != nil {
		return nil, nil, err
	}

	numStone := 0
	fmt.Sscanf(lines[33], "%d", &numStone)

	stoneBase := make([]*StoneBase, numStone)
	for i := 0; i < numStone; i++ {
		newStone, err := NewStoneBase(i, lines[34+9*i:34+9*i+8])
		if err != nil {
			return nil, nil, err
		}
		stoneBase[i] = newStone
	}

	return field, stoneBase, nil
}

func getLinesFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\n")
	return lines, err
}

type BestMgr struct {
	best *Plan
	mu   sync.Mutex

	// cache
	score      int
	numOfPiece int
}

func (b *BestMgr) Score() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.best == nil {
		return 0x8fffffff
	}
	return b.score
}

func (b *BestMgr) Set(candidate *Plan) {
	if candidate == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	if b.best == nil || b.score > candidate.Score() {
		b.set(candidate)
	} else if b.score == candidate.Score() && b.numOfPiece > candidate.NumberOfPiece() {
		b.set(candidate)
	}
}

func (b *BestMgr) set(candidate *Plan) {
	b.best = candidate
	b.score = b.best.Score()
	b.numOfPiece = b.best.NumberOfPiece()
	println(b.score, b.numOfPiece)
}

func (b *BestMgr) Get() *Plan {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.best
}
