// wap in go language to swap two numbers using call by reerence concept

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("Enter two number for swap")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Printf("%v and %v Before swapping \n", num1, num2)
	swap(&num1, &num2)
	fmt.Printf("%v and %v After swapping \n", num1, num2)
}
func swap(num1 *int, num2 *int) {
	*num1 = *num1 + *num2
	*num2 = *num1 - *num2
	*num1 = *num1 - *num2
}
