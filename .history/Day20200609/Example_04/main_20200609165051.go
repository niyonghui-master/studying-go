package main

import (
	"fmt"
	"time"
)

func main() {
	chananel := make(chan bool, 1)

	go func() {
		fmt.Println("2222")
		chananel <- true
	}()

	<-chananel

	fmt.Println("1111")

	
}
