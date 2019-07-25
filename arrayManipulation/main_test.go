package main

import "testing"

func Test01(t *testing.T) {
	arrayManipulation(5, [][]int32{
		[]int32{1, 2, 100},
		[]int32{2, 5, 100},
		[]int32{3, 4, 100},
	})
}

func Test02(t *testing.T) {
	arrayManipulation(10, [][]int32{
		[]int32{2, 6, 8},
		[]int32{3, 5, 7},
		[]int32{1, 8, 1},
		[]int32{5, 9, 15},
	})
}
