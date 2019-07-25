package main

import (
	"fmt"
)

func main() {
	var count, n, d, m int
	fmt.Scanf("%d\n", &n)
	chocolateBar := readLine(n)
	fmt.Scanf("%d %d\n", &d, &m)

	for idx := range chocolateBar {
		sum := 0
		if idx+m <= len(chocolateBar) {
			for _, num := range chocolateBar[idx : idx+m] {
				sum += num
			}

			if sum == d {
				count++
			}
		}
	}

	fmt.Printf("%d\n", count)
}

func readLine(cap int) []int {
	output := make([]int, 0, cap)
	for i := 0; i < cap; i++ {
		var num int
		fmt.Scanf("%d", &num)
		output = append(output, num)
	}

	fmt.Scanf("\n")
	return output
}
