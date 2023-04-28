// wap in go language to find the largest and smallest number in an array

package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	fmt.Println("Enter 10 elements in the array : ")
	for i := 0; i < 10; i++ {
		fmt.Printf("arr[%d] : ", i+1)
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
	//bubble sort for array
	for i := 0; i < 10; i++ {
		for j := 0; j < 10-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] + arr[j+1]
				arr[j+1] = arr[j] - arr[j+1]
				arr[j] = arr[j] - arr[j+1]
			}
		}
	}
	
	fmt.Println(arr)
	fmt.Printf("The maximum element of the array is : %v\n", arr[len(arr)-1])
	fmt.Printf("The minimum element of the array is : %v\n", arr[0])

}


