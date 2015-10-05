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
		positions: make([]*Position, 0, 32),
		numStone:  numStone,
	}
}

func (plan *Plan) Copy() *Plan {
	pos := make([]*Position, len(plan.positions))
	copy(pos, plan.positions)
	return &Plan{
		field:     plan.field,
		positions: pos,
		numStone:  plan.numStone,
	}
}

func (plan *Plan) Get(x, y int) bool {
	if plan.field.Get(x, y) {
		return true
	}
	return plan.GetStoneDot(x, y)
}

func (plan *Plan) GetStoneDot(x, y int) bool {
	for _, position := range plan.positions {
		if position.Get(x, y) {
			return true
		}
	}
	return false
}

func (plan *Plan) Pop() {
	plan.positions = plan.positions[0 : len(plan.positions)-1]
}

func (plan *Plan) Put(x, y int, stone *Stone) bool {
	if !plan.puttable(x, y, stone) {
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

func (plan *Plan) puttable(x, y int, stone *Stone) bool {
	if plan.isDuplicateStone(stone) {
		return false
	}
	if !plan.canPutStone(x, y, stone) {
		return false
	}
	if !plan.isFirstStone() {
		if !plan.isExistRelatedStone(x, y, stone) {
			return false
		}
	}
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

func (plan *Plan) isFirstStone() bool {
	return len(plan.positions) == 0
}

func (plan *Plan) isExistRelatedStone(x, y int, stone *Stone) bool {
	for stoneX := 0; stoneX < stone.Width(); stoneX++ {
		for stoneY := 0; stoneY < stone.Height(); stoneY++ {
			if !stone.Get(stoneX, stoneY) {
				continue
			}
			if plan.GetStoneDot(stoneX+x-1, stoneY+y) ||
				plan.GetStoneDot(stoneX+x, stoneY+y-1) ||
				plan.GetStoneDot(stoneX+x+1, stoneY+y) ||
				plan.GetStoneDot(stoneX+x, stoneY+y+1) {
				return true
			}
		}
	}
	return false
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
	if plan == nil {
		return 0x8fffffff
	}
	score := plan.field.buffer.Height()*plan.field.buffer.Width() - plan.field.buffer.Count()
	for _, pos := range plan.positions {
		score -= pos.stone.Count()
	}
	return score
}

func (plan *Plan) PartialScore(rect Rect) int {
	if plan == nil {
		return 0x8fffffff
	}
	score := 0
	for x := rect.X; x < rect.X+rect.Width; x++ {
		for y := rect.Y; y < rect.Y+rect.Height; y++ {
			if !plan.Get(x, y) {
				score += 1
			}
		}
	}
	return score
}

func (plan *Plan) PartialScoreByExistStones() int {
	if plan == nil {
		return 0x8fffffff
	}
	score := 0
	for _, pos := range plan.positions {
		for x := pos.x; x < pos.x+pos.stone.Width(); x++ {
			for y := pos.y; y < pos.y+pos.stone.Height(); y++ {
				if !plan.Get(x, y) {
					score += 1
				}
			}
		}
	}
	return score
}

func (plan *Plan) CountIsolation() int {
	score := 0
	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			if !plan.Get(x, y) &&
				plan.Get(x-1, y) &&
				plan.Get(x+1, y) &&
				plan.Get(x, y-1) &&
				plan.Get(x, y+1) {
				score += 1
			}
		}
	}
	return score
}

func (plan *Plan) NumberOfPiece() int {
	return len(plan.positions)
}
