package matrix

// Retrieve value at [row, col] safely.
func (A *Matrix) SafeGet(row, col int) (float64, error) {

	if A.isOutOfBounds(row, col) {
		return 0, ErrOutOfBounds
	}

	return A.Get(row, col), nil
}

// Multiplies 2 matricies with each other safely.
// Returns a new matrix.
func (A *Matrix) SafeMul(B *Matrix) (*Matrix, error) {
	if !isSquare(A, B) {
		return nil, ErrIncompatibleSizes
	}

	return A.Mul(B), nil
}

// Set the element at [row, col] to val safely.
func (A *Matrix) SafeSet(row, col int, val float64) error {
	if A.isOutOfBounds(row, col) {
		return ErrOutOfBounds
	}

	A.Set(row, col, val)

	return nil
}

// Calculates the element-wise product of 2 matricies safely.
// Returns a new matrix.
func (A *Matrix) SafeEWProd(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	return A.EWProd(B), nil
}

// Insert the given rows into the matrix, returning a new matrix.
// Passing 0 as the second argument is like making the
// passed rows the first few, whereas passing Rows() is like appending
// the additional rows to the matrix.
func (A *Matrix) SafeInsertRows(rows *Matrix, afterRow int) (*Matrix, error) {
	if rows.Cols != A.Cols {
		return nil, ErrIncompatibleSizes
	}

	if afterRow < 0 || afterRow > A.Rows {
		return nil, ErrInsertionOutOfBounds
	}

	return A.InsertRows(rows, afterRow), nil
}

// Insert the given columns into the matrix, returning a new matrix.
// Passing 0 as the second argument is like making the
// passed columns the first few (on the left), whereas passing Columns() is like appending
// the additional columns to the matrix (on the right).
func (A *Matrix) SafeInsertColumns(cols *Matrix, afterCol int) (*Matrix, error) {
	if cols.Rows != A.Rows {
		return nil, ErrIncompatibleSizes
	}

	if afterCol < 0 || afterCol > A.Cols {
		return nil, ErrInsertionOutOfBounds
	}

	return A.InsertColumns(cols, afterCol), nil
}

func (A *Matrix) isOutOfBounds(row, col int) bool {
	index := (row-1)*A.Cols + col - 1

	if index >= len(A.Vals) {
		return true
	}

	return false
}

func sameSize(A, B *Matrix) bool {
	return A.Rows == B.Rows && A.Cols == B.Cols
}

func isSquare(A, B *Matrix) bool {
	return A.Cols == B.Rows
}
