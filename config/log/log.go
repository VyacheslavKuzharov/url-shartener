package log

type Level string

const (
	InfoLevel  Level = "info"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
)

type Log struct {
	Level Level
}

func New() *Log {
	return &Log{
		Level: InfoLevel,
	}
}
