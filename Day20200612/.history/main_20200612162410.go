package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "NYH"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "NYH"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int 
		
	}
}
