package day06

import (
	"testing"
)

func TestTotalFish(t *testing.T) {
	tests := []struct {
		initialState string
		day          int
		want         int
	}{
		{
			initialState: "3,4,3,1,2",
			day:          18,
			want:         26,
		},
		{
			initialState: "3,4,3,1,2",
			day:          80,
			want:         5934,
		},
	}

	for _, tt := range tests {
		got := getTotalFish(tt.initialState, tt.day)
		if got != tt.want {
			t.Fatalf("TestTotalFish test failed: got %d, want %d", got, tt.want)
		}

	}
}
