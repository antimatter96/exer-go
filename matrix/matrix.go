package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a named alias for `[][]int`, since methods can't be
// defined for `[][]int`
type Matrix [][]int

// New creates a new Matrix from a string of newline and space seperated numbers
func New(s string) (Matrix, error) {
	var arr [][]int
	x := strings.Split(s, "\n")

	rows := len(x)
	if rows == 0 {
		return arr, fmt.Errorf("No rows")
	}
	cols := len(strings.Split(x[0], " "))
	if cols == 0 {
		return arr, fmt.Errorf("No cols")
	}
	var err error
	arr = make([][]int, rows)
	for i := 0; i < rows; i++ {
		trimmed := strings.Trim(x[i], " ")
		seperated := strings.Split(trimmed, " ")
		if len(seperated) != cols {
			return arr, fmt.Errorf("Non equal cols")
		}

		arr[i] = make([]int, cols)
		for k, v := range seperated {
			arr[i][k], err = strconv.Atoi(v)
			if err != nil {
				return arr, fmt.Errorf("Parse Error %v", v)
			}
		}
	}

	return arr, nil

}

// Rows returns the rows of the matrix
func (m *Matrix) Rows() [][]int {
	var arr [][]int
	rows := len(*m)
	if rows == 0 {
		return arr
	}
	cols := len((*m)[0])
	if cols == 0 {
		return arr
	}

	arr = make(Matrix, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
	}

	for i := 0; i < len(*m); i++ {
		for j := 0; j < cols; j++ {
			arr[i][j] = (*m)[i][j]
		}
	}
	return arr
}

// Cols returns the columns of the matrix
func (m *Matrix) Cols() [][]int {
	var arr [][]int
	newrows := len((*m)[0])
	if newrows == 0 {
		return arr
	}
	newcols := len(*m)
	if newcols == 0 {
		return arr
	}

	arr = make(Matrix, newrows)
	for i := 0; i < newrows; i++ {
		arr[i] = make([]int, newcols)
	}

	for i := 0; i < newrows; i++ {
		for j := 0; j < newcols; j++ {
			arr[i][j] = (*m)[j][i]
		}
	}

	return arr
}

// Set is used to change the value of given cell in the matrix
func (m *Matrix) Set(row, col, value int) bool {
	rows := len(*m)
	if rows == 0 {
		return false
	}
	cols := len((*m)[0])
	if cols == 0 {
		return false
	}

	if row < 0 || row >= rows || col < 0 || col >= cols {
		return false
	}
	(*m)[row][col] = value

	return true
}
