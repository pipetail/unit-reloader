package reloader

import (
	"context"
	"log"

	"github.com/coreos/go-systemd/dbus"
	"github.com/pipetail/unit-reloader/pkg/config"
	dbusImpl "github.com/pipetail/unit-reloader/pkg/dbus"
	sqsImpl "github.com/pipetail/unit-reloader/pkg/sqs"
)

func Run(ctx context.Context, cfg *config.Config, sqsClient sqsImpl.SqsClinet, dbusClient *dbus.Conn) {

	// consume messages in loop
	for {
		output, err := sqsImpl.Receive(sqsClient, cfg.QueueURL, cfg.QueueMaxMessages, cfg.QueueWaitTime)
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
