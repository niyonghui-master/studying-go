package main

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *s)