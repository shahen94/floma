package floma

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shahen94/floma/domain"
)

// Logger is a wrapper around color.Color
type Logger struct {
	info    *color.Color
	warn    *color.Color
	error   *color.Color
	spinner *LogSpinner
}

// LogSpinner is a wrapper around spinner.Spinner
type LogSpinner struct {
	info  *spinner.Spinner
	warn  *spinner.Spinner
	error *spinner.Spinner
}

// Info prints a message with spinner and blue color
func (l *LogSpinner) Info(message ...string) {
	for _, m := range message {
		l.info.Suffix = " " + m
		go l.info.Start()
	}
}

// Warn prints a message with spinner and yellow color
func (l *LogSpinner) Warn(message ...string) {
	for _, m := range message {
		l.warn.Suffix = " " + m
		go l.warn.Start()
	}
}

// Error prints a message with spinner and red color
func (l *LogSpinner) Error(message ...string) {
	for _, m := range message {
		l.error.Suffix = " " + m
		go l.error.Start()
	}
}

const TAG = "[floma]"

// Info prints a message with blue color
func (l *Logger) Info(message string) {
	l.info.Println(TAG, message)
}

// Warn prints a message with yellow color
func (l *Logger) Warn(message string) {
	l.warn.Println(TAG, message)
}

// Error prints a message with red color
func (l *Logger) Error(message string) {
	l.error.Println(TAG, message)
}

// WithSpinner returns a spinner
func (l *Logger) WithSpinner() domain.LogSpinner {
	return l.spinner
}

// NewLogger returns a new Logger
func NewLogger() domain.Logger {
	infoColor := color.New(color.FgBlue)
	warnColor := color.New(color.FgYellow)
	errorColor := color.New(color.FgRed)

	infoSpinner := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	warnSpinner := spinner.New(spinner.CharSets[32], 100*time.Millisecond)
	errorSpinner := spinner.New(spinner.CharSets[32], 100*time.Millisecond)

	infoSpinner.Color("blue")
	warnSpinner.Color("yellow")
	errorSpinner.Color("red")

	logSpinner := &LogSpinner{
		info:  infoSpinner,
		warn:  warnSpinner,
		error: errorSpinner,
	}

	return &Logger{
		info:  infoColor,
		warn:  warnColor,
		error: errorColor,

		spinner: logSpinner,
	}
}
