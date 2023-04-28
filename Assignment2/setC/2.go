// wap in go language to create a file and write hello world in it and close the file by using defer statement

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createFile() {
	//for creating file we use os.create("filename") function
	//os.create("filename") it will return two things 1.file object 2. error log
	//we want to handle this both things

	file, err := os.Create("2.txt")
	if err != nil {
		log.Fatalf("Failed to creating file : %v", err)
	}
	/*
		for writing in file we can use <file>.Writing(<string>) method
		it will return two things 1. length and 2 . err log

	*/
	len, err := file.WriteString("Hello World!")
	if err != nil {
		log.Fatalf("Failed to write in file: %v", err)
	}
	fmt.Printf("File Name : %v \n", file.Name())

	fmt.Printf("%v is len of file\n", len)

	data, err := ioutil.ReadFile("2.txt")
	if err != nil {
		log.Fatalf("Failed to read file : %v", err)

	}
	for _, elementInData := range data {
		fmt.Printf("%c", elementInData)
	}

	defer file.Close()

}

func main() {
	createFile()
}
