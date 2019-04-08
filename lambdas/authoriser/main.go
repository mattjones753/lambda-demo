package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

func Authorise(e events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	if token := e.AuthorizationToken; token != "secret" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}
	statement := events.IAMPolicyStatement{
		Action:   []string{"execute-api:Invoke"},
		Effect:   "Allow",
		Resource: []string{"arn:aws:execute-api:*"},
	}
	policyDoc := events.APIGatewayCustomAuthorizerPolicy{Version: "2012-10-17", Statement: []events.IAMPolicyStatement{statement}}
	return events.APIGatewayCustomAuthorizerResponse{PrincipalID: "principal", PolicyDocument: policyDoc}, nil
}

func main() {
	lambda.Start(Authorise)
}
