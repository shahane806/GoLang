// wap in go language to demonstrate working of slices (like append , remove, copy,etc)

package main

import (
	"fmt"
	"reflect"
)

func Remove(Slice []int, n int) []int {
	return append(Slice[:n], Slice[n+1:]...)
}

func main() {
	//Creating a slice by make function
	Slice1 := make([]int, 2, 10)
	Slice2 := make([]int, 2, 10)
	Slice1[0] = 2
	Slice1[1] = 3
	fmt.Printf("%v\n", Slice1)
	fmt.Printf("%v\n", reflect.ValueOf(Slice1).Kind())
	fmt.Printf("Length : %v\n", len(Slice1))
	fmt.Printf("Capacity : %v\n", cap(Slice1))

	Slice1 = append(Slice1, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("%v\n", Slice1)
	fmt.Printf("%v\n", reflect.ValueOf(Slice1).Kind())
	fmt.Printf("Length : %v\n", len(Slice1))
	fmt.Printf("Capacity : %v\n", cap(Slice1))

	Slice1 = Remove(Slice1, 5)
	Slice1 = Remove(Slice1, 9)
	fmt.Printf("%v\n", Slice1)
	fmt.Printf("%v\n", reflect.ValueOf(Slice1).Kind())
	fmt.Printf("Length : %v\n", len(Slice1))
	fmt.Printf("Capacity : %v\n", cap(Slice1))

	copy(Slice2, Slice1)
	fmt.Printf("%v\n", Slice2)
	fmt.Printf("%v\n", reflect.ValueOf(Slice2).Kind())
	fmt.Printf("Length : %v\n", len(Slice2))
	fmt.Printf("Capacity : %v\n", cap(Slice2))
}
