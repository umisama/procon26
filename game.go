package main

import (
	"fmt"
	"github.com/umisama/procon26/materials"
	"github.com/umisama/procon26/plan"
	"sync"
)

type Game interface {
	Run() *plan.Plan
}

func readGameMaterials(lines []string) (*materials.Field, []*materials.StoneBase, error) {
	field, err := materials.NewField(lines[0:32])
	if err != nil {
		return nil, nil, err
	}

	numStone := 0
	fmt.Sscanf(lines[33], "%d", &numStone)

	stoneBase := make([]*materials.StoneBase, numStone)
	for i := 0; i < numStone; i++ {
		newStone, err := materials.NewStoneBase(i, lines[34+9*i:34+9*i+8])
		if err != nil {
			return nil, nil, err
		}
		stoneBase[i] = newStone
	}

	return field, stoneBase, nil
}

type BestMgr struct {
	best *plan.Plan
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

func (b *BestMgr) Set(candidate *plan.Plan) {
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

func (b *BestMgr) set(candidate *plan.Plan) {
	b.best = candidate
	b.score = b.best.Score()
	b.numOfPiece = b.best.NumberOfPiece()
}

func (b *BestMgr) Get() *plan.Plan {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.best
}
