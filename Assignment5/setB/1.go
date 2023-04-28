// wap in go prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one using delay function
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		time.Sleep(2500000)
		fmt.Printf("%v\n", i)
	}
}
