package main

import (
	"fmt"
)

type Plan struct {
	field     *Field
	positions []*Position
	numStone  int
}

type Position struct {
	x, y  int
	stone *Stone
}

func (position *Position) Get(x, y int) bool {
	posx := x - position.x
	posy := y - position.y
	if position.stone.Width() <= posx || posx < 0 {
		return false
	}
	if position.stone.Height() <= posy || posy < 0 {
		return false
	}
	return position.stone.Get(posx, posy)
}

func NewPlan(field *Field, numStone int) *Plan {
	return &Plan{
		field:     field,
		positions: make([]*Position, 0),
		numStone:  numStone,
	}
}

func (plan *Plan) Get(x, y int) bool {
	if plan.field.Get(x, y) {
		return true
	}
	for _, position := range plan.positions {
		if position.Get(x, y) {
			return true
		}
	}
	return false
}

func (plan *Plan) Put(x, y int, stone *Stone) bool {
	if plan.isDuplicateStone(stone) {
		return false
	}
	if !plan.canPutStone(x, y, stone) {
		return false
	}

	// put stone
	plan.positions = append(plan.positions, &Position{
		x:     x,
		y:     y,
		stone: stone,
	})
	return true
}

func (plan *Plan) isDuplicateStone(stone *Stone) bool {
	for _, pos := range plan.positions {
		if pos.stone.Number() == stone.Number() {
			return true
		}
	}
	return false
}

func (plan *Plan) canPutStone(x, y int, stone *Stone) bool {
	for stoneX := 0; stoneX < stone.Width(); stoneX++ {
		for stoneY := 0; stoneY < stone.Height(); stoneY++ {
			if !stone.Get(stoneX, stoneY) {
				continue
			}
			if plan.Get(x+stoneX, y+stoneY) {
				return false
			}
		}
	}
	return true
}

func (plan *Plan) String() string {
	str := ""

	first := false
	for i := 0; i < plan.numStone; i++ {
		if !first {
			first = true
		} else {
			str += "\n"
		}

		// find
		position := plan.findPositionByStoneNumber(i)
		if position == nil {
			continue
		}
		stone := position.stone

		str += fmt.Sprintf("%d %d ", position.x-stone.rect.X, position.y-stone.rect.Y)
		if stone.flipped {
			str += "T "
		} else {
			str += "H "
		}

		switch stone.dig {
		case 0:
			str += "0"
		case 1:
			str += "270"
		case 2:
			str += "180"
		case 3:
			str += "90"
		}
	}
	return str
}

func (plan *Plan) findPositionByStoneNumber(num int) *Position {
	for _, position := range plan.positions {
		if position.stone.number == num {
			return position
		}
	}
	return nil
}

func (plan *Plan) Score() int {
	score := 0
	for i := 0; i < len(plan.field.buffer); i++ {
		for j := 0; j < len(plan.field.buffer[0]); j++ {
			if !plan.Get(i, j) {
				score += 1
			}
		}
	}
	return score
}
