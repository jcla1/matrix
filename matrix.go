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
	Rows, Cols int
	Values     []float64
}

// A function that will apply an abitary transformation to an element in the matrix
type ApplyFunc func(index int, value float64) float64

// Creates a new Matrix and initializes all values to 0
func Zeros(rows, cols int) *Matrix {
	A := new(Matrix)

	A.Rows = rows
	A.Cols = cols
	A.Values = make([]float64, rows*cols)

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

	for i := range A.Values {
		A.Values[i] = rand.Float64()
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

	for i, v := range A.Values {
		buffer.WriteString(fmt.Sprintf("%f ", v))

		if (i+1)%A.Cols == 0 {
			buffer.WriteString("; ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}

func FromSlice(s []float64, rows, cols int) *Matrix {
	A := Zeros(rows, cols)
	copy(A.Values, s)

	return A
}

// Gives the dimensions of the matrix
func (A *Matrix) Dim() (int, int) {
	return A.Rows, A.Cols
}

// Return the number of rows in the matrix
func (A *Matrix) R() int {
	return A.Rows
}

// Return the number of columns in the matrix
func (A *Matrix) C() int {
	return A.Cols
}

// Returns an array of all the values
func (A *Matrix) Values() []float64 {
	tmp := make([]float64, len(A.Values))
	copy(tmp, A.Values)
	return tmp
}

// Returns a string representation of the matrix
func (A *Matrix) String() string {
	buffer := new(bytes.Buffer)

	for i, elem := range A.Values {
		buffer.WriteString(fmt.Sprintf(" %.3f ", elem))

		if (i+1)%A.Cols == 0 && i+1 != len(A.Values) {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

// Returns an exact copy of the matrix
func (A *Matrix) Copy() *Matrix {
	B := Zeros(A.Rows, A.Cols)
	copy(B.Values, A.Values)

	return B
}

func (A *Matrix) Unroll() *Matrix {
	return A.Reshape(A.Rows*A.Cols, 1)
}

// Reshapes the matrix. Make sure they have the same number of elements
func (A *Matrix) Reshape(rows, cols int) *Matrix {
	B := Zeros(rows, cols)
	copy(B.Values, A.Values)

	return B
}

// Insert the given rows into the matrix, returning a new matrix.
// Passing 0 as the second argument is like making the
// passed rows the first few, whereas passing R() is like appending
// the additional rows to the matrix.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeInsertRows
func (A *Matrix) InsertRows(rows *Matrix, afterRow int) *Matrix {
	B := Zeros(A.Rows+rows.Rows, A.Cols)

	// copy rows before inserted rows
	copy(B.Values[0:afterRow*B.Cols], A.Values[0:afterRow*B.Cols])
	// insert new rows
	copy(B.Values[afterRow*B.Cols:afterRow*B.Cols+rows.Rows*B.Cols], rows.Values)
	// copy rows after inserted rows
	copy(B.Values[afterRow*B.Cols+rows.Rows*B.Cols:], A.Values[afterRow*B.Cols:])

	return B
}

// Remove a single row from the matrix.
// The indexing for the rows start with 1 and go upto A.R()
func (A *Matrix) RemoveRow(row int) *Matrix {
	B := Zeros(A.Rows-1, A.Cols)
	copy(B.Values, append(append([]float64{}, A.Values[:(row-1)*A.Cols]...), A.Values[(row-1)*A.Cols+A.Cols:]...))
	return B
}

// Insert the given columns into the matrix, returning a new matrix.
// Passing 0 as the second argument is like making the
// passed columns the first few (on the left), whereas passing Columns() is like appending
// the additional columns to the matrix (on the right).
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeInsertColumns
func (A *Matrix) InsertColumns(cols *Matrix, afterCol int) *Matrix {
	B := Zeros(A.Rows, A.Cols+cols.Cols)

	for i := 0; i < A.Rows; i++ {
		copy(B.Values[i*B.Cols:i*B.Cols+afterCol], A.Values[i*A.Cols:i*A.Cols+afterCol])
		copy(B.Values[i*B.Cols+afterCol:i*B.Cols+afterCol+cols.Cols], cols.Values[i*cols.Cols:(i+1)*cols.Cols])
		copy(B.Values[i*B.Cols+afterCol+cols.Cols:(i+1)*B.Cols], A.Values[i*A.Cols+afterCol:(i+1)*A.Cols])
	}

	return B
}

// Give a function, apply its transformation to every element in the matrix.
//
//	A := matrix.Rand(4, 8)
//	A.Apply(func(index int, value float64) float64 {
//		return sigmoid(value)
//	})
func (A *Matrix) Apply(f ApplyFunc) *Matrix {
	for i, v := range A.Values {
		A.Values[i] = f(i, v)
	}

	return A
}

// Get the element at [row, col].
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeGet
func (A *Matrix) Get(row, col int) float64 {
	return A.Values[(row-1)*A.Cols+col-1]
}

// Set the element at [row, col] to val.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeSet
func (A *Matrix) Set(row, col int, val float64) {

	A.Values[(row-1)*A.Cols+col-1] = val

}

// Transpose the matrix
func (A *Matrix) Transpose() *Matrix {
	B := Zeros(A.Cols, A.Rows)

	for i := 1; i <= A.Rows; i++ {
		for j := 1; j <= A.Cols; j++ {
			B.Values[(j-1)*B.Cols+i-1] = A.Values[(i-1)*A.Cols+j-1]
		}
	}

	return B
}

// Add B to the matrix A, produces new matrix
func (A *Matrix) Add(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	C := Zeros(A.Rows, A.Cols)

	for i, val := range A.Values {
		C.Values[i] = val + B.Values[i]
	}

	return C, nil
}

// Subtract B from the matrix A, produces a new matrix
func (A *Matrix) Sub(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	C := Zeros(A.Rows, A.Cols)

	for i, val := range A.Values {
		C.Values[i] = val - B.Values[i]
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

	C := Zeros(A.Rows, B.Cols)

	for i := 0; i < C.Rows; i++ {
		for j := 0; j < C.Cols; j++ {
			sum := float64(0)

			for k := 0; k < A.Cols; k++ {
				sum += A.Values[i*A.Cols+k] * B.Values[k*B.Cols+j]
			}

			C.Values[i*C.Cols+j] = sum
		}
	}

	return C
}

func (A *Matrix) Dot(B *Matrix) float64 {
	sum := 0.0
	for i, v := range A.Values {
		sum += v * B.Values[i]
	}

	return sum
}

// Element-wise multiplication of the matricies
// Returns a new matrix.
//
// Warning: This is an unsafe method to use, it does no boundary
// checking what so ever. If you'd like a safe version
// use: SafeEWProd
func (A *Matrix) EWProd(B *Matrix) *Matrix {
	C := Zeros(A.Rows, A.Cols)

	for i, val := range A.Values {
		C.Values[i] = val * B.Values[i]
	}

	return C
}

// Scale the matrix by the factor f
func (A *Matrix) Scale(f float64) *Matrix {
	B := Zeros(A.Rows, A.Cols)

	for i, val := range A.Values {
		B.Values[i] = val * f
	}

	return B
}

// Take every element of the matrix to the power of n
func (A *Matrix) Power(n float64) *Matrix {
	B := Zeros(A.Rows, A.Cols)

	for i, val := range A.Values {
		B.Values[i] = math.Pow(val, n)
	}

	return B
}

// Add n to all elements in the matrix (in-place)
func (A *Matrix) AddNum(n float64) *Matrix {
	for i := range A.Values {
		A.Values[i] += n
	}

	return A
}
