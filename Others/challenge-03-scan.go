// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
//  	var arraySize int
//  	reader := bufio.NewReader(os.Stdin)
//  	fmt.Scanf("%d", &arraySize)
//  	listOfElements, _ := reader.ReadString('\n')

// 	elements := convertStringListToInt(arraySize, listOfElements)

// 	result := sum(elements)
// 	fmt.Println(result)
// }

// func convertStringListToInt(size int, list string) []int64 {
// 	var output []int64
// 	var splittedList = strings.Fields(list)

// 	for idx, number := range splittedList {
// 		value, _ := strconv.ParseInt(number, 10, 32)

// 		if (idx < size) {
// 			output = append(output, value)
// 		}
// 	}

// 	return output
// }

// func sum(numbers []int64) int64 {
// 	var output int64
// 	for _, number := range numbers {
// 		output += number
// 	}
// 	return output
// }
