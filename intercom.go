package intercom

import (
	"fmt"
	"os"
	"strings"
)

const (
	// SilentLevel is the "silent" log level.
	SilentLevel = "silent"
	// ErrorLevel is the "error" log level.
	ErrorLevel = "error"
	// WarnLevel is the "warn" log level.
	WarnLevel = "warn"
	// InfoLevel is the "info" log level.
	InfoLevel = "info"
	// DebugLevel is the "debug" log level.
	DebugLevel = "debug"
)

const (
	silentLevel = iota
	errorLevel
	warnLevel
	infoLevel
	debugLevel
)

// Logger is an intercom logger
type Logger struct {
	Level int
}

// NewLogger returns a Logger instance with the provided log level,
// defaulting to "info" if the provided level is not recognized as
// one of SilentLevel, ErrorLevel, WarnLevel, InfoLevel, or DebugLevel.
func NewLogger(level string) *Logger {
	var intLevel int

	switch strings.ToLower(level) {
	case SilentLevel:
		intLevel = silentLevel
	case ErrorLevel:
		intLevel = errorLevel
	case WarnLevel:
		intLevel = warnLevel
	case InfoLevel:
		intLevel = infoLevel
	case DebugLevel:
		intLevel = debugLevel
	default:
		intLevel = infoLevel
	}

	return &Logger{Level: intLevel}
}

// Errorf logs a red formatted string followed by a new line.
func (l *Logger) Errorf(message string, args ...interface{}) {
	if l.Level < errorLevel {
		return
	}

	colorMessage := fmt.Sprintf("\033[1;31m%s\033[0m\n", message)
	fmt.Fprintf(os.Stderr, colorMessage, args...)
}

// Warnf logs a yellow formatted string followed by a new line.
func (l *Logger) Warnf(message string, args ...interface{}) {
	if l.Level < warnLevel {
		return
	}

	colorMessage := fmt.Sprintf("\033[1;33m%s\033[0m\n", message)
	fmt.Fprintf(os.Stderr, colorMessage, args...)
}

// Infof logs a green formatted string followed by a new line.
func (l *Logger) Infof(message string, args ...interface{}) {
	if l.Level < infoLevel {
		return
	}

	colorMessage := fmt.Sprintf("\033[1;32m%s\033[0m\n", message)
	fmt.Fprintf(os.Stderr, colorMessage, args...)
}

// Debugf logs a blue formatted string followed by a new line.
func (l *Logger) Debugf(message string, args ...interface{}) {
	if l.Level < debugLevel {
		return
	}

	colorMessage := fmt.Sprintf("\033[1;34m%s\033[0m\n", message)
	fmt.Fprintf(os.Stderr, colorMessage, args...)
}
