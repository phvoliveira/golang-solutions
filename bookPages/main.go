package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &p)

	front, back := p/2, ((n / 2) - (p / 2))

	if front < back {
		fmt.Printf("%d\n", front)
	} else {
		fmt.Printf("%d\n", back)
	}
}
