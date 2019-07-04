package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Foo int
}

func handler() (Response, error) {
	res := Response{Foo: 1}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
