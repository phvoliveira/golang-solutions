package main

import (
	"fmt"
)

func main() {
	var pairs, n int
	fmt.Scanf("%d %d\n", &n)
	numbers := readInts(n)
	socks := make(map[int]int)

	for _, num := range numbers {
		if _, ok := socks[num]; ok {
			pairs++
			delete(socks, num)
		} else {
			socks[num] = 1
		}
	}

	fmt.Printf("%d\n", pairs)
}

func readInts(cap int) []int {
	output := make([]int, 0, cap)
	for i := 0; i < cap; i++ {
		var num int
		fmt.Scanf("%d", &num)
		output = append(output, num)
	}

	fmt.Scanf("\n")
	return output
}
