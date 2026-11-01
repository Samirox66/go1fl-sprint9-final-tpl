package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "negative size", size: -5},
		{name: "one element", size: 1},
		{name: "many elements", size: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if tt.size <= 0 {
				if data != nil {
					t.Errorf("expected nil slice for size %d, got %v", tt.size, data)
				}
				return
			}

			if len(data) != tt.size {
				t.Errorf("expected slice of length %d, got %d", tt.size, len(data))
			}

			for _, value := range data {
				if value <= 0 {
					t.Errorf("expected positive value, got %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "max at the start", data: []int{10, 3, 7, 1}, want: 10},
		{name: "max at the end", data: []int{1, 3, 7, 10}, want: 10},
		{name: "max in the middle", data: []int{1, 10, 3}, want: 10},
		{name: "equal elements", data: []int{5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "fewer elements than chunks", data: []int{1, 2, 3}, want: 3},
		{name: "exactly divisible by chunks", data: []int{8, 1, 2, 3, 4, 5, 6, 7}, want: 8},
		{name: "not evenly divisible", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks() = %d, want %d", got, tt.want)
			}
		})
	}
}
