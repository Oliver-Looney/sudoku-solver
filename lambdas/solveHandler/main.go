package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

func isSafeRow(grid [][]int, row int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}
	return true
}

func isSafeBox(grid [][]int, row int, col int, num int) bool {
	startRow := row - row%3
	startCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func isSafeCol(grid [][]int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}
	return true
}

func isSafe(grid [][]int, i int, j int, num int) bool {
	rowresult := isSafeRow(grid, i, num)
	colresult := isSafeCol(grid, j, num)
	box := isSafeBox(grid, i, j, num)
	return box && rowresult && colresult
}

func checkGridIsFilledIn(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == -1 {
				return false
			}
		}
	}
	return true
}

type possibleSolution struct {
	Grid     [][]int
	position int
}

type Stack []possibleSolution

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(input possibleSolution) {
	*s = append(*s, input) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (possibleSolution, bool) {
	if s.IsEmpty() {
		return possibleSolution{}, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func revSolve(grid [][]int, result myResult) myResult {
	// i := position / 9
	// j := position - (position/9)*9
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == -1 {
				for n := 1; n < 10; n++ {
					if isSafe(grid, y, x, n) {
						grid[y][x] = n
						result = revSolve(grid, result)
						grid[y][x] = -1
					}
				}
				return result
			}
		}
	}
	fmt.Printf("recurisve og: \n%v\n", grid)
	result.Solutions = append(result.Solutions, grid)
	result.Solvable = true
	return result
}

func solveGrid(result myResult) myResult {
	return revSolve(result.Grid, myResult{Grid: result.Grid})
}

// func itSolve(grid [][]int) {
// 	var stack Stack
// 	stack.Push(possibleSolution{Grid: grid, position: 0})
// 	var flag bool
// 	count := 0
// 	for len(stack) > 0 {
// 		topOfStack, _ := stack.Pop()
// 		topOfStack.position = count
// 		fmt.Printf("\ntop of stack:%v", topOfStack.Grid)
// 		fmt.Printf("\nstack:\n")
// 		for i := 0; i < len(stack); i++ {
// 			fmt.Printf("%v\n", stack[i].Grid)
// 		}
// 		fmt.Printf("length of stack: %d\n\n", len(stack))
// 		flag = true
// 		for y := 0; y < 9; y++ {
// 			for x := 0; x < 9; x++ {
// 				if topOfStack.Grid[y][x] == -1 {
// 					for n := 1; n < 10; n++ {
// 						// if len(stack) == 2 {
// 						// 	fmt.Printf("\ny: %d, x: %d\n", y, x)
// 						// 	fmt.Printf("topOfStack.Grid[y][x]==-1 = %t", topOfStack.Grid[y][x] == -1)
// 						// 	fmt.Printf("\nn:%d\n", n)
// 						// 	fmt.Printf("working with: %v\n", topOfStack.Grid)
// 						// 	fmt.Printf("%t", isSafe(topOfStack.Grid, y, x, n))
// 						// }
// 						if isSafe(topOfStack.Grid, y, x, n) {
// 							newStackMember := topOfStack
// 							newStackMember.Grid[y][x] = n
// 							stack.Push(newStackMember)
// 							// fmt.Printf("\n%d\n", n)
// 							// fmt.Printf("pushed :\n%v\n", newStackMember.Grid)
// 						}
// 					}
// 					flag = false
// 				}
// 			}
// 		}
// 		// fmt.Printf("iterative :\n%v\n", topOfStack.Grid)
// 		// fmt.Printf("\n%d\n", topOfStack.position)
// 	}
// }

// func solveGrid(result myResult) myResult {
// 	revSolve(result.Grid)
// 	itSolve(result.Grid)
// 	var stack Stack
// 	stack.Push(possibleSolution{Grid: result.Grid, position: 0})
// 	fmt.Printf("%t", isSafe([][]int{
// 		{7, -1, 5, 6, 1, 2, 3, 9, 4},
// 		{-1, -1, 4, 7, -1, 3, 2, 6, 5},
// 		{3, -1, -1, 4, -1, -1, -1, -1, -1},
// 		{-1, -1, 1, -1, 7, 8, 5, -1, -1},
// 		{-1, -1, 7, -1, 5, 6, 9, 2, 8},
// 		{2, 5, -1, -1, 3, -1, 1, -1, -1},
// 		{-1, 2, 3, -1, -1, 7, 4, 8, -1},
// 		{-1, -1, 6, 3, 2, 9, 7, -1, 1},
// 		{5, -1, 9, -1, 4, -1, -1, -1, 2}}, 2, 7, 1))

// 	for !stack.IsEmpty() {
// 		// fmt.Printf("\n%d\n", len(stack))
// 		// fmt.Printf("\n\n")
// 		// for x := 0; x < len(stack); x++ {
// 		// 	fmt.Printf("\n%v\n", stack[x])
// 		// }
// 		// fmt.Printf("\n\n")
// 		topOfStack, _ := stack.Pop()
// 		// fmt.Printf("\n%v\n", topOfStack)
// 		i := topOfStack.position / 9
// 		j := topOfStack.position - (topOfStack.position/9)*9
// 		flag := topOfStack.Grid[i][j] != -1
// 		if topOfStack.Grid[2][7] == 1 {
// 			fmt.Printf("%v", topOfStack.Grid)
// 		}
// 		for i < 9 && j < 9 && flag {
// 			// fmt.Printf("%d\n", topOfStack.Grid[i][j])
// 			// fmt.Printf("%d\n", topOfStack.position)
// 			topOfStack.position = topOfStack.position + 1
// 			i = topOfStack.position % 9
// 			j = (topOfStack.position - i) / 9
// 			flag = (topOfStack.Grid[i][j] != -1)
// 		}

// 		// topOfStack.position = topOfStack.position + 1
// 		// fmt.Printf("%d", topOfStack.Grid[i][j])
// 		if topOfStack.position == 81 {
// 			// full grid tick
// 			result.Solutions = append(result.Solutions, topOfStack.Grid)
// 		}
// 		if topOfStack.Grid[i][j] == -1 {
// 			// try diff nums and add to stack
// 			for n := 1; n < 10; n++ {
// 				if topOfStack.position == 26 {
// 					fmt.Printf("%t", isSafe(topOfStack.Grid, i, j, n))
// 				}
// 				if isSafe(topOfStack.Grid, i, j, n) {
// 					newStackMember := topOfStack
// 					copier.Copy(newStackMember, topOfStack)
// 					newStackMember.Grid[i][j] = n
// 					stack.Push(newStackMember)
// 				}
// 			}
// 			//increment position
// 			topOfStack.position = topOfStack.position + 1
// 		}
// 		if topOfStack.position == 81 {
// 			// full grid tick
// 			result.Solutions = append(result.Solutions, topOfStack.Grid)
// 		}
// 		// for index := 0; index < 81; index++ {
// 		// 	i := index % 9
// 		// 	j := (index - i) / 9

// 		// 	if topOfStack.Grid[i][j] == -1 {
// 		// 		for n := 1; n < 10; n++ {
// 		// 			if isSafe(topOfStack.Grid, i, j, n) {
// 		// 				newStackMember := topOfStack
// 		// 				copier.Copy(newStackMember, topOfStack)
// 		// 				newStackMember.Grid[i][j] = n
// 		// 				newStackMember.position = i + (9 * j)
// 		// 				stack.Push(newStackMember)
// 		// 			}
// 		// 		}
// 		// 	}
// 		// }

// 		// for i := topOfStack.position - 9*(topOfStack.position%9); i < 9; i++ {
// 		// 	for j := topOfStack.position % 9; j < 9; j++ {
// 		// 		if topOfStack.Grid[i][j] == -1 {
// 		// 			for n := 1; n < 10; n++ {
// 		// 				if isSafe(topOfStack.Grid, i, j, n) {
// 		// 					newStackMember := topOfStack
// 		// 					copier.Copy(newStackMember, topOfStack)
// 		// 					newStackMember.Grid[i][j] = n
// 		// 					newStackMember.position = (i + (9 * j))
// 		// 					if newStackMember.position == 80 {
// 		// 						result.Solutions = append(result.Solutions, newStackMember.Grid)
// 		// 					} else {
// 		// 						stack.Push(newStackMember)
// 		// 					}
// 		// 					// stack.Push(newStackMember)
// 		// 				}
// 		// 			}
// 		// 		}
// 		// 	}
// 		// }
// 	}
// 	return result
// }

type jsonGrid struct {
	Grid [][]int `json:"grid"`
}

type myResult struct {
	Grid      [][]int
	Solutions [][][]int
	Solvable  bool
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	err := fastjson.Validate(request.Body)
	log.Printf("EVENT: %v", request.Body)
	var grid jsonGrid
	Data := []byte(request.Body)
	errJSON := json.Unmarshal(Data, &grid)
	if errJSON != nil {
		fmt.Println(errJSON)
	}
	solveGridInput := myResult{
		Grid:     grid.Grid,
		Solvable: false,
	}
	result := solveGrid(solveGridInput)
	if err != nil {
		body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + request.Body
		ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
	} else {
		ApiResponse = events.APIGatewayProxyResponse{Body: fmt.Sprint(result), StatusCode: 200}
	}
	// Response
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
