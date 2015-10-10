package plan

import (
	"fmt"
	"github.com/umisama/procon26/buffer"
	"github.com/umisama/procon26/materials"
)

type Plan struct {
	field     *materials.Field
	positions []*Position

	buffer   buffer.Buffer
	numStone int
}

type Position struct {
	x, y  int
	stone *materials.Stone
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

func NewPlan(field *materials.Field, numStone int) *Plan {
	p := &Plan{
		field:     field,
		positions: make([]*Position, 0, 32),
		numStone:  numStone,
		buffer:    buffer.NewBuffer(field.Width(), field.Height()),
	}
	p.refreshBuffer(buffer.Rect{0, 0, field.Height(), field.Width()})
	return p
}

func (plan *Plan) Copy() *Plan {
	pos := make([]*Position, len(plan.positions))
	copy(pos, plan.positions)
	return &Plan{
		field:     plan.field,
		positions: pos,
		numStone:  plan.numStone,
		buffer:    plan.buffer.Copy(),
	}
}

func (plan *Plan) Get(x, y int) bool {
	return plan.buffer.Get(x, y)
}

func (plan *Plan) strictGet(x, y int) bool {
	return plan.field.Get(x, y) || plan.GetStoneDot(x, y)
}

func (plan *Plan) refreshBuffer(rect buffer.Rect) {
	for x := rect.X; x < rect.X+rect.Width; x++ {
		for y := rect.Y; y < rect.Y+rect.Height; y++ {
			plan.buffer.Set(x, y, plan.strictGet(x, y))
		}
	}
	return
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
	pos := plan.positions[len(plan.positions)-1]
	plan.positions = plan.positions[0 : len(plan.positions)-1]
	plan.refreshBuffer(buffer.Rect{pos.x, pos.y, pos.stone.Height(), pos.stone.Width()})
	return
}

func (plan *Plan) Put(x, y int, stone *materials.Stone) bool {
	if !plan.puttable(x, y, stone) {
		return false
	}
	// put stone
	plan.positions = append(plan.positions, &Position{
		x:     x,
		y:     y,
		stone: stone,
	})
	plan.refreshBuffer(buffer.Rect{x, y, stone.Height(), stone.Width()})
	return true
}

func (plan *Plan) puttable(x, y int, stone *materials.Stone) bool {
	if !plan.canPutStone(x, y, stone) {
		return false
	}
	if plan.isDuplicateStone(stone) {
		return false
	}
	if !plan.isFirstStone() {
		if !plan.isExistRelatedStone(x, y, stone) {
			return false
		}
	}
	return true
}

func (plan *Plan) isDuplicateStone(stone *materials.Stone) bool {
	for _, pos := range plan.positions {
		if pos.stone.Number() == stone.Number() {
			return true
		}
	}
	return false
}

func (plan *Plan) canPutStone(x, y int, stone *materials.Stone) bool {
	for stoneX := 0; stoneX < stone.Width(); stoneX++ {
		for stoneY := 0; stoneY < stone.Height(); stoneY++ {
			if !stone.Get(stoneX, stoneY) {
				continue
			}
			if plan.GetStoneDot(x+stoneX, y+stoneY) {
				return false
			}
		}
	}
	return true
}

func (plan *Plan) isFirstStone() bool {
	return len(plan.positions) == 0
}

func (plan *Plan) isExistRelatedStone(x, y int, stone *materials.Stone) bool {
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
	str, first := "", false
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

		str += fmt.Sprintf("%d %d ", position.x-stone.Rect().X, position.y-stone.Rect().Y)
		if stone.IsFlipped() {
			str += "T "
		} else {
			str += "H "
		}
		str += fmt.Sprintf("%d", stone.Dig())
	}
	str += "\n"
	return str
}

func (plan *Plan) findPositionByStoneNumber(num int) *Position {
	for _, position := range plan.positions {
		if position.stone.Number() == num {
			return position
		}
	}
	return nil
}

func (plan *Plan) Score() int {
	if plan == nil {
		return 0x8fffffff
	}
	score := plan.field.Height()*plan.field.Width() - plan.field.Count()
	for _, pos := range plan.positions {
		score -= pos.stone.Count()
	}
	return score
}

func (plan *Plan) PartialScore(rect buffer.Rect) int {
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
