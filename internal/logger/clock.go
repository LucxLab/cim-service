package logger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type utcClock struct{}

func (c *utcClock) Now() time.Time {
	return time.Now().UTC()
}

func (c *utcClock) NewTicker(duration time.Duration) *time.Ticker {
	return time.NewTicker(duration)
}

func newUtcClock() zapcore.Clock {
	return &utcClock{}
}
