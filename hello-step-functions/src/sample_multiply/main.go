package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	AdditionResult struct {
		Var int `json:"var"`
	}
}

type Response struct {
	Baz int `json:"baz"`
}

func handler(ctx context.Context, event Event) (Response, error) {
	result := event.AdditionResult.Var * 3

	return Response{Baz: result}, nil
}

func main() {
	lambda.Start(handler)
}
