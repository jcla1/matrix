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

// Calculates the standard scalar product of 2 matrixies safely.
// Returns a new matrix.
func (A *Matrix) SafeDot(B *Matrix) (*Matrix, error) {
	if !sameSize(A, B) {
		return nil, ErrIncompatibleSizes
	}

	return A.Dot(B), nil
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

func isSquare(A, B *Matrix) bool {
	return A.cols == B.rows
}
