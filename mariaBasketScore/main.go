package main

import (
	"fmt"
)

func main() {
	var numberOfGames, highestNumber, highestCount, lowestNumber, lowestCount int
	fmt.Scanf("%d\n", &numberOfGames)
	score := readLine(numberOfGames)
	highestNumber = score[0]
	lowestNumber = score[0]

	for f := 1; f <= len(score)-1; f++ {
		if highestNumber < score[f] {
			highestCount++
			highestNumber = score[f]
		}

		if lowestNumber > score[f] {
			lowestCount++
			lowestNumber = score[f]
		}
	}

	fmt.Printf("%d %d\n", highestCount, lowestCount)
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
