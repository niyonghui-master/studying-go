package main

import "fmt"

type Tester interface {
	getName() string
}

type Tester2 interface {
	printName() 
}

type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}

func (p Person) printName() {
	fmt.Println(p.name)
}

func main() {
	// 第一种情况
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("%T\n%s\n", f, f.getName())
	}

	// 第二种情况
	if t, ok2 := t.(Tester2); ok2 { // 重用变量名t （无需重新声明）
		check2(t) // 若类型断言为true，则新的t被转型
	}
}

func check2(t Tester2) {
	t.printName()
}