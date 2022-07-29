package main

import (
	"testing"
)

func TestIsSafeRowValid1(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeRow(grid, 7, 1)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeRowValid2(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, 9, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeRow(grid, 2, 1)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeRowNotValid(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = false
	got := isSafeRow(grid, 8, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeColValid1(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, -1, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeCol(grid, 7, 5)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeColValid2(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, 9, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeCol(grid, 7, 1)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeBoxValid1(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, -1, 4, 7, 8, 3, 2, 6, 5},
		{3, -1, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeBox(grid, 1, 1, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeBoxValid2(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, 9, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	got := isSafeBox(grid, 7, 2, 1)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeBoxNotValid(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, -1, 4, 7, 8, 3, 2, 6, 5},
		{3, -1, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = false
	got := isSafeBox(grid, 1, 1, 2)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeColNotValid(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 6, 5},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, -1},
		{5, 7, 9, 8, 4, 9, 6, 4, 2}}
	t.Parallel()
	var want bool = false
	got := isSafeCol(grid, 8, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeValid1(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, -1, 4, 7, 8, 3, 2, -1, -1},
		{3, -1, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, -1}}
	t.Parallel()
	var want bool = true
	got := isSafe(grid, 1, 1, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeValid2(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, 9, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, 7},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	got := isSafe(grid, 2, 7, 1)
	row := isSafeRow(grid, 2, 1)
	col := isSafeCol(grid, 7, 1)
	box := isSafeBox(grid, 2, 7, 1)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
		t.Errorf("row %t", row)
		t.Errorf("col %t", col)
		t.Errorf("box %t", box)
		t.Errorf("and : %t", row && col && box)
	}
}

func TestIsSafeNotValid1(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, -1, 4, 7, 8, 3, 2, -1, 6},
		{3, -1, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, -1}}
	t.Parallel()
	var want bool = false
	got := isSafe(grid, 1, 1, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestIsSafeNotValid2(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, -1, 4, 7, 8, 3, 2, -1, -1},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, -1}}
	t.Parallel()
	var want bool = false
	got := isSafe(grid, 1, 1, 6)
	if want != got {
		t.Errorf("Want %t, got %t", want, got)
	}
}

func TestSolveGrid1(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, 9, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	input := myResult{
		Grid: grid,
	}
	got := solveGrid(input)
	got.Solvable = (len(got.Solutions) != 0)
	if len(got.Solutions) == 0 {
		t.Errorf("Want %t, got %t", want, got.Solvable)
		t.Errorf("%d", got.Solutions)
	}
	// t.Errorf("%d", len(got.Solutions))
	// for i := 0; i < len(got.Solutions); i++ {
	// 	t.Errorf("%d", got.Solutions[i])
	// }
}

func TestSolveGrid(t *testing.T) {
	grid := [][]int{
		{7, -1, 5, 6, 1, 2, 3, -1, 4},
		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
		{3, -1, -1, 4, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
		{2, 5, -1, -1, 3, -1, 1, -1, -1},
		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
		{5, -1, 9, -1, 4, -1, -1, -1, 2}}
	t.Parallel()
	var want bool = true
	input := myResult{
		Grid: grid,
	}
	got := solveGrid(input)
	got.Solvable = (len(got.Solutions) != 0)
	if want != got.Solvable {
		t.Errorf("Want %t, got %t", want, got.Solvable)
		t.Errorf("%v", got.Solutions)
	}
	// t.Errorf("%d", len(got.Solutions))
	// for i := 0; i < len(got.Solutions); i++ {
	// 	t.Errorf("%d", got.Solutions[i])
	// }
}

func TestGetNextBlank(t *testing.T) {
	grid := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 1, 1},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, 1}}
	t.Parallel()
	want1, want2 := -1, -1
	got1, got2 := getNextBlank(grid)
	if want1 != got1 {
		t.Errorf("Want %d, got %d", want1, got1)
	}
	if want2 != got2 {
		t.Errorf("Want %d, got %d", want2, got2)
	}
}
