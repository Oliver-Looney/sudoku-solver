package main

func isUnitValid(unit []int) bool {
	unitMap := make(map[int]int)
	for i := 0; i < 9; i++ {
		if unit[i] == -1 {
			return false
		}
		if _, ok := unitMap[unit[i]]; ok {
			return false
		}
		unitMap[unit[i]] = unit[i]
	}
	return true
}

func verifyRows(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		if !isUnitValid(grid[i]) {
			return false
		}
	}
	return true
}

func colToArray(grid [][]int, row int) []int {
	var result []int
	for i := 0; i < 9; i++ {
		// s = append(s, 1)
		result = append(result, grid[i][0])
		// result[i] = grid[row][i]
	}
	return result
}

func verifyCols(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		if !isUnitValid(colToArray(grid, i)) {
			return false
		}
	}
	return true
}

func boxToArray(grid [][]int, elementI int, elementJ int) []int {
	var result []int
	for i := elementI; i < elementI+3; i++ {
		for j := elementJ; j < elementJ+3; j++ {
			result = append(result, grid[i][j])
		}
	}
	return result
}

func verifyBoxes(grid [][]int) bool {
	for _, elementI := range []int{0, 3, 6} {
		for _, elementJ := range []int{0, 3, 6} {
			box := boxToArray(grid, elementI, elementJ)
			if !isUnitValid(box) {
				return false
			}
		}
	}
	return true
}

func VerifyGrid(grid [][]int) bool {
	return verifyBoxes(grid) && verifyRows(grid) && verifyCols(grid)
}
