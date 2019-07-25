package main

import (
	"fmt"
)

func main() {
	x, c, done := 20, make(chan int), make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	for j := 0; j < x; j++ {
		go read(c, done)
	}

	for k := 0; k < x; k++ {
		<-done
	}
}

func read(c chan int, done chan bool) {
	for n := range c {
		fmt.Println(n)
	}
	done <- true
}
