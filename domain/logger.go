package domain

type Logger interface {
	Info(message string)
	Warn(message string)
	Error(message string)
	WithSpinner() LogSpinner
}

type LogSpinner interface {
	Info(message ...string)
	Warn(message ...string)
	Error(message ...string)
}
