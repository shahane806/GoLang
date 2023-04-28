// wap in go language to accept two strings and compare them

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var s2 string

	fmt.Println("Enter two strings")
	fmt.Println("S1 :")

	fmt.Scan(&s1)
	fmt.Println("S2 :")
	fmt.Scan(&s2)
	if strings.Compare(s1, s2) == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
