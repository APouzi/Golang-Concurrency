package main

import "fmt"

func StartSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	chansSlice := []chan int{ch1, ch2, ch3, ch4}
	
	go NumtoChannel(ch1)
	go NumtoChannel(ch2)
	go NumtoChannel(ch3)
	go NumtoChannel(ch4)
	
	// SelectReader(chansSlice)
	SelectReaderV2(chansSlice)

	
}

func NumtoChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
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

func SelectReaderV2(chans []chan int) {
	nums := []int{}
	for {
		select {
			case v, ok := <-chans[0]:
				if ok != false{
					fmt.Println(v)
					nums = append(nums, v)
					
				}else{
					chans[0] = nil
					// nums = append(nums, v)
				}
			case v, ok := <-chans[1]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[1] = nil
					// nums = append(nums, v)
				}
			case v, ok := <-chans[2]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[2] = nil
					// nums = append(nums, v)
				}
			case v, ok := <-chans[3]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[3] = nil
					// nums = append(nums, v)
				}
		}
		if chans[0] == nil && chans[1] == nil && chans[2] == nil && chans[3] == nil{
			break
		}
	}
	fmt.Println(nums)
}

