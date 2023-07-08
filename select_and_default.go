package main

import "fmt"

func StartSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go NumtoChannel(ch1)
	go NumtoChannel(ch2)
	go NumtoChannel(ch3)
	go NumtoChannel(ch4)

	chansSlice := []chan int{ch1, ch2, ch3, ch4}
	SelectReader(chansSlice)
	// SelectReaderV2(chansSlice)
}

func NumtoChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func SelectReader(chans []chan int) {
	nums := []int{}
	for i := 0; i < 40; i++ {
		select {
		case v := <-chans[0]:
			nums = append(nums, v)
		case v := <-chans[1]:
			nums = append(nums, v)
		case v := <-chans[2]:
			nums = append(nums, v)
		case v := <-chans[3]:
			nums = append(nums, v)
		}
	}
	fmt.Println(nums)
}

