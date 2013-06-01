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

	for i, v := range A.Values() {
		if v != expected_values[i] {
			t.Error("Unexpected value")
			break
		}
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
	val, err := A.Get(8, 5)

	if err != nil {
		t.Error("There was an error getting the value:", err)
	}

	if val != 11 {
		t.Error("Wrong value was retrieved expected 11, but got:", val)
	}
}

func TestMatrixSet(t *testing.T) {
	A := Zeros(5, 3)
	A.Set(4, 3, 10)

	expected_values := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0}

	for i, v := range A.Values() {
		if v != expected_values[i] {
			t.Error("Unexpected value")
			break
		}
	}
}

func TestMatrixAdd(t *testing.T) {
	A := Rand(3, 5)
	B := Rand(3, 5)

	a_values := A.Values()
	b_values := B.Values()

	A.Add(B)

	for i, v := range A.Values() {
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

	A.Sub(B)

	for i, v := range A.Values() {
		if v != (a_values[i] - b_values[i]) {
			t.Error("Wrong value in subtraction")
			break
		}
	}
}

func TestMatrixMul(t *testing.T) {}

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
	vals := FromMatlab("[1 2; 3 4; 5 6]").Transpose().Values()

	expected_vals := []float64{1, 3, 5, 2, 4, 6}

	for i, v := range vals {
		if v != expected_vals[i] {
			t.Error("Unexpected value occurred")
			break
		}
	}
}

func TestMatrixScale(t *testing.T) {

}
func TestMatrixAddNum(t *testing.T) {}
