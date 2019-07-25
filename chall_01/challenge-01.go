package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arraySize int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &arraySize)
	listOfElements, _ := reader.ReadString('\n')

	elements := addElements(arraySize, listOfElements)

	result := sum(elements)
	fmt.Println(result)
}

func addElements(size int, list string) []int {
	var output []int
	var splittedList = strings.Fields(list)

	for _, number := range splittedList {
		if len(output) < size {
			value, _ := strconv.Atoi(number)
			output = append(output, value)
		}
	}

	return output
}

func sum(numbers []int) int {
	output := 0
	for _, number := range numbers {
		output += number
	}
	return output
}
