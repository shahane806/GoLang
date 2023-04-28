// wap in go to print table of given number
package main

import "fmt"

func main() {
	var num int

	fmt.Println("Enter number you want a multiplication table of")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d \n", num, i, num*i)
	}
}
