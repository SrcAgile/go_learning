package snip

import "testing"

func TestSnip(t *testing.T) {
	//s []int := {1;2;3;4}
	// var a1 int = 6

	var a [2]int
	// a2 := [5]int{1, 2, 3, 4}
	s := [4]int{1, 2, 3, 4}
	t.Log(len(s), cap(s))
	t.Log(a)
}
