package dbus

import (
	"context"
	"fmt"
)

func Do(ctx context.Context, dbusClient DbusClient, unit string, action string) error {
	switch action {
	case "start":
		return Start(ctx, dbusClient, unit)
	case "stop":
		return Stop(ctx, dbusClient, unit)
	case "restart":
		return Restart(ctx, dbusClient, unit)
	default:
		return fmt.Errorf("action %s is not defined", action)
	}
}

func Restart(ctx context.Context, dbusClient DbusClient, unit string) error {
	_, err := dbusClient.RestartUnitContext(ctx, unit, "replace", nil)
	return err
}

func Start(ctx context.Context, dbusClient DbusClient, unit string) error {
	_, err := dbusClient.StartUnitContext(ctx, unit, "replace", nil)
	return err
}

func Stop(ctx context.Context, dbusClient DbusClient, unit string) error {
	_, err := dbusClient.StopUnitContext(ctx, unit, "replace", nil)
	return err
}

type DbusClient interface {
	StopUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error)
	StartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error)
	RestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error)
}
