package matrix

import "testing"

func BenchmarkTransposeBig(b *testing.B) {
	A := Rand(100, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		A.Transpose()
	}
}

func BenchmarkMulBig(b *testing.B) {
	A := Rand(100, 100)
	B := Rand(100, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		A.SafeMul(B)
	}
}

func BenchmarkDotBig(b *testing.B) {
	var A, B *Matrix

	for i := 0; i < b.N; i++ {
		A = Rand(100, 100)
		B = Rand(100, 100)
		A.Dot(B)
	}
}

func BenchmarkRandBig(b *testing.B) {
	var A *Matrix
	for i := 0; i < b.N; i++ {
		A = Rand(100, 100)
		A.Rows()
	}
}
