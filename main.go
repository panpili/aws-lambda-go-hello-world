package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse,
	error) {
	fmt.Printf("method:%v\n", request.RequestContext.HTTP.Method)
	fmt.Printf("path:%v\n", request.RawPath)
	fmt.Printf("query:%v\n", request.QueryStringParameters)
	//fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("body = %v\n", request.Body)

	fmt.Println("headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	return events.APIGatewayV2HTTPResponse{Body: "hello world!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
