package main

import (
	"fmt"
)

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d\n", &x1, &v1, &x2, &v2)

	if x1 > x2 {
		x1, v1, x2, x2 = x2, x2, x1, v1
	}

	distanceA := x1 + v1
	distanceB := x2 + v2

	if (v1 < v2) || (x1 != x2 && v1 == v2) {
		fmt.Println("NO")
	} else {
		for {
			if distanceA == distanceB {
				fmt.Println("YES")
				break
			}

			newDistanceA := distanceA + v1

			if distanceA > distanceB || ((distanceB - newDistanceA) > (distanceB - distanceA)) {
				fmt.Println("NO")
				break
			}

			distanceA = newDistanceA
			distanceB += v2
		}
	}
}
