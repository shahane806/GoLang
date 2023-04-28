/*
Wap in go language to swap the number without temporary variable
*/
package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Enter two numbers ")
	fmt.Printf("Enter a : ")
	fmt.Scan(&a)
	fmt.Printf("Enter b : ")
	fmt.Scan(&b)
	fmt.Printf("Before Swapping \n")
	fmt.Println("A  : ", a)
	fmt.Println("B  : ", b)
	a = a + b // a = 10 b = 20 a = 30
	b = a - b // b = 30 - 20 b = 10
	a = a - b // a = 30 - 10 a = 20
	fmt.Printf("After Swapping \n")
	fmt.Println("A  : ", a)
	fmt.Println("B  : ", b)

}
