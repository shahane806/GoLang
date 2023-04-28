// wap in go language to print addition of two numbers using function
package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("Enter num1")
	fmt.Scan(&num1)
	fmt.Println("Enter num2")
	fmt.Scan(&num2)
	add := additionOfTwoNumbers(num1, num2)
	fmt.Printf("%v + %v = %v\n", num1, num2, add)
}

func additionOfTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}
