package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/likexian/whois"
)


func main (){
	lambda.Start(handler)
}

// whois query for an ipv6
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		result, err := whois.Whois("2603:8000:1a01:3c84:f54e:4230:2378:1457")
		ApiResponse := events.APIGatewayProxyResponse{Body: result,	StatusCode: 200}
		return ApiResponse, err
}