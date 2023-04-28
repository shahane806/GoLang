// wap in go language to print fibonaci series of n terms

package main

import "fmt"

func main() {
	var n int
	var first int
	var second int
	var third int

	fmt.Println("Enter number to calculate fibonacci series of that number")
	fmt.Scanf("%d", &n)
	first = 0
	second = 1

	fmt.Printf("%d %d", first, second)
	for i := 2; i < n; i++ {
		third = first + second
		first = second
		second = third
		fmt.Printf(" %d ", third)
	}

}
