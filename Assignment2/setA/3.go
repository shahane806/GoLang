// wap in go language using function to check wheterh accepts number is palindrome or not

package main

import "fmt"

func main() {
	num := 0
	fmt.Println("Enter number to check it is palindrome or not")
	fmt.Scan(&num)

	palindrome(num);
	

}
func palindrome(num int){
	temp := num
	rem := 0
	rev := 0

	for i := 0; num != 0; i++ {
		rem = num % 10
		rev = rev*10 + rem
		num = num / 10
	}
	if temp == rev {
		fmt.Println("The number is palindrome")

	} else {
		fmt.Println("The number is not palindrome")
	}
}