// wap in go language to accept n records of employee information(eno,ename,salary) and display record of employees
// having maximum salary

package main

import (
	"fmt"
)

type employee struct {
	eno    int
	ename  string
	salary float32
}

func display(emp []employee, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%10d %10s %10.2f\n", emp[i].eno, emp[i].ename, emp[i].salary)
	}
}
func MaxSalary(e []employee, n int) {
	var max float32
	for i := 0; i < n; i++ {
		if max < e[i].salary {
			max = e[i].salary
		}
	}
	for i := 0; i < n; i++ {
		if max == e[i].salary {
			fmt.Println("Max Salary Employee : ")
			fmt.Printf("%10d %10s %10.2f\n", e[i].eno, e[i].ename, e[i].salary)
		}
	}

}
func main() {
	var n int
	emp := make([]employee, 10, 30)
	fmt.Println("How many employees you want to enter")
	fmt.Scan(&n)
	fmt.Println("Enter employees details")

	for i := 0; i < n; i++ {
		fmt.Println("Enter eno")
		fmt.Scan(&emp[i].eno)
		fmt.Println("Enter Name")
		fmt.Scan(&emp[i].ename)
		fmt.Println("Enter Salary")
		fmt.Scan(&emp[i].salary)
		emp = append(emp, emp[i])
	}
	fmt.Printf("%10s %10s %10s\n", "Eno", "Ename", "Salary")
	display(emp, n)
	MaxSalary(emp, n)

}
