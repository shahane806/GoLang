// wap in go language to sort array elements in ascending order

package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Printf("Enter array elements\n")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Sorting Array elements in ascending order\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] + arr[j+1]
				arr[j+1] = arr[j] - arr[j+1]
				arr[j] = arr[j] - arr[j+1]
			}
		}
	}
	fmt.Printf("Sorted Process complete.... here the Array now : ")
	for _, data := range arr {
		fmt.Printf("%v ", data)
	}

}
