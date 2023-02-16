package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	for i := 0; i < len(rows); i++ {
		cols := strings.Fields(rows[i])
		if i > 0 && len(cols) != len(matrix[0]) {
			return nil, errors.New("invalid input")
		}

		matrix[i] = make([]int, len(cols))

		for j := 0; j < len(cols); j++ {
			val, err := strconv.Atoi(cols[j])
			if err != nil {
				return nil, errors.New("invalid input")
			}
			matrix[i][j] = val
		}
	}
	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	matrix := make(Matrix, len(m[0]))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, len(m))
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = m[j][i]
		}
	}

	return matrix
}

func (m Matrix) Rows() [][]int {
	matrix := make(Matrix, len(m))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, len(m[0]))
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = m[i][j]
		}
	}

	return matrix
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}
