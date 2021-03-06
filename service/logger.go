package service

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/piotrkowalczuk/auth-service/lib"
)

const (
	// LoggerAdapterConsole ...
	LoggerAdapterConsole = "console"
)

var (
	// Logger is singleton instance of logger.
	Logger *logrus.Logger
	// ErrLoggerAdapterNotSupported ...
	ErrLoggerAdapterNotSupported = errors.New("service: logger adapter not supported")
)

// LoggerConfig ...
type LoggerConfig struct {
	Adapter string `xml:"adapter"`
	Level   uint8  `xml:"level"`
}

// InitLogger ...
func InitLogger(config LoggerConfig) {
	Logger = logrus.New()
	Logger.Level = logrus.Level(config.Level)

	switch config.Adapter {
	case LoggerAdapterConsole:
		Logger.Formatter = &lib.ConsoleFormatter{}
	default:
		Logger.Fatal(ErrLoggerAdapterNotSupported)
	}
}
