package main

import "fmt"

var (
	size     = 1024 // 原文 -> size := 1024 
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}

/*
	简短声明限制：
		1. 必须使用显示初始化
		2. 
*/ 
