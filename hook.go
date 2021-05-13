package logger

import "github.com/rs/zerolog"

type LevelNameHook struct{}

func (h LevelNameHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	if l != zerolog.NoLevel {
		e.Str("level_name", l.String())
	} else {
		e.Str("level_name", "NoLevel")
	}
}
