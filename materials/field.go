package materials

import (
	"github.com/umisama/procon26/buffer"
)

type Field struct {
	buffer buffer.Buffer
}

func NewField(input []string) (*Field, error) {
	buf, err := buffer.NewBufferByInput(input)
	if err != nil {
		return nil, err
	}

	return &Field{
		buffer: buf,
	}, nil
}

func (field *Field) Get(x, y int) bool {
	if len(field.buffer) <= y || y < 0 {
		return true
	}
	if len(field.buffer[0]) <= x || x < 0 {
		return true
	}
	return field.buffer[y][x]
}

func (field *Field) Width() int {
	return field.buffer.Width()
}

func (field *Field) Height() int {
	return field.buffer.Height()
}

func (field *Field) Count() int {
	return field.buffer.Count()
}
