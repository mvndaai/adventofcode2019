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
