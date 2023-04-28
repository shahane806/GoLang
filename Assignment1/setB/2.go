// wap in go language to print PASCALS triangle
package main

import "fmt"

func main() {
	rows := 5
	temp := 1
	for i := 0; i <= rows; i++ {
		for j := 0; j <= rows-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			if i == 0 || k == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			fmt.Printf("%3d", temp)
		}
		fmt.Println()
	}

}
