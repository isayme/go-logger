package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogFormat string

var (
	FORMAT_CONSOLE LogFormat = "console"
	FORMAT_JSON    LogFormat = "json"
)

// SetFormat set log format
func SetFormat(format LogFormat) {
	switch format {
	case FORMAT_CONSOLE:
		l = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05.000Z",
		})
	case FORMAT_JSON:
		l = log.With().Caller().Logger().Output(os.Stderr)
	}
}
