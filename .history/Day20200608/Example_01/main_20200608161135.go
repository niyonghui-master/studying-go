package main

import "fmt"

func main() {
	// 浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a地址1: %p\n", a)
	b := a[0:3]
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)

	fmt.Printf("a地址2：%p\n", a)
	fmt.Printf("b地址 ：%p\n", b)

	b[0] = 100
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)

	c := make([]int, 3, 3)
	c = a[0:3]
	fmt.Println("c is:", c)
	fmt.Printf("c地址 : %p\n", c)

	// 深拷贝
	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1
	fmt.Printf("s1地址: %p, s2地址: %p\n", s1, s2)

	s1 = append(s1, 100)
	s2[0] = 200
	fmt.Printf("s1地址: %p, s2地址: %p\n", s1, s2)
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is:", s2)

	// copy函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	s3 := make([]int, 4, 4)
	s4 := make([]int, 5, 5)
	fmt.Println(copy(s3, s1))
	fmt.Println(copy(s4, s1))
}
