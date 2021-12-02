package day01

import (
	"testing"
)

func TestCountIncreasingPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"},
			want:  7,
		},
	}

	for _, test := range tests {
		got := countIncreasing(test.input, 1)

		if got != test.want {
			t.Fatalf("TestCountIncreasingPart1 failed: got %d, want %d", got, test.want)
		}
	}
}

func TestCountIncreasingPart2(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"},
			want:  5,
		},
	}

	for _, test := range tests {
		got := countIncreasing(test.input, 3)

		if got != test.want {
			t.Fatalf("TestCountIncreasingPart3 failed: got %d, want %d", got, test.want)
		}
	}
}
