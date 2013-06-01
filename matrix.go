// Package matrix provides functions for simple linear algebra
package matrix

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type ErrOutOfBounds bool
type ErrIncompatibleSize bool

func (e ErrOutOfBounds) Error() string {
	return "No value there! Out of range!"
}

func (e ErrIncompatibleSize) Error() string {
	return "The matricies aren't the right dimensions."
}

type Matrix struct {
	rows, cols int
	values     []float64
}

// Creates a new Matrix and initializes all values to 0
func Zeros(rows, cols int) *Matrix {
	A := new(Matrix)

	A.rows = rows
	A.cols = cols
	A.values = make([]float64, rows*cols)

	return A
}

// Creates a new Matrix and initializes all values to 1
func Ones(rows, cols int) *Matrix {
	A := Zeros(rows, cols)
	A.AddNum(1)

	return A
}

// Creates a new identity matrix
func Eye(size int) *Matrix {
	A := Zeros(size, size)

	for i := 1; i <= size; i++ {
		A.Set(i, i, 1)
	}

	return A
}

// Constructs a new matrix with random values in range [0;1)
func Rand(rows, cols int) *Matrix {
	A := Zeros(rows, cols)

	for i := range A.values {
		A.values[i] = rand.Float64()
	}

	return A
}

// Constructs a new Matrix from a Matlab style representation
func FromMatlab(str string) *Matrix {
	rows := strings.Split(str, ";")

	for i, row := range rows {
		rows[i] = strings.Replace(row, ",", " ", -1)
	}

	n_rows := len(rows)
	n_columns := len(strings.Fields(rows[0]))

	A := Zeros(n_rows, n_columns)

	for i, row := range rows {
		row = strings.Trim(row, "[] ")
		str_nums := strings.Fields(row)

		for j, num := range str_nums {
			n, _ := strconv.ParseFloat(num, 64)
			A.Set(i+1, j+1, n)
		}
	}

	return A
}

// Return a Matlab representation of the matrix
func (A *Matrix) ToMatlab() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("[")

	for i, v := range A.values {
		buffer.WriteString(fmt.Sprintf("%f ", v))

		if (i+1)%A.cols == 0 {
			buffer.WriteString("; ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}

// Gives the dimensions of the matrix
func (A *Matrix) Dim() (int, int) {
	return A.rows, A.cols
}

// Return the number of rows in the matrix
func (A *Matrix) Rows() int {
	return A.rows
}

// Return the number of columns in the matrix
func (A *Matrix) Columns() int {
	return A.cols
}

// Returns an array of all the values
func (A *Matrix) Values() []float64 {
	tmp := make([]float64, len(A.values))
	copy(tmp, A.values)
	return tmp
}

// Returns an exact copy of the matrix
func (A *Matrix) Copy() *Matrix {
	B := Zeros(A.rows, A.cols)
	copy(B.values, A.values)

	return B
}

// Returns a string representation of the matrix
func (A *Matrix) String() string {
	buffer := new(bytes.Buffer)

	for i, elem := range A.values {
		buffer.WriteString(fmt.Sprintf("%f ", elem))

		if (i+1)%A.cols == 0 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

// Retrieve value at [row, col]
func (A *Matrix) Get(row, col int) (float64, error) {

	if A.isOutOfBounds(row, col) {
		return 0, ErrOutOfBounds(true)
	}

	return A.values[(row-1)*A.cols+col-1], nil
}

// Set the element at [row, col] to val
func (A *Matrix) Set(row, col int, val float64) error {

	if A.isOutOfBounds(row, col) {
		return ErrOutOfBounds(true)
	}

	A.values[(row-1)*A.cols+col-1] = val

	return nil
}

// Transpose the matrix (in-place, costly)
func (A *Matrix) Transpose() *Matrix {
	B := Zeros(A.cols, A.rows)

	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			v, _ := A.Get(i, j)
			B.Set(j, i, v)
		}
	}

	A.rows, A.cols = A.cols, A.rows
	copy(A.values, B.values)
	return A
}

// Add B to the matrix A (in-place)
func (A *Matrix) Add(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSize(true)
	}

	for i, val := range B.values {
		A.values[i] += val
	}

	return A, nil
}

// Subtract B from the matrix A (in-place)
func (A *Matrix) Sub(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSize(true)
	}

	for i, val := range B.values {
		A.values[i] -= val
	}

	return A, nil
}

// Multiply 2 matricies with each other returning a new matrix
func (A *Matrix) Mul(B *Matrix) (*Matrix, error) {
	if !columnIsRow(A, B) {
		return nil, ErrIncompatibleSize(true)
	}

	C := Zeros(A.rows, B.cols)

	for i := 0; i < C.rows; i++ {
		for j := 0; j < C.cols; j++ {
			sum := float64(0)

			for k := 0; k < A.cols; k++ {
				sum += A.values[i*A.cols+k] * B.values[k*B.cols+j]
			}

			C.values[i*C.rows+j] = sum
		}
	}

	return C, nil
}

// Scale the matrix in-place by the factor f
func (A *Matrix) Scale(f float64) *Matrix {
	for i := range A.values {
		A.values[i] *= f
	}

	return A
}

// Take every element of the matrix to the power of n (in-place)
func (A *Matrix) Power(n float64) *Matrix {
	for i, elem := range A.values {
		A.values[i] = math.Pow(elem, n)
	}

	return A
}

// Add n to all elements in the matrix (in-place)
func (A *Matrix) AddNum(n float64) *Matrix {
	for i := range A.values {
		A.values[i] += n
	}

	return A
}

func (A *Matrix) isOutOfBounds(row, col int) bool {
	index := (row-1)*A.cols + col - 1

	if index >= len(A.values) {
		return true
	}

	return false
}

func sameSize(A, B *Matrix) bool {
	return A.rows == B.rows && A.cols == B.cols
}

func columnIsRow(A, B *Matrix) bool {
	return A.cols == B.rows
}
