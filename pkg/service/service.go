package service

import (
	component "github.com/andreylm/nats-component"
)

// Version - server version
var Version = "v1"

// Service - nats service
type Service struct {
	Component component.Component
}
