package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Grid [][]int `json:"grid"`
}

type MyResponse struct {
	Message bool `json:""`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: verifyGrid(event.Grid)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
