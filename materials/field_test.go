package materials

import (
	"github.com/umisama/procon26/buffer"
	"testing"
)

func TestNewField(t *testing.T) {
	cases := []struct {
		title  string
		input  []string
		expect *Field
	}{{
		title: "No.1",
		input: []string{
			"00000000000000001111111111111111",
			"00000000000000001111111111111111",
			"01000000000000001111111111111111",
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
		},
		expect: &Field{
			buffer: buffer.Buffer{
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
				buffer.Line{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			},
		},
	}}
	for _, c := range cases {
		t.Log("start: ", c.title)
		field, err := NewField(c.input)
		if err != nil {
			t.Fatal("error")
		}

		if field.buffer.String() != c.expect.buffer.String() {
			t.Log(field.buffer.String())
			t.Log(c.expect.buffer.String())
			t.Error("")
		}
	}
}
