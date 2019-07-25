package main

import (
	"fmt"
)

type myInt int

func (i myInt) init() {
	fmt.Println("Init!")
	i = 1
}

func main() {
	var inter myInt
	fmt.Println(inter)
}
