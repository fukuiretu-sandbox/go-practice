package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Oh.. Hello %s!", name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
