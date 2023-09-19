package config

import "time"

type App struct {
	ServiceName    string        `envconfig:"SERVICE_NAME" default:"JWT_Login"`
	Mode           string        `envconfig:"APP_MODE" default:"development"`
	ContextTimeout time.Duration `envconfig:"CONTEXT_TIMEOUT" default:"2s"`
}
