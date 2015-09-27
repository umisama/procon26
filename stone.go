package main

import (
	"fmt"
	"reflect"
)

func NewStoneBase(number int, input []string) (*StoneBase, error) {
	buf, err := NewBufferByInput(input)
	if err != nil {
		return nil, err
	}
	rect := buf.GetRect()
	trimmedBuf := buf.Trim(rect)

	m := &StoneBase{
		number: number,
		buffer: trimmedBuf,
		rect:   rect,
	}

	m.createVariations()
	return m, nil
}

type StoneBase struct {
	number     int
	buffer     Buffer
	rect       Rect // A trimmed rect in input
	variations []*Stone
}

func (m *StoneBase) Number() int {
	return m.number
}

func (m *StoneBase) Get(x, y int) bool {
	return m.buffer.Get(x, y)
}

func (m *StoneBase) Height() int {
	return m.rect.Height
}

func (m *StoneBase) checkHeight() int {
	height := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if m.buffer.Get(x, y) {
				height = x
				break
			}
		}
	}
	return height + 1
}

func (m *StoneBase) Width() int {
	return m.rect.Width
}

func (m *StoneBase) checkWidth() int {
	width := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if m.buffer.Get(x, y) {
				width = y
				break
			}
		}
	}
	return width + 1
}

func (m *StoneBase) GetVariations() []*Stone {
	return m.variations
}

func (m *StoneBase) createVariations() {
	variations := make([]*Stone, 0, 8)

	buf := m.buffer
	rect := m.rect
	bufF := m.buffer.Flip()
	rectF := m.rect.Flip(8)
	for i := 0; i < 4; i++ {
		variations = append(variations, &Stone{
			number: m.number,
			buffer: buf,
			rect:   rect,
			dig:    i,
		})
		variations = append(variations, &Stone{
			number:  m.number,
			buffer:  bufF,
			rect:    rectF,
			dig:     i,
			flipped: true,
		})

		buf = buf.Rotate()
		bufF = bufF.Rotate()
		rect = rect.Rotate(8)
		rectF = rectF.Rotate(8)
	}
	m.variations = removeDuplicateItems(variations)
	return
}

func removeDuplicateItems(stones []*Stone) []*Stone {
	ret := make([]*Stone, 0, len(stones))
	for i := 0; i < len(stones); i++ {
		found := false
		for j := 0; j < i; j++ {
			if reflect.DeepEqual(stones[i].buffer, stones[j].buffer) {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, stones[i])
		}
	}
	return ret
}

type Stone struct {
	number  int
	buffer  Buffer
	rect    Rect
	dig     int
	flipped bool
}

func (m *Stone) Number() int {
	return m.number
}

func (m *Stone) Get(x, y int) bool {
	return m.buffer.Get(x, y)
}

func (m *Stone) Height() int {
	return m.rect.Height
}

func (m *Stone) Width() int {
	return m.rect.Width
}

func (m *Stone) String() string {
	str := ""
	str += "stone:\n" + m.buffer.String()
	str += "dig:" + fmt.Sprintf("%d\n", m.dig)
	str += "flipped: " + fmt.Sprintf("%#v\n", m.flipped)
	str += "rect: " + fmt.Sprintf("%#v\n", m.rect)
	return str
}
