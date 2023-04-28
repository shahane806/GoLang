// wap in go language to accept user choice and print answer of using arithmetical operators

package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Printf("Enter a and b\n")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v / %v = %v\n", a, b, a/b)
	fmt.Printf("%v %% %v = %v\n", a, b, a%b)

}
