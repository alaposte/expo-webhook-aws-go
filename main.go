package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Println("Before lambda Start calls the handler function")
	lambda.Start(Handler)
}

func Handler() {
	fmt.Println("my function has been invoked")
}
