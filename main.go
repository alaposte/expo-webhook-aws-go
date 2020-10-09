package main

import (
	"context"
	"fmt"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ExpoWebhookPayload struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	ArtifactUrl string `json:"artifactUrl"`
	Platform    string `json:"platform"`
}

type ApiGatewayMap struct {
	// Body      ExpoWebhookPayload `json:"body"`
	Signature string `json:"signature"`
	// Headers   map[string][]string `json:"headers"`
}

func HandleRequest(ctx context.Context, payload ApiGatewayMap) (events.APIGatewayProxyResponse, error) {
	// fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	// fmt.Printf("Body size = %d.\n", len(payload.Body))
	fmt.Printf("Body content: %+v\n", payload.Body)

	fmt.Println("Headers:")
	for key, value := range payload.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
