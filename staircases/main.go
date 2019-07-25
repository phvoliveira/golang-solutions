package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type floor string
type stairs []floor

func (f floor) String() string {
	switch {
	case f == "":
		return " "
	default:
		return string(f)
	}
}

func (s stairs) String() (r string) {
	for _, f := range s {
		r += fmt.Sprintf("%s", f)
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringText := scanner.Text()
	stringNumbers := strings.Split(stringText, " ")
	number, _ := strconv.Atoi(stringNumbers[0])
	stairs := make(stairs, number)

	for i := (number - 1); i >= 0; i-- {
		stairs[i] = "#"
		fmt.Printf("%s\n", stairs)
	}
}
