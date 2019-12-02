package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{in: 12, out: 2},
		{in: 14, out: 2},
		{in: 1969, out: 654},
		{in: 100756, out: 33583},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			if f := fuel(tt.in); f != tt.out {
				t.Error("final fuel was: ", f)
			}
		})
	}
}

func TestOne2(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{in: 12, out: 2},
		{in: 14, out: 2},
		{in: 1969, out: 966},
		{in: 100756, out: 50346},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			if f := fuel2(tt.in); f != tt.out {
				t.Error("final fuel was: ", f)
			}
		})
	}
}
