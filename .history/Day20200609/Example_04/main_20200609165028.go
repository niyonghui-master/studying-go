package main

import (
	"fmt"
	"time"
)

func main() {
	chananel := make(chan bool)

	go func() {
		fmt.Println("2222")
		chananel <- true
	}()

	<- chananel 

	fmt.Println("1111")

	time.Sleep(2 * time.Second)
}