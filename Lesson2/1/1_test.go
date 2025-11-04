package main

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		in   []int
		want map[int]int
	}{
		{[]int{1, 2, 3}, map[int]int{1: 1, 2: 1, 3: 1}},
		{[]int{1, 2, 3, 3}, map[int]int{1: 1, 2: 1, 3: 2}},
		{[]int{1, 5, 4, 4, 5}, map[int]int{1: 1, 4: 2, 5: 2}},
		{[]int{1, 2, 3, 4, 5, 1}, map[int]int{1: 2, 2: 1, 3: 1, 4: 1, 5: 1}},
	}
	for _, c := range cases {
		got := Count(c.in)
		if reflect.DeepEqual(got, c.want) != true {
			t.Errorf("Count(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
