// wap in go language to create and print multidimensional slice

package main

import (
	"fmt"
	"reflect"
)

func main() {
	multiDimensionalSlice := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	for index, data := range multiDimensionalSlice {
		fmt.Printf("%v : %v \n", index, data)
	}
	fmt.Println(reflect.ValueOf(multiDimensionalSlice).Kind())
}
