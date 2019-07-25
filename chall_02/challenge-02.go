package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Reading input values
	aliceBoard, _ := reader.ReadString('\n')
	bobBoard, _ := reader.ReadString('\n')

	scoreboards := map[string][]int{
		"Alice": make([]int, len(aliceBoard)),
		"Bob":   make([]int, len(bobBoard)),
	}

	// transform into array of int values
	scoreboards["Alice"] = convertStringListToInt(aliceBoard)
	scoreboards["Bob"] = convertStringListToInt(bobBoard)

	// comparing values and return scoreboard
	result := scoreboardChallenge(scoreboards)

	fmt.Println(result)
}

func convertStringListToInt(list string) []int {
	var output []int
	var splittedList = strings.Fields(list)

	for _, number := range splittedList {
		value, err := strconv.ParseInt(number, 10, 32)
		if err == nil {
			// raise an error
			os.Exit(1)
		}
		output = append(output, int(value))
	}

	return output
}

func scoreboardChallenge(scoreboards map[string][]int) string {
	totalLen := len(scoreboards["Alice"])

	if totalLen > len(scoreboards["Bob"]) {
		totalLen = len(scoreboards["Bob"])
	}

	challengers := map[string]int{
		"Alice": 0,
		"Bob":   0,
	}

	for idx := 0; idx < totalLen; idx++ {
		if scoreboards["Alice"][idx] > scoreboards["Bob"][idx] {
			challengers["Alice"]++
		} else if scoreboards["Bob"][idx] > scoreboards["Alice"][idx] {
			challengers["Bob"]++
		}
	}

	aliceResult := strconv.Itoa(challengers["Alice"])
	bobResult := strconv.Itoa(challengers["Bob"])

	return aliceResult + " " + bobResult
}
