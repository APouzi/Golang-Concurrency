package main

import "fmt"

func SelectInit(){
	fmt.Println("\n Select and Case Start:")
	StartSelectWithoutClose()
	StartSelectWithClose()
}

func StartSelectWithoutClose() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	chansSlice := []chan int{ch1, ch2, ch3, ch4}
	
	go NumtoChannel(ch1)
	go NumtoChannel(ch2)
	go NumtoChannel(ch3)
	go NumtoChannel(ch4)
	
	fmt.Println("\n Select with numbered loop and no close:")
	SelectReader(chansSlice)
}


func NumtoChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func StartSelectWithClose() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	chansSlice := []chan int{ch1, ch2, ch3, ch4}
	
	go NumtoChannelWithClose(ch1)
	go NumtoChannelWithClose(ch2)
	go NumtoChannelWithClose(ch3)
	go NumtoChannelWithClose(ch4)
	
	fmt.Println("\n Select with inifite loop and close, making it dynamic:")
	SelectReaderV2(chansSlice)
	
}


func NumtoChannelWithClose(ch chan int) {
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
					nums = append(nums, v)
				}else{
					chans[0] = nil
				}
			case v, ok := <-chans[1]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[1] = nil
				}
			case v, ok := <-chans[2]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[2] = nil
				}
			case v, ok := <-chans[3]:
				if ok != false{
					nums = append(nums, v)
				}else{
					chans[3] = nil
				}
		}
		if chans[0] == nil && chans[1] == nil && chans[2] == nil && chans[3] == nil{
			break
		}
	}
	fmt.Println(nums)
}

