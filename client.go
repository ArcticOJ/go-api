package arctic

import (
	"github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
	l Logger
}

type Logger interface {
	Info(format string, args ...interface{})
	Success(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

func NewClient(baseUrl string, logger Logger) (c *Client) {
	if logger == nil {
		logger = &NilLogger{}
	}
	c = &Client{
		c: req.C().SetBaseURL(baseUrl).SetCommonHeader("User-Agent", "iceberg/0.1.0-alpha (https://github.com/ArcticOJ/iceberg)"),
		l: logger,
	}
	return c
}
