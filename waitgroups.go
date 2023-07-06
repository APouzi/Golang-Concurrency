package main

import (
	"fmt"
	"sync"
)

func WaitGroupStarter() {
	var wg sync.WaitGroup
	added := 0
	go AddingPartOne(1, 1, &added, &wg)
	wg.Add(1)
	go AddingPartTwo(1, 1, &added, &wg)
	wg.Add(1)
	wg.Wait()
	fmt.Println("waitgroup.go result is:",added)
}

func AddingPartOne(a, b int, pass *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*pass += a + b
}

func AddingPartTwo(a, b int, pass *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*pass += a + b
}