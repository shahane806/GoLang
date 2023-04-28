// wap in go language to create an interface and display it's values with the help of type assertion
package main

import "fmt"

func main() {
	var i interface{} = "This is String"
	n, t := i.(string)
	fmt.Printf("Assertion is successful %v %T\n", t, n)
}
