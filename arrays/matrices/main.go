// rows := [][]float65{}
// that's the syntax of creating a matrix

// complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers
// where the value of each cell is i * j where i and j are the indexes of the row and column, respectively.

package main

import "fmt"

func createMatrix(rows, cols int) [][]int {

	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix

}

func main() {

	fmt.Printf("%v", createMatrix(10, 10))
}
