package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

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
	log.Printf("EVENT: %t", verifyBoxes(grid))
	log.Printf("EVENT: %t", verifyRows(grid))
	log.Printf("EVENT: %t", verifyCols(grid))

	return verifyBoxes(grid) && verifyRows(grid) && verifyCols(grid)
}

// type MyEvent struct {
// 	Grid [][]int `json:"grid"`
// }

// type MyResponse struct {
// 	Message bool `json:"message"`
// }

// func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
// 	log.Printf("EVENT: %v", event)
// 	log.Printf("EVENT.GRID: %v", event.Grid)
// 	log.Printf("length of EVENT.GRID: %d", len(event.Grid))
// 	return MyResponse{
// 		Message: VerifyGrid(event.Grid)}, nil
// }

type jsonGrid struct {
	Grid [][]int `json:"grid"`
}

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
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
	result := VerifyGrid(grid.Grid)
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
