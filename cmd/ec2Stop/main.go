package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func handleRequest(ctx context.Context) error {
	instances := strings.Split(os.Getenv("EC2_INSTANCES"), ",")

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	_, err = ec2Client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: instances,
	})
	if err != nil {
		return fmt.Errorf("failed to stop instances: %w", err)
	}

	fmt.Printf("Stopped instances: %v\n", instances)
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
