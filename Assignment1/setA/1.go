/*WAP in go language to print Student name, rollno, division and college name*/

package main

import "fmt"

func main() {
	var name string
	var rollno int
	var division string
	var college_name string

	fmt.Println("Enter Name : ")
	fmt.Scan(&name)
	fmt.Println("Enter rollno : ")
	fmt.Scan(&rollno)

	fmt.Println("Enter Division : ")
	fmt.Scan(&division)

	fmt.Println("Enter college name : ")
	fmt.Scan(&college_name)

	fmt.Printf("%10s %10s %10s %10s\n", "Name", "Roll NO", "Division", "CollegeName")
	fmt.Printf("%10s %10d %10s %10s\n", name, rollno, division, college_name)

}
