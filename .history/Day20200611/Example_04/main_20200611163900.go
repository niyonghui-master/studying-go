package main

import "fmt"

type Test interface {}

type Test1 interface {
	TestFunc()
}

type Structure struct {
	a int
}

func (s *Structure) TestFunc() {
	fmt.Println("OK, Let's rock and roll!")
}

func fTest(t Test) {
	fmt.Printf("Test %T", t)
	fmt.Println(t == nil) 
}

func fTest1(t1 Test1) {
	fmt.Printf("Test1 %T", t1)
	fmt.Println(t1 == nil)
}

func fStructure(s *Structure) {
	fmt.Printf("Structure %T", s)
	fmt.Println(s == nil)
}

func main() {
	var s *Structure = nil
	fTest(s) // false
	fTest1(s) // false
	fStructure(s) // true
	s.TestFunc() // OK, Let's rock and roll!
}