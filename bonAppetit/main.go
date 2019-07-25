package main

import (
	"fmt"
)

func main() {
	var bCharged, bActual, n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	numbers := readInts(n)
	fmt.Scanf("%d\n", &bCharged)

	for idx, num := range numbers {
		if idx != k {
			bActual += num
		}
	}

	bActual = bActual / 2
	overcharge := bCharged - bActual

	if overcharge == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Printf("%d\n", overcharge)
	}
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
