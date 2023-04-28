// WAP in go language to print whether number is even or odd
package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter Number to check it is even or odd\n")
	fmt.Scanf("%d", &num)
	if num%2 == 0 {
		fmt.Println("The Number is EVEN")
	} else {

		fmt.Println("The Number is ODD")
	}
}
