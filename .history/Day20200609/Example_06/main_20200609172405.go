package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func ()  {
			fmt.P
		}
	}
}