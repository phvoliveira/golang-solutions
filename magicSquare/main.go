package main

import (
	"fmt"
)

func main() {
	square := make(map[int][]int, 3)

	for idx := range square {
		square[idx] = readInts(3)
	}

	rows, columns, diagonals := make(chan int), make(chan int), make(chan int)

	go checkRows(square, rows)
	go checkRows(square, columns)
	go checkRows(square, diagonals)

}

func checkRows(sqr map[int][]int, output chan int) {

}

func readInts(cap int) []int {
	var output []int
	for i := 0; i < cap; i++ {
		var num int
		fmt.Scanf("%d", &num)
		output = append(output, num)
	}

	fmt.Scanf("\n")
	return output
}
