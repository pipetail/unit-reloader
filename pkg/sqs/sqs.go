package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Delete(sqsClient SqsClinet, queue string, message *sqs.Message) error {
	request := &sqs.DeleteMessageInput{
		QueueUrl:      &queue,
		ReceiptHandle: message.ReceiptHandle,
	}

	_, err := sqsClient.DeleteMessage(request)
	return err
}

func Receive(sqsClient SqsClinet, queue string, maxNumberOfMessages int, waitTimeSeconds int) (*sqs.ReceiveMessageOutput, error) {
	return sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queue,
		MaxNumberOfMessages: aws.Int64(int64(maxNumberOfMessages)),
		WaitTimeSeconds:     aws.Int64(int64(waitTimeSeconds)),
	})
}

type SqsClinet interface {
	DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
}
