package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	timeImput := scanner.Text()
	t, _ := time.Parse("15:04:05PM", timeImput)
	fmt.Printf("%v\n", t.Format("15:04:05"))
}
