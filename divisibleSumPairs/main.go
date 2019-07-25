package main

import (
	"fmt"
)

func main() {
	var count, n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	numbers := readInts(n)

	for i, j := 0, 1; i < len(numbers)-1; j++ {
		if (numbers[i]+numbers[j])%k == 0 {
			count++
		}

		if j == len(numbers)-1 {
			i++
			j = i
		}
	}

	fmt.Printf("%d\n", count)
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
