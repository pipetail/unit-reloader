module github.com/pipetail/unit-reloader

go 1.19

require (
	github.com/aws/aws-sdk-go v1.44.162
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
)

require (
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.5.0
