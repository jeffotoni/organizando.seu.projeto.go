// Go Api server
// @jeffotoni
// 2019-06-01

package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type JobSqsFifo struct {
	Resource string    `json:"resource"`
	UserID   int64     `json:"user_id"`
	Topic    string    `json:"topic"`
	Received time.Time `json:"received"`
	Sent     time.Time `json:"sent"`
}

func handlerApiGateway(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	log.Printf("Body size = %d.\n", len(request.Body))
	log.Printf("Body %s.\n", request.Body)

	log.Println("Headers:")
	for key, value := range request.Headers {
		log.Printf("    %s: %s\n", key, value)
	}

	log.Println("-----------------------------")
	var byteBody io.Reader
	byteBody = strings.NewReader(request.Body)

	var jobFifo JobSqsFifo
	err := json.NewDecoder(byteBody).Decode(&jobFifo)
	if err != nil {
		log.Println("Error NewDecoder:: ", err)
		return events.APIGatewayProxyResponse{Body: `{"message":"Error Decode json!"}`, StatusCode: 400}, nil
	}

	log.Println(jobFifo.Resource)
	log.Println(jobFifo.UserID)

	ok, erros := sqs.SendJobFifo(request.Body)
	if !ok {
		log.Println("Error SendJobFifo:: ", erros)
		return events.APIGatewayProxyResponse{Body: `{"message":"error SendJobFifo!"}`, StatusCode: 400}, nil
	}

	// remove nofificacao do meli..
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {

	lambda.Start(handlerApiGateway)
}
