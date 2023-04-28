// wap in go language to create structure author. write a method show whose receiver in struct author

package main

import "fmt"

type author struct {
	id   int
	name string
}

func (a author) show() {
	fmt.Printf("I am Receiver %v \n", a.id)
	fmt.Printf("My name is %v\n", a.name)
}

func main() {
	a := author{
		1,
		"Om",
	}
	a.show()
	b := author{
		2,
		"Tom",
	}
	b.show()
}
