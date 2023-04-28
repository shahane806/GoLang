// wap in go language to show the compiler throws an error if a variable is declared but not used

package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println(a)
}
