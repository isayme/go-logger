package logger

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z"
	zerolog.CallerSkipFrameCount++

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

type Logger struct {
	l zerolog.Logger

	ctx context.Context
}

func NewLogger() *Logger {
	l := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	return &Logger{
		l: l,
	}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := *l
	nl.ctx = ctx
	return &nl
}

// GetLevel get current level
func (l *Logger) GetLevel() string {
	return l.l.GetLevel().String()
}

// SetLevel set log level
func (l *Logger) SetLevel(lvl string) error {
	lvl = strings.ToLower(lvl)
	level, err := zerolog.ParseLevel(lvl)
	if err != nil {
		return err
	}

	l.l.Level(level)
	return nil
}

// Trace trace log
func (l *Logger) Trace(args ...interface{}) {
	l.l.Trace().Msg(fmt.Sprint(args...))
}

// Tracef trace log with format
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.l.Trace().Msgf(format, args...)
}

// Tracew trace log with additional context
func (l *Logger) Tracew(msg string, args ...interface{}) {
	l.l.Trace().Fields(l.sweetenFields(args)).Msg(msg)
}

// Debug debug log
func (l *Logger) Debug(args ...interface{}) {
	l.l.Debug().Msg(fmt.Sprint(args...))
}

// Debugf debug log with format
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.l.Debug().Msgf(format, args...)
}

// Debugw debug log with additional context
func (l *Logger) Debugw(msg string, args ...interface{}) {
	l.l.Debug().Fields(l.sweetenFields(args)).Msg(msg)
}

// Info info log
func (l *Logger) Info(args ...interface{}) {
	l.l.Info().Msg(fmt.Sprint(args...))
}

// Infof info log with format
func (l *Logger) Infof(format string, args ...interface{}) {
	l.l.Info().Msgf(format, args...)
}

// Infow info log with additional context
func (l *Logger) Infow(msg string, args ...interface{}) {
	l.l.Info().Fields(l.sweetenFields(args)).Msg(msg)
}

// Warn warning log
func (l *Logger) Warn(args ...interface{}) {
	l.l.Warn().Msg(fmt.Sprint(args...))
}

// Warnf warning log with format
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.l.Warn().Msgf(format, args...)
}

// Warnw warn log with additional context
func (l *Logger) Warnw(msg string, args ...interface{}) {
	l.l.Warn().Fields(l.sweetenFields(args)).Msg(msg)
}

// Error error log
func (l *Logger) Error(args ...interface{}) {
	l.l.Error().Msg(fmt.Sprint(args...))
}

// Errorf error log with format
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.l.Error().Msgf(format, args...)
}

// Errorw error log with additional context
func (l *Logger) Errorw(msg string, args ...interface{}) {
	l.l.Error().Fields(l.sweetenFields(args)).Msg(msg)
}

// Panic panic log
func (l *Logger) Panic(args ...interface{}) {
	l.l.Panic().Msg(fmt.Sprint(args...))
}

// Panicf panic log with format
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.l.Panic().Msgf(format, args...)
}

// Panicw panic log with additional context
func (l *Logger) Panicw(msg string, args ...interface{}) {
	l.l.Panic().Fields(l.sweetenFields(args)).Msg(msg)
}

func (l *Logger) sweetenFields(args []interface{}) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}

	m := map[string]interface{}{}
	for i := 0; i < len(args)-1; i += 2 {
		key, value := args[i], args[i+1]

		if keyStr, ok := key.(string); ok {
			m[keyStr] = value
		} else {
			l.Warnf("%v not string type, ignored with value(%v)", key, value)
		}
	}

	return m
}
