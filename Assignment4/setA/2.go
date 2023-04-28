// wap in go language to print multiplication of two numbers using method

package main

import "fmt"

func mul(num1 int, num2 int) int {
	return num1 * num2
}
func main() {
	var multiplication int
	var num1 int
	var num2 int
	fmt.Printf("Enter two numbers \n")
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	multiplication = mul(num1, num2)
	fmt.Printf("The multiplication of two numbers using method is : %v", multiplication)
}
