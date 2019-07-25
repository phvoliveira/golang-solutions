package main

import (
	"fmt"
)

func main() {
	var s, t, a, b, m, n int
	fmt.Scanf("%d %d\n", &s, &t)
	fmt.Scanf("%d %d\n", &a, &b)
	fmt.Scanf("%d %d\n", &m, &n)

	scoreFromAPoint := countScore(s, t, a, m)
	scoreFromBPoint := countScore(s, t, b, n)

	fmt.Printf("%d\n", scoreFromAPoint)
	fmt.Printf("%d\n", scoreFromBPoint)
}

func countScore(s, t, startPoint, numberOfThrows int) (c int) {
	var d int
	for i := 0; i < numberOfThrows; i++ {
		fmt.Scanf("%d", &d)
		f := startPoint + d

		if f >= s && f <= t {
			c++
		}
	}
	return
}
