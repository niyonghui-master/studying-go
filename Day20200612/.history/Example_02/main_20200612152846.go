package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]int)

	for key, val := range slice {
		fmt.Printf("%p\n", val)
		m[key] = val
	}

	for k, v := range m {
		fmt.Println(k, " -> ", v, v)
	}
}
