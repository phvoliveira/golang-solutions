package main

import (
	"fmt"
)

func main() {
	var (
		c                 string
		n, level, valleys int
	)

	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &c)

	for _, r := range []rune(c) {
		switch {
		case r == 'U':
			level++
		case r == 'D':
			level--
		}

		if level == 0 && r == 'U' {
			valleys++
		}
	}

	fmt.Printf("%d\n", valleys)
}
