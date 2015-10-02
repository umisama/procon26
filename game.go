package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("pprof.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	game, err := NewGameFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	plan := game.Run()

	fmt.Println(plan)
	fmt.Println(plan.Score())
}

type Game struct {
	field     *Field
	stoneBase []*StoneBase
	numStone  int
}

func NewGameFromString(input string) (*Game, error) {
	lines := strings.Split(input, "\n")
	return newGame(lines)
}

func NewGameFromFile(path string) (*Game, error) {
	lines, err := getLinesFromFile(path)
	if err != nil {
		return nil, err
	}
	return newGame(lines)
}

func newGame(lines []string) (*Game, error) {
	field, err := NewField(lines[0:32])
	if err != nil {
		return nil, err
	}

	numStone := 0
	fmt.Sscanf(lines[33], "%d", &numStone)

	stoneBase := make([]*StoneBase, numStone)
	for i := 0; i < numStone; i++ {
		newStone, err := NewStoneBase(i, lines[34+9*i:34+9*i+8])
		if err != nil {
			return nil, err
		}
		stoneBase[i] = newStone
	}

	return &Game{
		field:     field,
		stoneBase: stoneBase,
		numStone:  numStone,
	}, nil
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

func (game *Game) Run() *Plan {
	return game.AlgorithmCheckingPartialScore()
}

func (game *Game) AlgorithmRandom(times int) *Plan {
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
			fmt.Println("new answer! -> score(%d) on %d times", p.Score(), i)
			best = p
		}
	}
	return best
}

func (game *Game) AlgorithmCheckingPartialScore() *Plan {
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
		for N := 0; N < 1; N++ {
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
	case <-time.Tick(100 * time.Minute):
		println("time is up!")
	case <-end:
		println("well done")
	}
	return best.Get()
}

func (game *Game) algorithmCheckingPartialScore(x, y, score int) *Plan {
	var best *Plan
	for _, fStone := range game.stoneBase[0].GetVariations() {
		p := NewPlan(game.field, game.numStone)
		if !p.Put(x, y, fStone) {
			continue
		}
		p = game.sub()
		if p != nil && (best == nil || best.Score() > p.Score()) {
			best = p
		}
	}
	return best
}

func (game *Game) sub() *Plan {
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
					if !p.TestPut(x, y, stone) {
						continue
					}
					partialScore := p.PartialScoreByExistStones()
					if bestScore > partialScore {
						bestScore = partialScore
						bestStone = []bestStoneCont{{stone, x, y}}
					} else if bestScore == partialScore {
						bestStone = append(bestStone, bestStoneCont{stone, x, y})
					}
					p.ClearTestStone()
				}
			}
		}

		if len(bestStone) != 0 {
			bestStoneI := bestStone[rand.Intn(len(bestStone))]
			p.Put(bestStoneI.x, bestStoneI.y, bestStoneI.stone)
		}
	}
	return p
}

type BestMgr struct {
	best  *Plan
	score int
	mu    sync.Mutex
}

func (b *BestMgr) Score() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.score
}

func (b *BestMgr) Set(candidate *Plan) {
	if candidate == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	if b.best == nil || b.best.Score() > candidate.Score() {
		b.best = candidate
		b.score = b.best.Score()
		println(b.score)
	}
}

func (b *BestMgr) Get() *Plan {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.best
}
