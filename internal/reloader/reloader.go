package reloader

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/coreos/go-systemd/dbus"
	"github.com/pipetail/unit-reloader/pkg/config"
	dbusImpl "github.com/pipetail/unit-reloader/pkg/dbus"
	sqsImpl "github.com/pipetail/unit-reloader/pkg/sqs"
)

func Run(ctx context.Context, cfg *config.Config, sqsClient *sqs.SQS, dbusClient *dbus.Conn) {

	// consume messages in loop
	for {

		output, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            &cfg.QueueURL,
			MaxNumberOfMessages: aws.Int64(int64(cfg.QueueMaxMessages)),
			WaitTimeSeconds:     aws.Int64(int64(cfg.QueueWaitTime)),
		})

		if err != nil {
			log.Printf("could not receive SQS messages: %s", err)
		}

		for _, message := range output.Messages {
			log.Printf("processing SQS message %s: %s", *message.MessageId, *message.Body)
			// TODO: filter messages

			// perform systemd action
			err = dbusImpl.Do(ctx, dbusClient, cfg.UnitName, cfg.UnitAction)
			if err != nil {
				log.Printf("could not perform dbus action: %s", err)
			}

			// delete message
			err = sqsImpl.Delete(sqsClient, cfg.QueueURL, message)
			if err != nil {
				log.Printf("could not delete message: %s", err)
			}
		}

	}
}
