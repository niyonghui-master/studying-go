package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(getNumber())
	fmt.Println(getNumberByChan())
}

func getNumber() int {
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i = 5
		wg.Done()
	}()
	wg.Wait()
	return i
}

func getNumberByChan() int {
	var i int
	done := make(chan struct{})
	go func() {
		i = 5
		done <- struct{}{}
	}()
	<-done
	return i
}

//  首先，创建
type SafeNumber struct {
	val int
	m   sync.Mutex
}


