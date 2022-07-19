package verify

func isUnitValid(unit []int) bool {
	unitMap := make(map[int]int)
	for i := 0; i < len(unit); i++ {
		if unit[i] == -1 {
			return false
		}

		if _, ok := unitMap[unit[i]]; ok {
			unitMap[unit[i]] = unit[i]
		}
	}
	return true

}

func verifyGrid(grid [][]int) bool {
	return true
}
