package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var n, m int
	var numbers []int
	fmt.Scanf("%d %d\n", &n, &m)
	a, b := readLine(n), readLine(m)
	init, end := a[len(a)-1], b[0]

	for f := init; f <= end; f++ {
		wg.Add(1)
		go func(number int, nbrs *[]int) {
			if factorBy(a, number, false) && factorBy(b, number, true) {
				*nbrs = append(*nbrs, number)
			}
			wg.Done()
		}(f, &numbers)
	}
	wg.Wait()

	fmt.Printf("%d\n", len(numbers))
}

func readLine(cap int) []int {
	output := make([]int, 0, cap)
	for i := 0; i < cap; i++ {
		var num int
		fmt.Scanf("%d", &num)
		output = append(output, num)
	}

	fmt.Scanf("\n")
	sort.Ints(output)
	return output
}

func factorBy(numbers []int, number int, reverse bool) bool {
	output := true

	for _, num := range numbers {
		if (reverse && num%number != 0) || (!reverse && number%num != 0) {
			output = false
		}
	}

	return output
}
