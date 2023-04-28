// wap in go language to accept two matrices and display its multiplication

package main

import "fmt"

func main() {
	var rows1 int
	var cols1 int
	var rows2 int
	var cols2 int
	var rows3 int
	var cols3 int
	var matrix1 [5][5]int
	var matrix2 [5][5]int
	var validateDimensions bool
	fmt.Println("Enter rows and columns / dimensions of first matrix")
	fmt.Printf("Matrix %v : ", 1)
	fmt.Printf("\nRow  : \n")
	fmt.Scan(&rows1)
	fmt.Printf("\nCols  : \n")
	fmt.Scan(&cols1)
	fmt.Printf("Matrix %v : ", 2)
	fmt.Printf("\nRow  : \n")
	fmt.Scan(&rows2)
	fmt.Printf("\nCols  : \n")
	fmt.Scan(&cols2)
	rows3 = rows1
	cols3 = cols2
	fmt.Println("Checkings Dimensions of matrises")
	validateDimensions = cols1 != rows2

	if validateDimensions {
		fmt.Println("In Matrix 1 Number of Columns is not equal to Number of Rows of Matrix 2 or viseversa")
	} else {
		fmt.Println("Enters elements of matrix 1")
		for i := 0; i < rows1; i++ {
			for j := 0; j < cols1; j++ {
				fmt.Scan(&matrix1[i][j])

			}
		}
		fmt.Println("Enters elements of matrix 2")
		for i := 0; i < rows2; i++ {
			for j := 0; j < cols2; j++ {
				fmt.Scan(&matrix2[i][j])

			}
		}

		for i := 0; i < rows1; i++ {
			for j := 0; j < cols2; j++ {
				fmt.Printf(" matrix1 %v%v = %v and matrix2 %v%v = %v\n ", i, j, matrix1[i][j], i, j, matrix2[i][j])

			}
		}

		for i := 0; i < rows3; i++ {
			for j := 0; j < cols3; j++ {
				fmt.Printf("%v ", matrix1[i][j]*matrix2[j][j]+matrix1[i][j+1]*matrix2[j+1][j])
			}
			fmt.Println()
		}
	}

}
