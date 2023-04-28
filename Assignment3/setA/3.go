// wap in go language to initialize a slice using multi line syntax and display

package main

import (
	"fmt"
	"reflect"
)

func main() {
	multiLineSlice := []int{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 0,
	}
	for index, data := range multiLineSlice {
		fmt.Printf("%v : %v\n", index, data)
	}
	fmt.Println(reflect.ValueOf(multiLineSlice).Kind())
}
