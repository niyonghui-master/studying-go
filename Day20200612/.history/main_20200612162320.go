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

	if s1 == s2 {
		fmt.Println()
	}
}
