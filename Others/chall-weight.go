package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func numberOfTrips(tripWeight int, packagesWeight []int) int {
	var output int
	var splittedList = strings.Fields(list)

	for _, number := range packagesWeight {
		if len(output) < size {
			value, _ := strconv.Atoi(number)
			output = append(output, value)
		}
	}

	return output
}