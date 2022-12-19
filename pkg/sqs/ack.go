package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Delete(sqsClient *sqs.SQS, queue string, message *sqs.Message) error {
	request := &sqs.DeleteMessageInput{
		QueueUrl:      &queue,
		ReceiptHandle: message.ReceiptHandle,
	}

	_, err := sqsClient.DeleteMessage(request)
	return err
}
