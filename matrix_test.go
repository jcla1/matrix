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
	A := Eye(3)
	expected_values := []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}

	if !arraysIdentical(A.values, expected_values) {
		t.Error("Unexpected value")
	}
}

func TestMatrixRand(t *testing.T) {
	A := Rand(4, 5)

	for _, v := range A.Values() {
		if v >= 1.0 || v < 0.0 {
			t.Error("Random number out of range!")
		}
	}
}

func TestMatrixDim(t *testing.T) {
	A := Zeros(44, 32)
	r, c := A.Dim()
	if r != 44 || c != 32 {
		t.Error("Dimention mismatch")
	}
}

func TestMatrixGet(t *testing.T) {
	A := Ones(8, 5).AddNum(10)
	val, err := A.SafeGet(8, 5)

	if err != nil {
		t.Error("There was an error getting the value:", err)
	}

	if val != 11 {
		t.Error("Wrong value was retrieved expected 11, but got:", val)
	}
}

func TestMatrixSet(t *testing.T) {
	A := Zeros(5, 3)
	A.SafeSet(4, 3, 10)

	expected_values := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0}

	if !arraysIdentical(A.Values(), expected_values) {
		t.Error("Unexpected value")
	}
}

func TestMatrixAdd(t *testing.T) {
	A := Rand(3, 5)
	B := Rand(3, 5)

	a_values := A.Values()
	b_values := B.Values()

	C, _ := A.Add(B)

	for i, v := range C.Values() {
		if v != (a_values[i] + b_values[i]) {
			t.Error("Wrong value in addition")
			break
		}
	}
}

func TestMatrixSub(t *testing.T) {
	A := Rand(4, 7)
	B := Rand(4, 7)

	a_values := A.Values()
	b_values := B.Values()

	C, _ := A.Sub(B)

	for i, v := range C.Values() {
		if v != (a_values[i] - b_values[i]) {
			t.Error("Wrong value in subtraction")
			break
		}
	}
}

func TestMatrixMul(t *testing.T) {
	A := Eye(5)
	B := FromMatlab("[1 2 3 4 5; 6 7 8 9 10; 11 12 13 14 15; 16 17 18 19 20; 21 22 23 24 25]")

	C, _ := A.SafeMul(B)
	values := C.Values()

	expected_values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	if !arraysIdentical(values, expected_values) {
		t.Error("Wrong value in Multiplication with Identity matrix")
	}

	D := FromMatlab("[555,7068;5567,9370;3498,25]")
	E := FromMatlab("[4840,4607,1493;9326,1652,4872]")

	F, _ := D.SafeMul(E)
	values2 := F.Values()

	expected_values2 := []float64{68602368, 14233221, 35263911, 114328900, 41126409, 53962171, 17163470, 16156586, 5344314}

	if !arraysIdentical(values2, expected_values2) {
		t.Error("Wrong value in Multiplication with matrix")
	}
}

func TestMatrixDot(t *testing.T) {
	A := Ones(3, 3).AddNum(3)
	B := Ones(3, 3).AddNum(2)

	C, _ := A.SafeDot(B)

	for _, val := range C.values {
		if val != 12 {
			t.Error("Error in scalar multiplication")
		}
	}
}

func TestMatrixPower(t *testing.T) {
	A := Zeros(9, 4).AddNum(2).Power(2)

	for _, v := range A.Values() {
		if v != 4 {
			t.Error("Wrong result, expected 4 got:", v)
			break
		}
	}
}

func TestMatrixTranspose(t *testing.T) {
	values := FromMatlab("[1 2; 3 4; 5 6]").Transpose().Values()

	expected_values := []float64{1, 3, 5, 2, 4, 6}

	if !arraysIdentical(values, expected_values) {
		t.Error("Unexpected value occurred")
	}
}

func TestMatrixScale(t *testing.T) {
	vals := Ones(5, 4).Scale(5).Values()

	for _, v := range vals {
		if v != 5 {
			t.Error("Unexpected value occurred")
			break
		}
	}
}

func TestMatrixAddNum(t *testing.T) {
	vals := Zeros(3, 9).AddNum(42).Values()

	for _, v := range vals {
		if v != 42 {
			t.Error("Unexpected value occurred")
			break
		}
	}
}

func arraysIdentical(arr1, arr2 []float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}

	return true
}
