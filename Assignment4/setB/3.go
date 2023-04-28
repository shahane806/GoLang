// wap in go language to copy all elements of one array into another array using method

package main

import "fmt"

func main() {
	var arr1 [2]int
	var arr2 [2]int
	fmt.Println("Enter 2 elements in the array")
	for i := 0; i < 2; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Println("Enter 5 elements in array2")
	for i := 0; i < 2; i++ {
		fmt.Scan(&arr2[i])
	}
	fmt.Printf("Array 1 : %v\n", arr1)
	fmt.Printf("Array 2 : %v\n", arr2)
	fmt.Printf("Copying content of Array2 in Array1\n")
	copy(arr1[:], arr2[:])

	fmt.Printf("Array 1 : %v\n", arr1)

}
