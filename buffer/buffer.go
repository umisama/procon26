package buffer

import (
	"fmt"
)

type Rect struct {
	X, Y          int
	Height, Width int
}

func (m Rect) Rotate(boxWidth int) Rect {
	return Rect{
		X:      m.Y,
		Y:      boxWidth - m.X - m.Width,
		Height: m.Width,
		Width:  m.Height,
	}
}

func (m Rect) Flip(boxWidth int) Rect {
	return Rect{
		X:      boxWidth - m.Width - m.X,
		Y:      m.Y,
		Height: m.Height,
		Width:  m.Width,
	}
}

type Line []bool

func NewLineByInput(input string) (Line, error) {
	line := make(Line, len(input))
	for i, c := range input {
		switch c {
		case '0':
			line[i] = false
		case '1':
			line[i] = true
		default:
			return nil, fmt.Errorf("0 or 1 expected, but got %c", c)
		}
	}
	return line, nil
}

func NewLine(length int) Line {
	return make(Line, length)
}

func (line Line) Get(x int) bool {
	if len(line) <= x || 0 > x {
		return true
	}
	return line[x]
}

func (line Line) Set(x int, val bool) {
	if len(line) <= x || 0 > x {
		return
	}
	line[x] = val
	return
}

func (line Line) GetRange() (start, end int) {
	start, end = -1, -1
	for pos, dot := range line {
		if dot {
			if start == -1 {
				start = pos
			}
			end = pos
		}
	}
	return start, end
}

func (line Line) Count() int {
	cnt := 0
	for i := 0; i < len(line); i++ {
		if line[i] {
			cnt += 1
		}
	}
	return cnt
}

func (line Line) IsEmpty() bool {
	b := false
	for _, dot := range line {
		if dot {
			b = true
			break
		}
	}
	return !b
}

func (line Line) Trim(start, end int) Line {
	nbuf := NewLine(end - start)
	if start > line.Len() || end > line.Len() || end < start {
		return Line{}
	}
	for i := 0; i < end-start; i++ {
		nbuf[i] = line[i+start]
	}
	return nbuf
}

func (line Line) String() string {
	str := ""
	for _, d := range line {
		if d {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func (line Line) Len() int {
	return len(line)
}

type Buffer []Line

func NewBufferByInput(lines []string) (Buffer, error) {
	buf := make(Buffer, len(lines))
	for i, lStr := range lines {
		if len(lStr) != len(lines[0]) {
			return nil, fmt.Errorf("invalid length in lines")
		}
		line, err := NewLineByInput(lStr)
		if err != nil {
			return nil, err
		}
		buf[i] = line
	}
	return buf, nil
}

func NewBuffer(width, height int) Buffer {
	buf := make(Buffer, height)
	for i := 0; i < height; i++ {
		buf[i] = NewLine(width)
	}
	return buf
}

func (buf Buffer) Get(x, y int) bool {
	if len(buf) <= y || 0 > y {
		return true
	}
	return buf[y].Get(x)
}

func (buf Buffer) Set(x, y int, val bool) {
	if len(buf) <= y || 0 > y {
		return
	}
	buf[y].Set(x, val)
	return
}

func (buf Buffer) GetRect() Rect {
	var startx, starty, endx, endy = -1, -1, -1, -1
	for pos, line := range buf {
		if line.IsEmpty() {
			continue
		}

		if starty == -1 {
			starty = pos
		}
		endy = pos

		linestart, lineend := line.GetRange()
		if linestart < startx || startx == -1 {
			startx = linestart
		}
		if lineend > endx || endx == -1 {
			endx = lineend
		}
	}

	if startx == -1 || starty == -1 {
		return Rect{
			X:      0,
			Y:      0,
			Width:  0,
			Height: 0,
		}
	} else {
		return Rect{
			X:      startx,
			Y:      starty,
			Width:  endx - startx + 1,
			Height: endy - starty + 1,
		}
	}
}

func (buf Buffer) Trim(rect Rect) Buffer {
	nBuf := NewBuffer(rect.Width, rect.Height)
	nLnum := 0

	for lnum := rect.Y; lnum < rect.Y+rect.Height; lnum++ {
		nBuf[nLnum] = buf[lnum].Trim(rect.X, rect.X+rect.Width)
		nLnum += 1
	}
	return nBuf
}

func (buf Buffer) String() string {
	str := ""
	for _, l := range buf {
		str += l.String() + "\n"
	}
	return str
}

func (buf Buffer) Height() int {
	return len(buf)
}

func (buf Buffer) Width() int {
	if len(buf) == 0 {
		return 0
	}
	return buf[0].Len()
}

func (buf Buffer) Rotate() Buffer {
	nBuf := NewBuffer(buf.Height(), buf.Width())
	for x := 0; x < buf.Width(); x++ {
		for y := 0; y < buf.Height(); y++ {
			nBuf[buf.Width()-x-1][y] = buf[y][x]
		}
	}
	return nBuf
}

func (buf Buffer) Flip() Buffer {
	nBuf := NewBuffer(buf.Width(), buf.Height())
	for x := 0; x < buf.Width(); x++ {
		for y := 0; y < buf.Height(); y++ {
			nBuf[y][buf.Width()-x-1] = buf[y][x]
		}
	}
	return nBuf
}

func (buf Buffer) Count() int {
	cnt := 0
	for i := 0; i < len(buf); i++ {
		cnt += buf[i].Count()
	}
	return cnt
}

func (buf Buffer) Copy() Buffer {
	nbuf := make(Buffer, buf.Height())
	for y, line := range buf {
		nline := make(Line, len(line))
		copy(nline, line)
		nbuf[y] = nline
	}
	return nbuf
}
