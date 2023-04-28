// wap in go language to concatenate two strings using pointers

package main

import (
	"fmt"
)

func main() {
	var s1 string
	var s2 string

	var s *string
	var b *string

	s = &s1
	b = &s2

	fmt.Println("Enter two strings")
	fmt.Println("S1 :")

	fmt.Scan(&s1)
	fmt.Println("S2 :")
	fmt.Scan(&s2)

	for i := 0; i < 1; i++ {
		*s += *b
	}
	fmt.Println(s1)
}
