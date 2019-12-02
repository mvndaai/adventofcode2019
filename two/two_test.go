package main

import (
	"fmt"
	"testing"
)

func TestTwo(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{in: []int{1, 0, 0, 0, 99}, out: []int{2, 0, 0, 0, 99}},
		{in: []int{2, 3, 0, 3, 99}, out: []int{2, 3, 0, 6, 99}},
		{in: []int{2, 4, 4, 5, 99, 0}, out: []int{2, 4, 4, 5, 99, 9801}},
		{in: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, out: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			ic := intCode(tt.in)
			if fmt.Sprint(ic) != fmt.Sprint(tt.out) {
				t.Error("final values didn't match", ic)
			}
		})
	}
}
