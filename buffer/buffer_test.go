package buffer

import (
	"reflect"
	"testing"
)

func TestLineGetRange(t *testing.T) {
	cases := []struct {
		title      string
		input      Line
		rangeStart int
		rangeEnd   int
	}{{
		title:      "When line is FULL",
		input:      Line{true, true, true, true, true, true, true, true},
		rangeStart: 0,
		rangeEnd:   7,
	}, {
		title:      "When dot on firlst",
		input:      Line{true, false, false, false, false, false, false, false},
		rangeStart: 0,
		rangeEnd:   0,
	}, {
		title:      "When dot on last",
		input:      Line{false, false, false, false, false, false, false, true},
		rangeStart: 7,
		rangeEnd:   7,
	}, {
		title:      "When line is EMPTY",
		input:      Line{false, false, false, false, false, false, false, false},
		rangeStart: -1,
		rangeEnd:   -1,
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		start, end := c.input.GetRange()
		if start != c.rangeStart {
			t.Errorf("expect %d but got %d", c.rangeStart, start)
		}
		if end != c.rangeEnd {
			t.Errorf("expect %d but got %d", c.rangeEnd, end)
		}
	}
}

func TestLineTrim(t *testing.T) {
	cases := []struct {
		title     string
		input     Line
		trimStart int
		trimEnd   int
		expect    Line
	}{{
		title:     "All range",
		input:     Line{true, true, true, true, true, true, true, true},
		trimStart: 0,
		trimEnd:   8,
		expect:    Line{true, true, true, true, true, true, true, true},
	}, {
		title:     "First item only",
		trimStart: 0,
		trimEnd:   1,
		input:     Line{true, false, false, false, false, false, false, false},
		expect:    Line{true},
	}, {
		title:     "Cut item",
		trimStart: 2,
		trimEnd:   5,
		input:     Line{false, false, true, true, true, false, false, false},
		expect:    Line{true, true, true},
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		val := c.input.Trim(c.trimStart, c.trimEnd)
		if !reflect.DeepEqual(val, c.expect) {
			t.Errorf("expect %s but got %s", c.expect, val)
		}
	}
}

func TestLineIsEmpty(t *testing.T) {
	cases := []struct {
		title string
		input Line
		empty bool
	}{{
		title: "When line is FULL",
		input: Line{true, true, true, true, true, true, true, true},
		empty: false,
	}, {
		title: "When one dot in line",
		input: Line{false, true, false, false, false, false, false, false},
		empty: false,
	}, {
		title: "When line is EMPTY",
		input: Line{false, false, false, false, false, false, false, false},
		empty: true,
	}}

	for _, c := range cases {
		t.Log("start: ", c.title)
		if c.input.IsEmpty() != c.empty {
			t.Errorf("expect %#v but got %#v", c.empty, c.input.IsEmpty())
		}
	}
}

func TestBufferGetRect(t *testing.T) {
	cases := []struct {
		title  string
		input  Buffer
		expect Rect
	}{{
		title: "FULL",
		input: Buffer{
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
		},
		expect: Rect{X: 0, Y: 0, Width: 8, Height: 8},
	}, {
		title: "shape likes L",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, true, true, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		expect: Rect{X: 1, Y: 1, Width: 3, Height: 6},
	}, {
		title: "shape likes T",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, true, true, true, true, true, false, false},
			Line{false, false, true, false, false, false, false, false},
			Line{false, false, true, false, false, false, false, false},
			Line{false, false, true, false, false, false, false, false},
			Line{false, false, true, false, false, false, false, false},
			Line{false, false, true, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		expect: Rect{X: 1, Y: 1, Width: 5, Height: 6},
	}, {
		title: "EMPTY",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		expect: Rect{X: 0, Y: 0, Width: 0, Height: 0},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		if c.input.GetRect() != c.expect {
			t.Errorf("expect %#v but got %#v", c.expect, c.input.GetRect())
		}
	}
}

func TestBufferTrim(t *testing.T) {
	cases := []struct {
		title  string
		input  Buffer
		rect   Rect
		expect Buffer
	}{{
		title: "FULL",
		input: Buffer{
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
		},
		rect: Rect{X: 0, Y: 0, Width: 8, Height: 8},
		expect: Buffer{
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
			Line{true, true, true, true, true, true, true, true},
		},
	}, {
		title: "shape likes L",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, false, false, false, false, false, false},
			Line{false, true, true, true, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		rect: Rect{X: 1, Y: 1, Width: 3, Height: 6},
		expect: Buffer{
			Line{true, false, false},
			Line{true, false, false},
			Line{true, false, false},
			Line{true, false, false},
			Line{true, false, false},
			Line{true, true, true},
		},
	}, {
		title: "shape likes T",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, true, true, true, true, true, false, false},
			Line{false, false, false, true, false, false, false, false},
			Line{false, false, false, true, false, false, false, false},
			Line{false, false, false, true, false, false, false, false},
			Line{false, false, false, true, false, false, false, false},
			Line{false, false, false, true, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		rect: Rect{X: 1, Y: 1, Width: 5, Height: 6},
		expect: Buffer{
			Line{true, true, true, true, true},
			Line{false, false, true, false, false},
			Line{false, false, true, false, false},
			Line{false, false, true, false, false},
			Line{false, false, true, false, false},
			Line{false, false, true, false, false},
		},
	}, {
		title: "EMPTY",
		input: Buffer{
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
			Line{false, false, false, false, false, false, false, false},
		},
		rect:   Rect{X: 0, Y: 0, Width: 0, Height: 0},
		expect: Buffer{},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		val := c.input.Trim(c.rect)
		if !reflect.DeepEqual(val, c.expect) {
			t.Errorf("expect %s but got %s", c.expect, val)
		}
	}
}

func TestBufferRotate(t *testing.T) {
	cases := []struct {
		title  string
		input  Buffer
		expect Buffer
	}{{
		title: "No.1",
		input: Buffer{
			Line{true, true, true},
			Line{false, false, false},
		},
		expect: Buffer{
			Line{true, false},
			Line{true, false},
			Line{true, false},
		},
	}, {
		title: "No.2",
		input: Buffer{
			Line{true, false},
			Line{true, false},
			Line{true, false},
		},
		expect: Buffer{
			Line{false, false, false},
			Line{true, true, true},
		},
	}, {
		title: "No.3",
		input: Buffer{
			Line{false, false, false},
			Line{true, true, true},
		},
		expect: Buffer{
			Line{false, true},
			Line{false, true},
			Line{false, true},
		},
	}, {
		title: "No.4",
		input: Buffer{
			Line{false, true},
			Line{false, true},
			Line{false, true},
		},
		expect: Buffer{
			Line{true, true, true},
			Line{false, false, false},
		},
	}, {
		title: "No.5",
		input: Buffer{
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
		},
		expect: Buffer{
			Line{false, false, false, false},
			Line{false, false, false, false},
			Line{true, true, true, true},
			Line{false, false, false, false},
		},
	}, {
		title: "No.6",
		input: Buffer{
			Line{false, false, false, false},
			Line{false, false, false, false},
			Line{true, true, true, true},
			Line{false, false, false, false},
		},
		expect: Buffer{
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
		},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		val := c.input.Rotate()
		if !reflect.DeepEqual(val, c.expect) {
			t.Errorf("expect %s but got %s", c.expect, val)
		}
	}
}

func TestBufferFlip(t *testing.T) {
	cases := []struct {
		title  string
		input  Buffer
		expect Buffer
	}{{
		title: "No.1",
		input: Buffer{
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
		},
		expect: Buffer{
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
		},
	}, {
		title: "No.2",
		input: Buffer{
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
			Line{false, false, true, false},
		},
		expect: Buffer{
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
			Line{false, true, false, false},
		},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		val := c.input.Flip()
		if !reflect.DeepEqual(val, c.expect) {
			t.Errorf("expect %s but got %s", c.expect, val)
		}
	}
}

func TestBufferCopy(t *testing.T) {
	input := Buffer{
		Line{false, true, false, false},
		Line{false, true, false, false},
		Line{false, true, false, false},
		Line{false, true, false, false},
	}

	input2 := input.Copy()
	input2[0][0] = true
	input2[1][0] = true
	input2[2][0] = true
	input2[3][0] = true
	for k := range input {
		if reflect.DeepEqual(input[k], input2[k]) {
			t.Error("error on ", k)
			t.Log("input :", input[k].String())
			t.Log("copied:", input2[k].String())
		}
	}

	input[0][0] = true
	input[1][0] = true
	input[2][0] = true
	input[3][0] = true
	for k := range input {
		if !reflect.DeepEqual(input[k], input2[k]) {
			t.Error("error on ", k)
			t.Log("input :", input[k].String())
			t.Log("copied:", input2[k].String())
		}
	}
}
