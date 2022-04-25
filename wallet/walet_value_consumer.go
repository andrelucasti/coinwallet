package wallet

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
)

type SQSGetLPMsgAPI interface {
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optsFns ...func(options *sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optsFns ...func(options *sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

func GetQueueURL(ctx context.Context, api SQSGetLPMsgAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(ctx, input)
}

func GetLPMessages(ctx context.Context, api SQSGetLPMsgAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(ctx, input)
}

type ValueConsumer struct {
}

func Consumer() {
	qName := flag.String("q", "wallet_value", "The name of the queue")
	waitTime := flag.Int("w", 10, "How long the queue waits for messages")
	awsRegion := "us-east-1"
	awsEndpoint := "http://localhost:4566"

	//TODO Improve
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {

		return aws.Endpoint{
			URL:           awsEndpoint,
			SigningRegion: region,
			PartitionID:   "aws",
		}, nil
	})

	if defaultConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver)); err == nil {
		client := sqs.NewFromConfig(defaultConfig)
		qInput := &sqs.GetQueueUrlInput{
			QueueName: qName,
		}

		url, err := GetQueueURL(context.TODO(), client, qInput)
		if err != nil {
			log.Println("Got a error getting the queueURL: ", err)
			return
		}

		queueUrl := url.QueueUrl

		//TODO Need Understand more this parameters
		mInput := &sqs.ReceiveMessageInput{
			QueueUrl:              queueUrl,
			AttributeNames:        []types.QueueAttributeName{"SentTimestamp"},
			MaxNumberOfMessages:   10,
			MessageAttributeNames: []string{"All"},
			WaitTimeSeconds:       int32(*waitTime),
		}

		messages, err := GetLPMessages(context.TODO(), client, mInput)

		if err != nil {
			log.Println("Got a error receiving messages: ", err)
		}

		for _, msg := range messages.Messages {
			fmt.Println("msg id: " + *msg.MessageId)
			fmt.Println("msg body: " + *msg.Body)
		}
	} else {
		log.Println("Got a error loading default configs: ", err)
	}
}
