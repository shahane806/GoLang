// wap in go language to print recursive sum of digits of given number

package main

import "fmt"

func main() {
	var num1 int
	sum := 0
	fmt.Println("Enter number")
	fmt.Scan(&num1)
	sum = recursiveSum(num1)
	fmt.Printf("%v is sum of all digits\n", sum)
}
func recursiveSum(num1 int) int {
	sum := 0
	if num1 == 0{
		return 0
	}
	sum = num1%10+recursiveSum(num1/10)
	return sum
}