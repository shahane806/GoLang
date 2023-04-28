package main

import "fmt"

type student struct {
	name       string
	rollno     int
	percentage int
}

func main() {
	var n int
	var stud [10]student
	fmt.Println("How many record you want to enter")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter student %d , info\n", i+1)
		fmt.Printf("Roll NO  : ")
		fmt.Scan(&stud[i].rollno)
		fmt.Printf("Name : ")
		fmt.Scan(&stud[i].name)
		fmt.Printf("Percentage  : ")
		fmt.Scan(&stud[i].percentage)

	}
	fmt.Println("Display Details of all students")
	fmt.Printf("%20s %20s %20s\n", "Roll no", "Name", "Percentage")
	for i := 0; i < n; i++ {
		fmt.Printf("%20v %20v %20v\n", stud[i].rollno, stud[i].name, stud[i].percentage)

	}
	fmt.Println("In Descending Order of Percentage :::")
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if stud[j].percentage < stud[j+1].percentage {
				temp := stud[j].percentage
				stud[j].percentage = stud[j+1].percentage
				stud[j+1].percentage = temp
				temp2 := stud[j].rollno
				stud[j].rollno = stud[j+1].rollno
				stud[j+1].rollno = temp2
				temp3 := stud[j].name
				stud[j].name = stud[j+1].name
				stud[j+1].name = temp3

			}
		}
	}
	fmt.Println("Display Details of all students in descending order or percentage")
	fmt.Printf("%20s %20s %20s\n", "Roll no", "Name", "Percentage")
	for i := 0; i < n; i++ {
		fmt.Printf("%20v %20v %20v\n", stud[i].rollno, stud[i].name, stud[i].percentage)

	}
}
