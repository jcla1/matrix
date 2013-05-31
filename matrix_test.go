package matrix

import (
	"testing"
)

func TestMatrixZeros(t *testing.T) {
	A := Zeros(10, 8)
	vals := A.Values()

	if len(vals) != 10*8 {
		t.Error("Wrong number of values")
	}

	for _, v := range vals {
		if v != 0 {
			t.Error("Value isn't 0")
			break
		}
	}
}

func TestMatrixOnes(t *testing.T) {
	A := Ones(5, 14)
	vals := A.Values()

	if len(vals) != 5*14 {
		t.Error("Wrong number of values")
	}

	for _, v := range vals {
		if v != 1 {
			t.Error("Value isn't 1")
			break
		}
	}
}

func TestMatrixEye(t *testing.T) {
	A := Eye(8)

	for i, v := range A.Values() {
		if v == 1 {
			if (i+1)%A.cols != (i+1)/A.cols {
				t.Error("Value not in right position")
				break
			}
		} else if v == 0 {

		} else {
			t.Error("Value besides 0 or 1 encountered")
			break
		}
	}
}
func TestMatrixGet(t *testing.T)       {}
func TestMatrixSet(t *testing.T)       {}
func TestMatrixAdd(t *testing.T)       {}
func TestMatrixSub(t *testing.T)       {}
func TestMatrixMul(t *testing.T)       {}
func TestMatrixPower(t *testing.T)     {}
func TestMatrixDim(t *testing.T)       {}
func TestMatrixTranspose(t *testing.T) {}
func TestMatrixScale(t *testing.T)     {}
func TestMatrixAddNum(t *testing.T)    {}
