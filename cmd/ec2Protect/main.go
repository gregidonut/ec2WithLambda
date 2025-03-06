package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	fmt.Printf("Received event: %+v", event)

	var detail struct {
		InstanceID string `json:"instance-id"`
	}
	if err := json.Unmarshal(event.Detail, &detail); err != nil {
		return fmt.Errorf("failed to parse event detail: %w", err)
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %v", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	if _, err = ec2Client.StartInstances(ctx, &ec2.StartInstancesInput{
		InstanceIds: []string{detail.InstanceID},
	}); err != nil {
		return fmt.Errorf("failed to start instance %s: %v", detail.InstanceID, err)
	}

	fmt.Printf("Protected instance stopped - starting up instance: %s", detail.InstanceID)
	return nil
}

func main() {
	lambda.Start(handler)
}
