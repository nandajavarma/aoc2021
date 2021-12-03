package day03

import (
	"testing"
)

func TestCountGammaAndEpsilon(t *testing.T) {
	tests := []struct {
		input []string
		want  int64
	}{
		{
			input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want:  198,
		},
	}

	for _, test := range tests {
		got := countGammaEpsilon(test.input)

		if got != test.want {
			t.Fatalf("TestCountGammaAndEpsilon failed: got %d want %d", got, test.want)
		}

	}
}

func TestCountOxygenCo2(t *testing.T) {
	tests := []struct {
		input []string
		want  int64
	}{
		{
			input: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want:  230,
		},
	}

	for _, test := range tests {
		got := countOxygenCo2(test.input)

		if got != test.want {
			t.Fatalf("TestCountOxygenCo2 failed: got %d want %d", got, test.want)
		}

	}
}
