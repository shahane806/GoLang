// wap in go to illustrate how to create an anonyous goroutine

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go func() {
		fmt.Println("Generate")
		wg.Done()
	}()
	go func() {
		time.Sleep(2000000000)
		fmt.Println("World")
		wg.Done()
	}()
	go func() {
		time.Sleep(3000000000)
		fmt.Println("Next")
		wg.Done()
	}()

	go func() {
		time.Sleep(1000000000)
		fmt.Println("Hello")
		wg.Done()
	}()
	wg.Wait()
}
