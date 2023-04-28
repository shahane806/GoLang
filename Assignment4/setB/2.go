// wap in go language to demonstrate working type switch in interface

package main

import (
	"fmt"
)

func main() {
	var in interface{}
	in = "Hello"
	in = 3
	fmt.Printf("%T is type \n", in)

	switch typ := in.(type) {
	case int:
		fmt.Printf("It is %T\n", typ)
		break
	case string:
		fmt.Printf("It is %T\n", typ)
		break
	case float32:
		fmt.Printf("It is %T\n", typ)
		break
	default:
		fmt.Println("It is not in case")

	}
}
