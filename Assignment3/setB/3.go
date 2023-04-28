// wap in go language to accept n student details like roll_no,stud_name, mark1,mark2,mark3.
// Calculate the total and average of marks using structure

package main

import "fmt"

type studentDetails struct {
	roll_no   int
	stud_name string
	mark1     float32
	mark2     float32
	mark3     float32
	average   float32
	total     float32
}

func main() {
	var n int
	studentObjects := make([]studentDetails, 10, 20)

	fmt.Println("How many Students Details you want to enter (max 10)")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter Student %v RollNO : \n", i+1)
		fmt.Scan(&studentObjects[i].roll_no)
		fmt.Printf("Enter Student %v Name : \n", i+1)
		fmt.Scan(&studentObjects[i].stud_name)

		fmt.Printf("Enter Student %v Mark1 : \n", i+1)
		fmt.Scan(&studentObjects[i].mark1)

		fmt.Printf("Enter Student %v Mark2 : \n", i+1)
		fmt.Scan(&studentObjects[i].mark2)

		fmt.Printf("Enter Student %v Mark3 : \n", i+1)
		fmt.Scan(&studentObjects[i].mark3)
		studentObjects[i].average = (studentObjects[i].mark1 + studentObjects[i].mark2 + studentObjects[i].mark3) / 3

		studentObjects[i].total = (studentObjects[i].mark1 + studentObjects[i].mark2 + studentObjects[i].mark3)

	}

	fmt.Println("Displaying Alldetails of students in tabular form : ")
	fmt.Printf("%10s %10s %10s %10s %10s %10s %10s\n", "RollNo", "Name", "Mark1", "Mark2", "Mark3", "Average", "Total")
	for i := 0; i < n; i++ {
		fmt.Printf("%10d %10s %10.2f %10.2f %10.2f %10.2f %10.2f\n", studentObjects[i].roll_no, studentObjects[i].stud_name, studentObjects[i].mark1, studentObjects[i].mark2, studentObjects[i].mark3, studentObjects[i].average, studentObjects[i].total)

	}
}
