package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	game, err := NewGameFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}

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
	p := NewPlan(game.field, game.numStone)
base:
	for _, sBase := range game.stoneBase {
		for x := 0; x < 32; x++ {
			for y := 0; y < 32; y++ {
				for _, stone := range sBase.GetVariations() {
					if p.Put(x, y, stone) {
						continue base
					}
				}
			}
		}
	}
	return p
}
