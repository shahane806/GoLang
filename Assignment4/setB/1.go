// wap in go language to create structure student . Write a method show() whose receiver is a pointer of struct student

package main

import "fmt"

type student struct {
	id   int
	name string
}
type studinfo interface {
	show()
}

func (s *student) show() {
	fmt.Printf("I am Receiver %v \n", s.id)
	fmt.Printf("My name is %v\n", s.name)
}
func main() {
	s1 := student{
		id:   1,
		name: "Om",
	}
	s2 := student{
		id:   2,
		name: "Tom",
	}
	s3 := student{
		id:   3,
		name: "com",
	}
	var i studinfo
	i = &s1
	i.show()
	i = &s2
	i.show()
	i = &s3
	i.show()

}
