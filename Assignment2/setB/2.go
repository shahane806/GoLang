// wap in go language to demonstrate use of names returns variables

package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter a and b integers")
	fmt.Scan(&a)
	fmt.Scan(&b)
	sum := add(a, b)
	fmt.Println(sum)
}
func add(a int, b int) (sum int) {
	sum = 0
	sum = a + b
	return sum
}
