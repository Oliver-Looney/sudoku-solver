package main

import (
	"testing"
)

func TestIsUnitValidValidTest(t *testing.T) {
	t.Parallel()
	var want bool = true
	got := isUnitValid([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsUnitValidNotValidTest1(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{1, 1, 3, 4, 5, 6, 7, 8, 9})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsUnitValidNotValidTest2(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{1, 1, 6, 4, 5, 6, 7, 8, 5})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}
func TestIsUnitValidNotValidTest3(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{3, 3, 3, 3, 3, 3, 3, 3, 3})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}
func TestIsUnitValidNotValidTest4(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{1, 2, 3, 4, 5, 6, 9, 8, 9})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}
func TestIsUnitValidNotFullTest1(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{-1, 2, 3, 4, 5, 6, 7, 8, 9})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}
func TestIsUnitValidNotFullTest2(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{1, -1, 3, 4, 5, -1, 9, 8, 9})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}
func TestIsUnitValidNotFullTest3(t *testing.T) {
	t.Parallel()
	var want bool = false
	got := isUnitValid([]int{1, -1, -1, -1, -1, -1, -1, -1, -1})
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

// verify rows
func TestVerifyRowsValid1(t *testing.T) {
	validSudoku := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	t.Parallel()
	var want bool = true
	got := verifyRows(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestVerifyRowsNotValid1(t *testing.T) {
	validSudoku := [][]int{
		{1, 2, 1, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 2, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 5, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 9, 3, 4, 5, 6, 7, 8, 9}}
	t.Parallel()
	var want bool = false
	got := verifyRows(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

// verify cols to array
func TestColToArray(t *testing.T) {
	validSudoku := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9}}
	t.Parallel()
	var want []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var got []int = colToArray(validSudoku, 0)
	for i := 0; i < 9; i++ {
		if want[i] != got[i] {
			t.Errorf("Want %v, got %v", want, got)
		}
	}
}

// verify cols
func TestVerifyColsValid1(t *testing.T) {
	validSudoku := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9}}
	t.Parallel()
	var want bool = true
	got := verifyCols(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestVerifyColsNotValid1(t *testing.T) {
	validSudoku := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 3, 3, 3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 9, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 2, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 1}}
	t.Parallel()
	var want bool = false
	got := verifyCols(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

// Verify grid to boxes
func TestBoxToArray(t *testing.T) {
	validSudoku := [][]int{
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{4, 5, 6, 4, 5, 6, 4, 5, 6},
		{7, 8, 9, 7, 8, 9, 7, 8, 9},
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{4, 5, 6, 4, 5, 6, 4, 5, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{8, 8, 8, 8, 8, 8, 8, 8, 8},
		{7, 8, 9, 7, 8, 9, 7, 8, 9}}
	t.Parallel()
	var want []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var got []int = boxToArray(validSudoku, 0, 0)
	for i := 0; i < 9; i++ {
		if want[i] != got[i] {
			t.Errorf("Want %d, got %d", want[i], got[i])
		}
	}
}

func TestVerifyBoxesValid1(t *testing.T) {
	validSudoku := [][]int{
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{4, 5, 6, 4, 5, 6, 4, 5, 6},
		{7, 8, 9, 7, 8, 9, 7, 8, 9},
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{4, 5, 6, 4, 5, 6, 4, 5, 6},
		{7, 8, 9, 7, 8, 9, 7, 8, 9},
		{1, 2, 3, 1, 2, 3, 1, 2, 3},
		{4, 5, 6, 4, 5, 6, 4, 5, 6},
		{7, 8, 9, 7, 8, 9, 7, 8, 9}}
	t.Parallel()
	var want []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, elementI := range []int{0, 3, 6} {
		for _, elementJ := range []int{0, 3, 6} {
			got := boxToArray(validSudoku, elementI, elementJ)
			for i := 0; i < 9; i++ {
				if want[i] != got[i] {
					t.Errorf("got %v", got)
					t.Errorf("want %v", want)
					t.Errorf("elementI %d, elementJ %d", elementI, elementJ)
					t.Errorf("Want %d, got %d", want[i], got[i])
				}
			}
		}
	}
}

func TestVerifyGridValid1(t *testing.T) {
	validSudoku := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 1, 6, 3, 2}}
	t.Parallel()
	var want bool = true
	got := VerifyGrid(validSudoku)
	rows := verifyRows(validSudoku)
	cols := verifyCols(validSudoku)
	boxes := verifyBoxes(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
		t.Errorf("rows %t", rows)
		for i := 0; i < 9; i++ {
			t.Errorf("row %t", isUnitValid(validSudoku[i]))
		}
		t.Errorf("cols %t", cols)
		t.Errorf("boxes %t", boxes)
	}
}

func TestVerifyGridNotValid1(t *testing.T) {
	validSudoku := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = false
	got := VerifyGrid(validSudoku)
	rows := verifyRows(validSudoku)
	cols := verifyCols(validSudoku)
	boxes := verifyBoxes(validSudoku)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
		t.Errorf("rows %t", rows)
		t.Errorf("cols %t", cols)
		t.Errorf("boxes %t", boxes)
	}
}

// func TestHandleLambdaEvent(t *testing.T) {
// 	grids := [][]int{
// 		{7, 8, 5, 6, 1, 2, 3, 9, 4},
// 		{9, 1, 4, 7, 8, 3, 2, 6, 5},
// 		{3, 6, 2, 4, 9, 5, 8, 1, 7},
// 		{6, 9, 1, 2, 7, 8, 5, 4, 3},
// 		{4, 3, 7, 1, 5, 6, 9, 2, 8},
// 		{2, 5, 8, 9, 3, 4, 1, 7, 6},
// 		{1, 2, 3, 5, 6, 7, 4, 8, 9},
// 		{8, 4, 6, 3, 2, 9, 7, 5, 1},
// 		{5, 7, 9, 8, 4, 1, 6, 3, 2}}
// 	event := MyEvent{
// 		Grid: grids,
// 	}
// 	t.Parallel()
// 	want := MyResponse{
// 		Message: true,
// 	}
// 	got, _ := HandleLambdaEvent(event)
// 	if want != got {
// 		t.Errorf("Want %t, got %t", want.Message, got.Message)
// 	}
// }
