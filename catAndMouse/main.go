package main

import (
	"fmt"
)

func main() {
	var q, i int
	distanceA, distanceB := make(chan int), make(chan int)
	fmt.Scanf("%d", &q)

	for ; i < q; i++ {
		positions := readInts(3)
		catA, catB, mouseC := positions[0], positions[1], positions[2]
		go distance(distanceA, catA, mouseC)
		go distance(distanceB, catB, mouseC)

		valueA := <-distanceA
		valueB := <-distanceB
		if valueA == valueB {
			fmt.Println("Mouse C")
		} else if valueA < valueB {
			fmt.Println("Cat A")
		} else {
			fmt.Println("Cat B")
		}
	}
}

func distance(output chan int, a, b int) {
	distance := (a - b)
	if distance < 0 {
		distance *= -1
	}
	output <- distance
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
