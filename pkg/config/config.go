package config

import (
	"flag"
	"fmt"

	"github.com/pipetail/unit-reloader/pkg/util"
)

type Config struct {
	QueueURL         string
	UnitName         string
	UnitAction       string
	QueueMaxMessages int
	QueueWaitTime    int
}

func New() (*Config, error) {
	queueURL := flag.String("queue-url", "", "SQS queue URL")
	unit := flag.String("unit", "", "systemd unit name")
	action := flag.String("action", "", "systemd unit action")

	flag.Parse()

	// perform basic sanity checks
	if *queueURL == "" {
		return nil, fmt.Errorf("queue-url can't be emtpy")
	}

	if *unit == "" {
		return nil, fmt.Errorf("unit can't be emtpy")
	}

	if *action == "" {
		return nil, fmt.Errorf("action can't be emtpy")
	}

	if !util.StringInSlice(*action, []string{"start", "stop", "restart"}) {
		return nil, fmt.Errorf("action can be only start, stop or restart")
	}

	// get the config
	return &Config{
		QueueURL:         *queueURL,
		UnitName:         *unit,
		UnitAction:       *action,
		QueueMaxMessages: 1,
		QueueWaitTime:    20,
	}, nil
}
