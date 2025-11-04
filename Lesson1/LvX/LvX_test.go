package main

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	cases := []struct {
		in1, in2 int
		want     float64
	}{
		{1, 1, 1},
		{1, 2, 0.5},
		{9, 3, 3},
	}
	for _, c := range cases {
		got := Average(c.in1, c.in2)
		if math.Abs(got-c.want) > 0.1 {
			t.Fatalf("got %.2f, want %.2f", got, c.want)
		}
	}
}
