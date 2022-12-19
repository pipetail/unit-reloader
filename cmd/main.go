package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/coreos/go-systemd/dbus"
	"github.com/pipetail/unit-reloader/internal/reloader"
	"github.com/pipetail/unit-reloader/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("could not load configuration: %s", err)
	}

	// initialize dbus client
	dbusClient, err := dbus.New()
	if err != nil {
		log.Fatalf("could not initialize dbus client: %s", err)
	}

	// initialize SQS client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sqsClient := sqs.New(sess)

	// start the main loop
	ctx := context.Background()
	reloader.Run(ctx, cfg, sqsClient, dbusClient)
}
