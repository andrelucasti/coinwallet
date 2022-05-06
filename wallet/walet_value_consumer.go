package wallet

import (
	"context"
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

var i int

func Consumer() []WalletValueOutputMessage {
	var outputMessages []WalletValueOutputMessage

	for true {
		i++
		qName := aws.String("wallet_value")
		waitTime := aws.Int(10)
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
				return nil
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

				outputMessage := WalletValueOutputMessage{
					MessageId: *msg.MessageId,
					Body:      *msg.Body,
				}

				outputMessages = append(outputMessages, outputMessage)
			}
		} else {
			log.Println("Got a error loading default configs: ", err)
		}

	}

	return outputMessages
}

type WalletValueOutputMessage struct {
	MessageId string
	Body      string
}
