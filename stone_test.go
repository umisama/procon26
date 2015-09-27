package main

import (
	"reflect"
	"testing"
)

func TestNewStoneBase(t *testing.T) {
	cases := []struct {
		title  string
		input  []string
		expect *StoneBase
	}{{
		title: "shape likes L",
		input: []string{
			"01000000",
			"01000000",
			"01000000",
			"01000000",
			"01000000",
			"01000000",
			"01110000",
			"00000000",
		},
		expect: &StoneBase{
			buffer: Buffer{
				{true, false, false},
				{true, false, false},
				{true, false, false},
				{true, false, false},
				{true, false, false},
				{true, false, false},
				{true, true, true},
			},
			rect: Rect{X: 1, Y: 0, Width: 3, Height: 7},
		},
	}, {
		title: "shape likes T",
		input: []string{
			"00000000",
			"00111110",
			"00001000",
			"00001000",
			"00001000",
			"00001000",
			"00001000",
			"00001000",
		},
		expect: &StoneBase{
			buffer: Buffer{
				{true, true, true, true, true},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
				{false, false, true, false, false},
			},
			rect: Rect{X: 2, Y: 1, Width: 5, Height: 7},
		},
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		m, err := NewStoneBase(0, c.input)
		if err != nil {
			t.Fatalf("error: %s", err.Error())
		}

		if !reflect.DeepEqual(m.buffer, c.expect.buffer) {
			t.Errorf("expect 1 but got 2\n1:%v\n2:%v", *c.expect, *(m))
		}
		if !reflect.DeepEqual(m.rect, c.expect.rect) {
			t.Errorf("expect 1 but got 2\n1:%v\n2:%v", *c.expect, *(m))
		}
	}
}

func TestStoneBaseVariations(t *testing.T) {
	cases := []struct {
		title  string
		input  *StoneBase
		expect []*Stone
	}{{
		title: "No.1",
		input: &StoneBase{
			buffer: Buffer{
				{true, false},
				{true, true},
			},
			rect: Rect{X: 1, Y: 0, Width: 2, Height: 2},
		},
		expect: []*Stone{{
			buffer: Buffer{
				{true, false},
				{true, true},
			},
			rect:    Rect{X: 1, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: false,
		}, {
			buffer: Buffer{
				{false, true},
				{true, true},
			},
			rect:    Rect{X: 5, Y: 0, Width: 2, Height: 2},
			dig:     0,
			flipped: true,
		}, {
			buffer: Buffer{
				{true, true},
				{false, true},
			},
			rect:    Rect{X: 0, Y: 1, Width: 2, Height: 2},
			dig:     1,
			flipped: true,
		}, {
			buffer: Buffer{
				{true, true},
				{true, false},
			},
			rect:    Rect{X: 1, Y: 6, Width: 2, Height: 2},
			dig:     2,
			flipped: true,
		}},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		c.input.createVariations()
		val := c.input.GetVariations()
		if len(val) != len(c.expect) {
			t.Errorf("expect %d but got %d.", len(val), len(c.expect))
		}
		for i := 0; i < len(val); i++ {
			if val[i].String() != c.expect[i].String() {
				t.Errorf("expect 1 but got 2. \n1: %s\n2: %s\non %d", c.expect[i], val[i], i)
			}
		}
	}
}
