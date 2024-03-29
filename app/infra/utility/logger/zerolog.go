package logx

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func Create() *zerolog.Logger {

	init := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &init
}
