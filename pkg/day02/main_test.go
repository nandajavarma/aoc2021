package day02

import (
	"testing"
)

func TestGetPositionPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			want:  150,
		},
	}

	for _, test := range tests {
		got := get_position(test.input, false)

		if got != test.want {
			t.Fatalf("TestGetPositionPart1 failed: want: %d got %d", test.want, got)
		}
	}
}

func TestGetPositionPart2(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			want:  900,
		},
	}

	for _, test := range tests {
		got := get_position(test.input, true)

		if got != test.want {
			t.Fatalf("TestGetPositionPart2 failed: want: %d got %d", test.want, got)
		}
	}
}
