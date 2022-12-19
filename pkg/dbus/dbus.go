package dbus

import (
	"context"
	"fmt"

	"github.com/coreos/go-systemd/dbus"
)

func Do(ctx context.Context, dbusClient *dbus.Conn, unit string, action string) error {
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

func Restart(ctx context.Context, dbusClient *dbus.Conn, unit string) error {
	_, err := dbusClient.RestartUnitContext(ctx, unit, "replace", nil)
	return err
}

func Start(ctx context.Context, dbusClient *dbus.Conn, unit string) error {
	_, err := dbusClient.StartUnitContext(ctx, unit, "replace", nil)
	return err
}

func Stop(ctx context.Context, dbusClient *dbus.Conn, unit string) error {
	_, err := dbusClient.StopUnitContext(ctx, unit, "replace", nil)
	return err
}
