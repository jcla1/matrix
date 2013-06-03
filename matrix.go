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

var (
	ErrIncompatibleSizes    = &MatrixError{"Incompatible sizes of matricies"}
	ErrInsertionOutOfBounds = &MatrixError{"The matrix you are trying to insert to hasn't got that row or column."}
	ErrOutOfBounds          = &MatrixError{"The element you are trying to access is out of bounds."}
)

// Describes error during calculations
type MatrixError struct {
	ErrorString string
}

// Return the Error's string
func (err *MatrixError) Error() string { return err.ErrorString }

// Matrix construct that holds all the information about a matrix
type Matrix struct {
	rows, cols int
	values     []float64
}

// A function that will apply an abitary transformation to an element in the matrix
type ApplyFunc func(index int, value float64) float64

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
//
//	A := matrix.FromMatlab("[1 2 3; 4 5 6]")
func FromMatlab(str string) *Matrix {
	rows := strings.Split(str, ";")

	for i, row := range rows {
		rows[i] = strings.Replace(row, ",", " ", -1)
	}

	nRows := len(rows)
	nColumns := len(strings.Fields(rows[0]))

	A := Zeros(nRows, nColumns)

	for i, row := range rows {
		row = strings.Trim(row, "[] ")
		strNums := strings.Fields(row)

		for j, num := range strNums {
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

// Returns a string representation of the matrix
func (A *Matrix) String() string {
	buffer := new(bytes.Buffer)

	for i, elem := range A.values {
		buffer.WriteString(fmt.Sprintf("%.3f ", elem))

		if (i+1)%A.cols == 0 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

// Returns an exact copy of the matrix
func (A *Matrix) Copy() *Matrix {
	B := Zeros(A.rows, A.cols)
	copy(B.values, A.values)

	return B
}

// Insert the given rows into the matrix, returning a new matrix.
// Passing 0 as the second argument is like making the
// passed rows the first few, whereas passing Rows() is like appending
// the additional rows to the matrix.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeInsertRows
func (A *Matrix) InsertRows(rows *Matrix, afterRow int) *Matrix {
	B := Zeros(A.rows+rows.rows, A.cols)

	// copy rows before inserted rows
	copy(B.values[0:afterRow*B.cols], A.values[0:afterRow*B.cols])
	// insert new rows
	copy(B.values[afterRow*B.cols:afterRow*B.cols+rows.rows*B.cols], rows.values)
	// copy rows after inserted rows
	copy(B.values[afterRow*B.cols+rows.rows*B.cols:], A.values[afterRow*B.cols:])

	return B
}

func (A *Matrix) InsertColumns(cols *Matrix, afterCol int) *Matrix {
	return Zeros(1, 1)
}

// Give a function, apply its transformation to every element in the matrix.
//
//	A := matrix.Rand(4, 8)
//	A.Apply(func(index int, value float64) float64 {
//		return sigmoid(value)
//	})
func (A *Matrix) Apply(f ApplyFunc) *Matrix {
	for i, v := range A.values {
		A.values[i] = f(i, v)
	}

	return A
}

// Get the element at [row, col].
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeGet
func (A *Matrix) Get(row, col int) float64 {
	return A.values[(row-1)*A.cols+col-1]
}

// Set the element at [row, col] to val.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeSet
func (A *Matrix) Set(row, col int, val float64) {

	A.values[(row-1)*A.cols+col-1] = val

}

// Transpose the matrix
func (A *Matrix) Transpose() *Matrix {
	B := Zeros(A.cols, A.rows)

	for i := 1; i <= A.rows; i++ {
		for j := 1; j <= A.cols; j++ {
			B.values[(j-1)*B.cols+i-1] = A.values[(i-1)*A.cols+j-1]
		}
	}

	return B
}

// Add B to the matrix A, produces new matrix
func (A *Matrix) Add(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	C := Zeros(A.rows, A.cols)

	for i, val := range A.values {
		C.values[i] = val + B.values[i]
	}

	return C, nil
}

// Subtract B from the matrix A, produces a new matrix
func (A *Matrix) Sub(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	C := Zeros(A.rows, A.cols)

	for i, val := range A.values {
		C.values[i] = val - B.values[i]
	}

	return C, nil
}

// Multiplies 2 matricies with each other.
// Returns a new matrix.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeMul
func (A *Matrix) Mul(B *Matrix) *Matrix {

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

	return C
}

// Standard scalar product of 2 matricies.
// Returns a new matrix.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeDot
func (A *Matrix) Dot(B *Matrix) *Matrix {
	C := Zeros(A.rows, A.cols)

	for i, val := range A.values {
		C.values[i] = val * B.values[i]
	}

	return C
}

// Scale the matrix by the factor f
func (A *Matrix) Scale(f float64) *Matrix {
	B := Zeros(A.rows, A.cols)

	for i, val := range A.values {
		B.values[i] = val * f
	}

	return B
}

// Take every element of the matrix to the power of n
func (A *Matrix) Power(n float64) *Matrix {
	B := Zeros(A.rows, A.cols)

	for i, val := range A.values {
		B.values[i] = math.Pow(val, n)
	}

	return B
}

// Add n to all elements in the matrix (in-place)
func (A *Matrix) AddNum(n float64) *Matrix {
	for i := range A.values {
		A.values[i] += n
	}

	return A
}
