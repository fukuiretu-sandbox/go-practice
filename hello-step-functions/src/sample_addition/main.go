package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	SettingResult struct {
		Foo int `json:"foo"`
	}
}

type Response struct {
	Var int
}

func handler(ctx context.Context, event Event) (Response, error) {
	result := event.SettingResult.Foo + 2

	res := Response{Var: result}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
