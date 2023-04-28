// wap in go language to illustrate the concept of call by value

package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enters a and b integers")
	fmt.Scan(&a, &b)
	fmt.Printf("%v  and %v ", a, b)
	swap(a, b)
}
func swap(a int, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("%v  and %v ", a, b)
}

//this is call by value the actual value will not swapping
// when we use call by reference the actual value will change
