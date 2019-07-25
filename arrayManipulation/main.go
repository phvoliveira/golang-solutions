package main

import (
	"fmt"
	"sort"
)

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int, n*2)

	for _, q := range queries {
		arr[q[0]] += int(q[2])
		if int(q[1]) < len(arr) {
			arr[q[1]+1] -= int(q[2])
		}
	}

	for j := 1; j < len(arr); j++ {
		arr[j] += arr[j-1]
	}

	sort.Ints(arr)
	return int64(arr[len(arr)-1])
}

func main() {
	fmt.Printf("Primeiro Resultado: %v\n", arrayManipulation(5, [][]int32{
		[]int32{1, 2, 100},
		[]int32{2, 5, 100},
		[]int32{3, 4, 100},
	}))

	fmt.Printf("Segundo Resultado: %v\n", arrayManipulation(10, [][]int32{
		[]int32{2, 6, 8},
		[]int32{3, 5, 7},
		[]int32{1, 8, 1},
		[]int32{5, 9, 15},
	}))

	fmt.Printf("Terceiro Resultado: %v\n", arrayManipulation(4, [][]int32{
		[]int32{2, 3, 603},
		[]int32{1, 1, 286},
		[]int32{4, 4, 882},
		[]int32{3, 4, 2},
		[]int32{3, 3, 1000},
	}))
}
