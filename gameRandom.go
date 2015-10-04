package main

import (
	"fmt"
	"math/rand"
)

type gameRandom struct {
	field     *Field
	stoneBase []*StoneBase
	numStone  int
}

func NewGameRandom(path string) (Game, error) {
	lines, err := getLinesFromFile(path)
	if err != nil {
		return nil, err
	}
	field, stoneBase, err := readGameMaterials(lines)
	if err != nil {
		return nil, err
	}

	return &gameRandom{
		field:     field,
		stoneBase: stoneBase,
		numStone:  len(stoneBase),
	}, nil
}

func (game *gameRandom) Run() *Plan {
	return game.Algorithm(300)
}

func (game *gameRandom) Algorithm(times int) *Plan {
	var best *Plan = nil
	for i := 0; i < times; i++ {
		p := NewPlan(game.field, game.numStone)
	base:
		for _, sBase := range game.stoneBase {
			for _, x := range rand.Perm(32) {
				for _, y := range rand.Perm(32) {
					stones := sBase.GetVariations()
					for _, i := range rand.Perm(len(stones)) {
						if p.Put(x, y, stones[i]) {
							continue base
						}
					}
				}
			}
		}
		if best == nil || best.Score() > p.Score() {
			fmt.Printf("new answer! -> score(%d) on %d times", p.Score(), i)
			best = p
		}
	}
	return best
}
