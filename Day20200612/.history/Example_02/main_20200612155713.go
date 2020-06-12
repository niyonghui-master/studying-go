package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, " -> ", v, *v)
	}
}

// for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为 3，所以输出都是3


// make() 与 new() 的区别
/*
	new(T) 和 make(T, args) 是 Go 语言內建函数，用来分配内存，但使用的类型不同。
	
	new(T) 会为 T 类型的新值分配已置零的内存空间，并返回
*/ 