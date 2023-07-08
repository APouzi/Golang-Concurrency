package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go addPartOne(1, 2, ch1)
	go addPartTwo(1, 2, ch2)
	fmt.Println(<-ch1 + <-ch2)

	// Examples being called
	CallThese()
	WaitGroupStarter()
	Channelinit()
	CloseChannel()
}

func addPartOne(a, b int, c chan int) {
	c <- a + b
}

func addPartTwo(a, b int, c chan int) {
	c <- a + b
}