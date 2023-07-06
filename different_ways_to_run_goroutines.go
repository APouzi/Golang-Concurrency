package main

import (
	"fmt"
	"time"
)

type versionThree struct{}

func versionOne(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "This is a goroutine that is being run by being called outside of the CallThese as a function"
}

func (m *versionThree) versionThreeMethodCall(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "This is a goroutine that is being run by being called outside of the CallThese as a method"
}

func CallThese() {
	
	chV1, chV2, chV3 := make(chan string), make(chan string), make(chan string)
	
	go versionOne(chV1)


	go func() {
		time.Sleep(2 * time.Second)
		chV2 <- "This is a goroutine that is being run by being called inside of CallThese"
	}()

	methodCall := versionThree{}
	go methodCall.versionThreeMethodCall(chV3)
	fmt.Println(<-chV1)
	fmt.Println(<-chV2)
	fmt.Println(<-chV3)

}

