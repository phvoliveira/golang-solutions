package main

import (
	"fmt"
	"sort"
)

func main() {
	var numOfLoops int
	fmt.Scanf("%d\n", &numOfLoops)

	for l := 0; l < numOfLoops; l++ {
		getNumberOfBribes()
	}
}

func getNumberOfBribes() {
	var maxNumPerson, result int
	fmt.Scanf("%d\n", &maxNumPerson)
	list, orderedList := getList(maxNumPerson)
	cList, i, j, c := false, 0, 0, 0

	for !cList {
		c++
		if list[i] > 0 {
			if list[i] == orderedList[i] {
				j++
			} else if list[i]-(i+1) > 2 {
				fmt.Println("Too chaotic")
				return
			} else if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				result++
				i = j - 1
			}
		}
		i++

		if i == (len(list) - 1) {
			cList = true
		}
	}

	fmt.Printf("%v\n", result)
}

func getList(maxNumPerson int) (list, orderedList []int) {
	for i := 0; i < maxNumPerson; i++ {
		var personNumber int
		fmt.Scanf("%d", &personNumber)
		if personNumber > 0 {
			list = append(list, personNumber)
			orderedList = append(orderedList, personNumber)
		}
	}
	sort.Ints(orderedList)
	return
}
