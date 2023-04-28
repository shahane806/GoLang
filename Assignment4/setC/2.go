// wap in go language to store n student information(rollno,nae,percentage)and write a method to display student information in descdnding order of percentage

package main

import "fmt"

type student struct {
	rollno     int
	name       string
	percentage float32
}

func sortingWithPercentage(s []student, n int) []student {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j].percentage < s[j+1].percentage {
				var temp int
				temp = s[j].rollno
				s[j].rollno = s[j+1].rollno
				s[j+1].rollno = temp
				temp = int(s[j].percentage)
				s[j].percentage = s[j+1].percentage
				s[j+1].percentage = float32(temp)
				var temp2 string
				temp2 = (s[j].name)
				s[j].name = s[j+1].name
				s[j+1].name = temp2

			}
		}

	}
	return s

}
func main() {
	var n int
	stud := make([]student, 5, 10)

	fmt.Println("How many student details you want to enter")
	fmt.Scan(&n)
	fmt.Println("Enter Student Details")
	for i := 0; i < n; i++ {
		fmt.Println("Enter Student rollno")
		fmt.Scan(&stud[i].rollno)
		fmt.Println("Enter Student name")
		fmt.Scan(&stud[i].name)
		fmt.Println("Enter Student percentage")
		fmt.Scan(&stud[i].percentage)
		stud = append(stud, stud[i])
	}
	fmt.Println("Displaying Student Details")
	fmt.Printf("%10s %10s %10s\n", "RollNo", "Name", "Percentage")
	for i := 0; i < n; i++ {
		fmt.Printf("%10v %10v %10v\n", stud[i].rollno, stud[i].name, stud[i].percentage)
	}
	stud = sortingWithPercentage(stud, n)
	fmt.Println("Displaying Student Details Sorted in Descending order as per percentage")
	fmt.Printf("%10s %10s %10s\n", "RollNo", "Name", "Percentage")
	for i := 0; i < n; i++ {
		fmt.Printf("%10v %10v %10v\n", stud[i].rollno, stud[i].name, stud[i].percentage)
	}
}
