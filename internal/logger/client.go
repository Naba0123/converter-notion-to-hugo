package logger

import (
	"errors"

	"github.com/Naba0123/converter-notion-to-hugo/internal/errorwrapper"
	"go.uber.org/zap"
)

type Client struct {
	logger *zap.Logger
}

var (
	ErrLoggerNotInitialized = errors.New("logger has not initialized")
)

// Init initialize logger Client
func Init() (_ *Client, err error) {
	defer errorwrapper.Wrap(&err, "Init()")

	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Client{logger: l}, nil
}

// Logger get *zap.Logger in this client
func (c *Client) Logger() (_ *zap.Logger, err error) {
	defer errorwrapper.Wrap(&err, "Logger()")

	if c.logger == nil {
		return nil, ErrLoggerNotInitialized
	}

	return c.logger, nil
}
