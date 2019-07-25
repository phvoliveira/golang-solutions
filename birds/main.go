package main

import (
	"fmt"
)

func main() {
	var highestNumber, highestCount, n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	numbers := readInts(n)
	birds := make(map[int]int, n)

	for _, num := range numbers {
		birds[num]++
	}

	for id, count := range birds {
		if highestCount < count {
			highestCount = count
			highestNumber = id
		}
	}

	fmt.Printf("%d\n", highestNumber)
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
