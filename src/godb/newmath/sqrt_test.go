package newmath

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 9, 3
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
        }
}