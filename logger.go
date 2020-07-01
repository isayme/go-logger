package logger

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z"
	zerolog.CallerSkipFrameCount++

	log.Logger = log.With().Caller().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// SetLevel set log level
func SetLevel(level string) error {
	level = strings.ToLower(level)
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(lvl)
	return nil
}

// Trace trace log
func Trace(args ...interface{}) {
	log.Trace().Msg(fmt.Sprint(args...))
}

// Tracef trace log with format
func Tracef(format string, args ...interface{}) {
	log.Trace().Msgf(format, args...)
}

// Tracew trace log with additional context
func Tracew(msg string, args ...interface{}) {
	log.Trace().Fields(sweetenFields(args)).Msg(msg)
}

// Debug debug log
func Debug(args ...interface{}) {
	log.Debug().Msg(fmt.Sprint(args...))
}

// Debugf debug log with format
func Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

// Debugw debug log with additional context
func Debugw(msg string, args ...interface{}) {
	log.Debug().Fields(sweetenFields(args)).Msg(msg)
}

// Info info log
func Info(args ...interface{}) {
	log.Info().Msg(fmt.Sprint(args...))
}

// Infof info log with format
func Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

// Infow info log with additional context
func Infow(msg string, args ...interface{}) {
	log.Info().Fields(sweetenFields(args)).Msg(msg)
}

// Warn warning log
func Warn(args ...interface{}) {
	log.Warn().Msg(fmt.Sprint(args...))
}

// Warnf warning log with format
func Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

// Warnw warn log with additional context
func Warnw(msg string, args ...interface{}) {
	log.Warn().Fields(sweetenFields(args)).Msg(msg)
}

// Error error log
func Error(args ...interface{}) {
	log.Error().Msg(fmt.Sprint(args...))
}

// Errorf error log with format
func Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

// Errorw error log with additional context
func Errorw(msg string, args ...interface{}) {
	log.Error().Fields(sweetenFields(args)).Msg(msg)
}

// Panic panic log
func Panic(args ...interface{}) {
	log.Panic().Msg(fmt.Sprint(args...))
}

// Panicf panic log with format
func Panicf(format string, args ...interface{}) {
	log.Panic().Msgf(format, args...)
}

// Panicw panic log with additional context
func Panicw(msg string, args ...interface{}) {
	log.Panic().Fields(sweetenFields(args)).Msg(msg)
}

func sweetenFields(args []interface{}) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}

	m := map[string]interface{}{}
	for i := 0; i < len(args)-1; i += 2 {
		key, value := args[i], args[i+1]

		if keyStr, ok := key.(string); ok {
			m[keyStr] = value
		} else {
			Warnf("%v not string type, ignored with value(%v)", key, value)
		}
	}

	return m
}
