package plan

import (
	"github.com/umisama/procon26/buffer"
	"github.com/umisama/procon26/materials"
	"reflect"
	"testing"
)

func TestPlanGet(t *testing.T) {
	field, err := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	if err != nil {
		panic(err)
	}

	cases := []struct {
		title          string
		m              *Plan
		inputX, inputY int
		expect         bool
	}{{
		title: "return false if there are not dot.",
		m: &Plan{
			field:     field,
			positions: nil,
			buffer:    buffer.NewBuffer(32, 32),
		},
		inputX: 0, inputY: 0,
		expect: false,
	}, {
		title: "returns true if dot exists in field.",
		m: &Plan{
			field:     field,
			positions: nil,
			buffer:    buffer.NewBuffer(32, 32),
		},
		inputX: 2, inputY: 2,
		expect: true,
	}, {
		title: "returns true if dot exists in stone.(1)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		inputX: 0, inputY: 0,
		expect: true,
	}, {
		title: "returns true if dot exists in stone.(2)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		inputX: 1, inputY: 1,
		expect: true,
	}, {
		title: "returns false if dot don't exists in stone and field.(1)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		inputX: 1, inputY: 0,
		expect: false,
	}, {
		title: "returns false if dot don't exists in stone and field.(2)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		inputX: 3, inputY: 1,
		expect: false,
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		c.m.refreshBuffer(buffer.Rect{0, 0, 32, 32})
		if c.m.Get(c.inputX, c.inputY) != c.expect {
			t.Error("failed")
		}
	}
}

func TestPlanPut(t *testing.T) {
	field, err := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	if err != nil {
		panic(err)
	}

	cases := []struct {
		title          string
		m              *Plan
		stone          *materials.Stone
		inputX, inputY int
		expect         bool
	}{{
		title: "allow if plan is empty",
		m: &Plan{
			field:     field,
			buffer:    buffer.NewBuffer(32, 32),
			positions: nil,
		},
		stone: materials.NewStone(0, buffer.Buffer{
			{true, false},
			{true, true},
		}),
		inputX: 0, inputY: 0,
		expect: true,
	}, {
		title: "allow if there is not duplicated stone",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		stone: materials.NewStone(1, buffer.Buffer{
			{true, false},
			{true, true},
		}),
		inputX: 2, inputY: 0,
		expect: true,
	}, {
		title: "deny if there is not duplicated stone(layerd)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		stone: materials.NewStone(2, buffer.Buffer{
			{true, true},
			{false, true},
		}),
		inputX: 1, inputY: 0,
		expect: true,
	}, {
		title: "allow if there is not duplicated stone(layerd)",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		stone: materials.NewStone(2, buffer.Buffer{
			{true, true},
			{false, true},
		}),
		inputX: 1, inputY: 0,
		expect: true,
	}, {
		title: "deny if there is not related stone.",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		stone: materials.NewStone(2, buffer.Buffer{
			{false, true},
			{true, true},
		}),
		inputX: 2, inputY: 2,
		expect: false,
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		c.m.refreshBuffer(buffer.Rect{0, 0, 32, 32})
		if c.m.Put(c.inputX, c.inputY, c.stone) != c.expect {
			t.Error("failed")
		}
	}
}

func TestPlanPartialScoreByExistStones(t *testing.T) {
	field, _ := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	cases := []struct {
		title  string
		m      *Plan
		expect int
	}{{
		title: "No.1",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		expect: 1,
	}, {
		title: "No.2",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(0, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}, {
				x: 2, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, false},
					{true, true},
				}),
			}},
		},
		expect: 2,
	}}
	for _, c := range cases {
		c.m.refreshBuffer(buffer.Rect{0, 0, 32, 32})
		v := c.m.PartialScoreByExistStones()
		if v != c.expect {
			t.Error("failed: got:", v, " expect:", c.expect)
		}
	}
}

func TestPlanCountIsolation(t *testing.T) {
	field, _ := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	cases := []struct {
		title  string
		m      *Plan
		expect int
	}{{
		title: "No.1",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, false, true},
					{true, true, true},
				}),
			}},
		},
		expect: 1,
	}, {
		title: "No.2",
		m: &Plan{
			field:  field,
			buffer: buffer.NewBuffer(32, 32),
			positions: []*Position{{
				x: 0, y: 0,
				stone: materials.NewStone(1, buffer.Buffer{
					{true, true},
					{true, false},
					{true, true},
				}),
			}, {
				x: 2, y: 0,
				stone: materials.NewStone(2, buffer.Buffer{
					{true},
					{true},
				}),
			}},
		},
		expect: 1,
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		c.m.refreshBuffer(buffer.Rect{0, 0, 32, 32})
		v := c.m.CountIsolation()
		if c.expect != v {
			t.Error("failed")
		}
	}
}

func TestPlanPutPop(t *testing.T) {
	field, _ := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	plan := &Plan{
		field:  field,
		buffer: buffer.NewBuffer(32, 32),
		positions: []*Position{{
			x: 0, y: 0,
			stone: materials.NewStone(0, buffer.Buffer{
				{true, false, true},
				{true, true, true},
			}),
		}},
		numStone: 3,
	}
	plan.refreshBuffer(buffer.Rect{0, 0, 32, 32})

	planc := plan.Copy()

	for i := range plan.buffer {
		if !reflect.DeepEqual(plan.buffer[i], planc.buffer[i]) {
			t.Errorf("failed on %d", i)
		}
	}

	if !planc.Put(3, 0, materials.NewStone(2, buffer.Buffer{
		{true, true, true},
		{true, true, true},
		{true, true, false},
	})) {
		t.Errorf("failed")
	}
	planc.Pop()

	for i := range plan.buffer {
		if !reflect.DeepEqual(plan.buffer[i], planc.buffer[i]) {
			t.Errorf("failed on %d", i)
		}
	}
}

func BenchmarkPlanScore(b *testing.B) {
	field, _ := materials.NewField([]string{
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00100000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000000000001111111111111111",
		"00000000010000001111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
		"11111111111111111111111111111111",
	})
	plan := &Plan{
		field:     field,
		positions: nil,
		buffer:    buffer.NewBuffer(32, 32),
	}
	for i := 0; i < b.N; i++ {
		plan.refreshBuffer(buffer.Rect{0, 0, 32, 32})
		plan.Score()
	}
}
