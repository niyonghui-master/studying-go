package main

func getNumber() {
	var i int
	go func() {
		i = 5
	}()
	return i
}