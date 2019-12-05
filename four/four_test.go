package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{in: 111111, out: true},
		{in: 223450, out: false},
		{in: 123789, out: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			if validNumber(tt.in) != tt.out {
				t.Error("did not match")
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{in: 112233, out: true},
		{in: 123444, out: false},
		{in: 111122, out: true},

		{in: 111111, out: false},
		{in: 223450, out: false},
		{in: 123789, out: false},

		{in: 137683, out: false},
		{in: 596253, out: false},

		{in: 577999, out: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			if validNumber2(tt.in) != tt.out {
				t.Error("did not match")
			}
		})
	}
}
