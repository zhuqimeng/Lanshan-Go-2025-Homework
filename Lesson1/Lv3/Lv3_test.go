package main

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
	}
	for _, c := range cases {
		got := Factorial(c.in)
		if got != c.want {
			t.Fatalf("Factorial(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
