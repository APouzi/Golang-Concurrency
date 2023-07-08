package main

import (
	"fmt"
)

func CloseChannel() {
	ch1 := make(chan int)

	

	nums := []int{1, 2, 3}

	go insertChanV1(nums,ch1)

	ch2 := make(chan int)
	ch3 := make(chan int)
	go insertChanV2(ch2)
	go insertChanV2(ch3)

	for i := range ch1 {
		fmt.Println("CloseChannel in loop:",i)
	}
	
	
	for{
		select{
			case v, ok := <-ch2:
				if !ok{
					ch2 = nil 
				}else{
					fmt.Println("ch2",v)	
				}
					
			case v, ok:= <-ch3:
				if !ok{
					ch3 = nil
				}else{
					fmt.Println("ch3",v)
				}
					
		}
		if ch2 == nil && ch3 == nil{
			fmt.Println("out of loop")
			break
		}
	}
}

func insertChanV1(nums []int,ch chan int){
	for _, v := range nums {
		ch <- v
	}
	close(ch)
}

func insertChanV2(ch chan int){
	for i := 0; i < 10; i++{
		ch <- i 
	}
	close(ch)
}