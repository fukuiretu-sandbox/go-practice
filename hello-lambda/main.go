package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, event Event) (string, error) {
	return fmt.Sprintf("Oh.. Hello %s!", event.Name), nil
}

func main() {
	lambda.Start(handler)
}
