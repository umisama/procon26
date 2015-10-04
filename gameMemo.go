package main

import (
	"math/rand"
	"runtime"
	"time"
)

type gameMemo struct {
	field     *Field
	stoneBase []*StoneBase
	numStone  int

	limit time.Duration
	Nmax  int
}

func NewGameMemo(path string, limit time.Duration, Nmax int) (Game, error) {
	lines, err := getLinesFromFile(path)
	if err != nil {
		return nil, err
	}
	field, stoneBase, err := readGameMaterials(lines)
	if err != nil {
		return nil, err
	}

	return &gameMemo{
		field:     field,
		stoneBase: stoneBase,
		numStone:  len(stoneBase),
		limit:     limit,
		Nmax:      Nmax,
	}, nil
}

func (game *gameMemo) Run() *Plan {
	return game.Algorithm()
}

func (game *gameMemo) Algorithm() *Plan {
	var best BestMgr

	type Job struct {
		X int
		Y int
	}
	queue := make(chan Job)
	bestCandidates := make(chan *Plan)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(queue chan Job, bestCandidates chan *Plan) {
			for {
				job := <-queue
				bestCandidates <- game.algorithmCheckingPartialScore(job.X, job.Y, best.Score())
			}
		}(queue, bestCandidates)
	}
	go func(candidates chan *Plan) {
		for {
			candidate := <-candidates
			best.Set(candidate)
		}
	}(bestCandidates)

	// push jobs
	end := make(chan struct{})
	go func(queue chan Job) {
		for N := 0; N < game.Nmax; N++ {
			for x := 0; x < 32; x++ {
				for y := 0; y < 32; y++ {
					queue <- Job{
						X: x,
						Y: y,
					}
				}
			}
		}
		end <- struct{}{}
	}(queue)

	select {
	case <-time.Tick(game.limit):
		println("time is up!")
	case <-end:
		println("well done")
	}
	return best.Get()
}

func (game *gameMemo) algorithmCheckingPartialScore(x, y, score int) *Plan {
	var best *Plan
	for _, fStone := range game.stoneBase[0].GetVariations() {
		p := NewPlan(game.field, game.numStone)
		if !p.Put(x, y, fStone) {
			continue
		}
		p = game.sub(score)
		if p != nil && (best == nil || best.Score() > p.Score()) {
			best = p
		}
	}
	return best
}

func (game *gameMemo) sub(latestBestScore int) *Plan {
	p := NewPlan(game.field, game.numStone)
	for it := 0; it < game.numStone; it++ {
		type bestStoneCont struct {
			stone *Stone
			x, y  int
		}
		var bestStone []bestStoneCont
		var bestScore = 0x8fffffff
		sBase := game.stoneBase[it]
		for x := 0; x < 32; x++ {
			for y := 0; y < 32; y++ {
				for _, stone := range sBase.GetVariations() {
					if !p.Put(x, y, stone) {
						continue
					}
					partialScore := p.PartialScoreByExistStones()
					if bestScore > partialScore {
						bestScore = partialScore
						bestStone = []bestStoneCont{{stone, x, y}}
					} else if bestScore == partialScore {
						bestStone = append(bestStone, bestStoneCont{stone, x, y})
					}
					p.Pop()
				}
			}
		}

		if len(bestStone) != 0 {
			bestStoneI := bestStone[rand.Intn(len(bestStone))]
			p.Put(bestStoneI.x, bestStoneI.y, bestStoneI.stone)
		}
		willBestScore := p.Score()
		for i := it; i < game.numStone; i++ {
			willBestScore -= game.stoneBase[i].count
		}
		if willBestScore > latestBestScore {
			break
		}
	}
	return p
}
