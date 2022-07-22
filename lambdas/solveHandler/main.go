package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

func solveGrid(grid [][]int) [][]int {
	return grid
}

type jsonGrid struct {
	Grid [][]int `json:"grid"`
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
	result := solveGrid(grid.Grid)
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
