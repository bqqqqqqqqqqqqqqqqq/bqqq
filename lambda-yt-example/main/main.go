package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"How old are you"`
}

type myResponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event myEvent) (myResponse, error) {
	return myResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
