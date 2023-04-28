// wap in go language to illustrate the concept of returning multiple values from a function

package main

import "fmt"

func operators(a int, b int) (int, int, int, int) {
	return a + b, a - b, a / b, a * b
}
func main() {
	var add, sub, div, mul int
	add, sub, div, mul = operators(3, 2)

	fmt.Printf("a = %v and b = %v\n", 3, 2)
	fmt.Printf("%v is addition\n", add)
	fmt.Printf("%v is sub\n", sub)
	fmt.Printf("%v is division\n", div)
	fmt.Printf("%v is mul\n", mul)
}
