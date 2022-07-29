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

func getNextBlank(grid [][]int) (int, int) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == -1 {
				return y, x
			}
		}
	}
	return -1, -1
}

func solveGrid(input myResult) myResult {
	var stack Stack
	stack.Push(possibleSolution{Grid: input.Grid})
	var result myResult
	for len(stack) > 0 {
		top, _ := stack.Pop()
		y, x := getNextBlank(top.Grid)
		if y == -1 {
			result.Solvable = true
			result.Solutions = append(result.Solutions, top.Grid)
		} else {
			for n := 1; n < 10; n++ {
				if isSafe(top.Grid, y, x, n) {
					newgrid := copyGrid(top.Grid)
					newgrid[y][x] = n
					stack.Push(possibleSolution{Grid: newgrid})
				}
			}
		}
	}
	return result
}

func copyGrid(original [][]int) [][]int {
	result := [][]int{
		{7, 8, 5, 6, 1, 2, 3, 9, 4},
		{9, 1, 4, 7, 8, 3, 2, 1, 1},
		{3, 6, 2, 4, 9, 5, 8, 1, 7},
		{6, 9, 1, 2, 7, 8, 5, 4, 3},
		{4, 3, 7, 1, 5, 6, 9, 2, 8},
		{2, 5, 8, 9, 3, 4, 1, 7, 6},
		{1, 2, 3, 5, 6, 7, 4, 8, 9},
		{8, 4, 6, 3, 2, 9, 7, 5, 1},
		{5, 7, 9, 8, 4, 9, 6, 4, 1}}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			result[i][j] = original[i][j]
		}
	}
	return result
}

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
		solutions, _ := json.Marshal(result.Solutions)
		body := "{\"Grid\": " + fmt.Sprint(result.Grid) + ",\"Solutions\": " + fmt.Sprint(string(solutions)) + ",\"Solvable\": " + fmt.Sprint(result.Solvable) + "}"
		ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 200}
	}
	// Response
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
