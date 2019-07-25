package main

import (
	"fmt"
)

func main() {
	var s, n, m, i, j, biggest, lowest int

	fmt.Scanf("%d %d %d\n", &s, &n, &m)
	highestSum, pricesA, pricesB := -1, readInts(n), readInts(m)
	biggest, lowest = n, m
	if m > n {
		biggest, lowest = m, n
		pricesA, pricesB = pricesB, pricesA
	}

	for ; i < lowest; j++ {
		sum := (pricesA[j] + pricesB[i])

		if sum > highestSum && sum <= s {
			highestSum = sum
			if sum == s {
				break
			}
		}

		if j == (biggest - 1) {
			i++
			j = 0
		}
	}

	fmt.Printf("%d\n", highestSum)
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
