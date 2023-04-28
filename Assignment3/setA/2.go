// wap in go language to accept the book details such as bookid,title,author,price.
// Read and display the details of n number of books

package main

import "fmt"

type book struct {
	bookId int
	title  string
	author string
	price  float32
}

func main() {
	var n int
	var books [10]book
	fmt.Println("How many books details you want to enter (max :10)")
	fmt.Scan(&n)
	fmt.Println("Enter book details")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter Book[%v] details \n", i+1)
		fmt.Printf("BookId : ")
		fmt.Scan(&books[i].bookId)
		fmt.Printf("Title : ")
		fmt.Scan(&books[i].title)
		fmt.Printf("author : ")
		fmt.Scan(&books[i].author)
		fmt.Printf("price : ")
		fmt.Scan(&books[i].price)

	}
	fmt.Println("Display Details of all books")
	fmt.Printf("%20s %20s %20s %20s\n", "BookId", "Title", "Author", "price")
	for i := 0; i < n; i++ {
		fmt.Printf("%20v %20v %20v %20v\n", books[i].bookId, books[i].title, books[i].author, books[i].price)

	}
}
