package main

import (
	"fmt"
	"sync"
)

func Channelinit(){
	ChannelBlocked()
	ChannelRead()
	ChannelReadV2()
}

func ChannelRead(){
	fmt.Println("\n This is ChannelRead() run of code \n ")
	ch1 := make(chan int)
	// ch2 := make(chan int)

	go SumTwo(1,2, ch1)
	go SumTwoPartTwo(2,3, ch1)
	fmt.Println("ChannelRead",<-ch1)
	fmt.Println("ChannelRead",<-ch1)
	fmt.Println("\n This is ChannelRead() end of code \n ")
}

func ChannelReadV2(){
	fmt.Println("\n This is ChannelReadV2() run of code \n ")
	ch1 := make(chan int)
	// ch2 := make(chan int)

	go SumTwo(1,2, ch1)
	go SumTwoPartTwo(2,3, ch1)
	fmt.Println("ChannelReadV2",<-ch1 + <-ch1)
	fmt.Println("\n This is ChannelReadV2() end of code \n ")
}

func SumTwo(a,b int, ch chan int){
	ch <- a +b
}
func SumTwoPartTwo(a,b int, ch chan int){
	ch <- a +b
}


func ChannelBlocked() {
	var wg *sync.WaitGroup
	fmt.Println("\n This is channel block run of code \n ")
	ch1 := make(chan int)
	// What we are doing here is we are sending off a goroutine with a loop that is calling to pull from a channel. 
	// This is a blocking event, meaning that the only way this goroutine can be continued is by simply inserting into the channel 10 times.
	go func(){
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch1)
		}
	}()
	ChannelReader(ch1, wg)
	fmt.Println("\n End of channel block \n ")
}

func ChannelReader(ch chan int, wg *sync.WaitGroup){
	 for i := 0; i < 10; i++ {
		ch <- 1
	}
}