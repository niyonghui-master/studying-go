package main

import (
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	ua := &UserAges{
		ages: make(map[string]int),
	}

	go func() {
		ua.Add("A", 21)
	}()

	fmt.ua.Get("A")

	time.Sleep(5 * time.Second)
}
