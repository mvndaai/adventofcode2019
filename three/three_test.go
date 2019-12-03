package main

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	tests := []struct {
		a string
		b string
		// out   int
		steps int
	}{
		{
			a: "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			b: "U62,R66,U55,R34,D71,R55,D58,R83",
			// out: 159,
			steps: 610,
		},
		{
			a: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			b: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			// out: 135,
			steps: 410,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.steps), func(t *testing.T) {
			if out := manhatten(tt.a, tt.b); out != tt.steps {
				t.Error("out did not match", out)
			}
		})
	}
}
