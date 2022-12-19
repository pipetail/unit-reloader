package sqs

import (
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

type SqsClinet interface {
	DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
}
