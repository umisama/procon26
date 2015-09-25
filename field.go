package main

type Field struct {
	buffer Buffer
}

func NewField(input []string) (*Field, error) {
	buf, err := NewBufferByInput(input)
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
