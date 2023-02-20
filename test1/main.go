package main

import "fmt"

type Example struct {
	Name string
}

func (e *Example) Original(name string) {
	e.Name = name
}

func main() {
	e := &Example{"hello"}
	fmt.Println(e)
	e.Original("goodbye")
	fmt.Println(e)
	fmt.Println("hello world")
}
