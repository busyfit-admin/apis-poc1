/*
This Lambda is triggered by the Example API POC
*/

package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"

	"github.com/aws/aws-lambda-go/lambda"
)

type Service struct {
	ctx    context.Context
	logger *log.Logger
}

func main() {

	ctx, root := xray.BeginSegment(context.TODO(), "example-lambda")
	defer root.Close(nil)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Cannot load config: %v\n", err)
	}

	awsv2.AWSV2Instrumentor(&cfg.APIOptions)

	svc := Service{
		ctx:    ctx,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	lambda.Start(svc.handleAPIRequests)

}

func (svc *Service) handleAPIRequests(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	reqContext := request.RequestContext

	svc.logger.Printf("Request Context data: %v", reqContext)

	return events.APIGatewayProxyResponse{}, nil
}
