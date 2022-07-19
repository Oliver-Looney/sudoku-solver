package verify_test

import (
	"testing"
)

func TestVerifyGrid(t *testing.T) {
	t.Parallel()
	var want bool = true
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	got := verify.verifyGrid(twoD)
	if want != got {
		t.Errorf("Want %f, got %f", want, got)
	}
}
