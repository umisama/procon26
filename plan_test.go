package main

import (
	"testing"
)

func TestPlanGet(t *testing.T) {
	field, err := NewField([]string{
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
		},
		inputX: 0, inputY: 0,
		expect: false,
	}, {
		title: "returns true if dot exists in field.",
		m: &Plan{
			field:     field,
			positions: nil,
		},
		inputX: 2, inputY: 2,
		expect: true,
	}, {
		title: "returns true if dot exists in stone.(1)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		inputX: 0, inputY: 0,
		expect: true,
	}, {
		title: "returns true if dot exists in stone.(2)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		inputX: 1, inputY: 1,
		expect: true,
	}, {
		title: "returns false if dot don't exists in stone and field.(1)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		inputX: 1, inputY: 0,
		expect: false,
	}, {
		title: "returns false if dot don't exists in stone and field.(2)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		inputX: 3, inputY: 1,
		expect: false,
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		if c.m.Get(c.inputX, c.inputY) != c.expect {
			t.Error("failed")
		}
	}
}

func TestPlanPut(t *testing.T) {
	field, err := NewField([]string{
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
		stone          *Stone
		inputX, inputY int
		expect         bool
	}{{
		title: "allow if plan is empty",
		m: &Plan{
			field:     field,
			positions: nil,
		},
		stone: &Stone{
			buffer: Buffer{
				{true, false},
				{true, true},
			},
			rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		},
		inputX: 0, inputY: 0,
		expect: true,
	}, {
		title: "allow if there is not duplicated stone",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					number: 1,
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		stone: &Stone{
			number: 2,
			buffer: Buffer{
				{true, false},
				{true, true},
			},
			rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		},
		inputX: 2, inputY: 0,
		expect: true,
	}, {
		title: "deny if there is not duplicated stone(layerd)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					number: 1,
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		stone: &Stone{
			number: 2,
			buffer: Buffer{
				{true, true},
				{false, true},
			},
			rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		},
		inputX: 1, inputY: 0,
		expect: true,
	}, {
		title: "allow if there is not duplicated stone(layerd)",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					number: 1,
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		stone: &Stone{
			number: 2,
			buffer: Buffer{
				{true, true},
				{false, true},
			},
			rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		},
		inputX: 1, inputY: 0,
		expect: true,
	}, {
		title: "deny if there is not related stone.",
		m: &Plan{
			field: field,
			positions: []*Position{{
				x: 0, y: 0,
				stone: &Stone{
					number: 1,
					buffer: Buffer{
						{true, false},
						{true, true},
					},
					rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
					dig:     0,
					flipped: false,
				},
			}},
		},
		stone: &Stone{
			number: 2,
			buffer: Buffer{
				{false, true},
				{true, true},
			},
			rect:    Rect{X: 0, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		},
		inputX: 2, inputY: 2,
		expect: false,
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		if c.m.Put(c.inputX, c.inputY, c.stone) != c.expect {
			t.Error("failed")
		}
	}
}

func BenchmarkPlanScore(b *testing.B) {
	field, _ := NewField([]string{
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
	}
	for i := 0; i < b.N; i++ {
		plan.Score()
	}
}
