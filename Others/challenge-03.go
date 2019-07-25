// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan() 
// 	arraySize, errInt := strconv.Atoi(scanner.Text())
// 	listOfElements, errRs := reader.ReadString('\n')
// 	if errInt != nil || errRs != nil {
// 		// raise an error
// 		os.Exit(1)
// 	}

// 	elements := convertStringListToInt(arraySize, listOfElements)

// 	result := sum(elements)
// 	fmt.Println(result)
// }

// func convertStringListToInt(size int, list string) []int64 {
// 	var output []int64
// 	var splittedList = strings.Fields(list)

// 	for idx, number := range splittedList {
// 		value, err := strconv.ParseInt(number, 10, 32)
// 		if err != nil {
// 			// raise an error
// 			os.Exit(1)
// 		}

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
