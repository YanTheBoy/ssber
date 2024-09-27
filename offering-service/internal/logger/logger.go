package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
)

type Log struct {
	original *slog.Logger
}

func New() Log {
	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if err, ok := a.Value.Any().(error); ok {
					aErr := tint.Err(err)
					aErr.Key = a.Key
					return aErr
				}
				return a
			},
		}),
	)
	return Log{logger}
}

func (l *Log) Error(msg string, args ...any) {
	l.original.Error(msg, args...)
}

func (l *Log) Info(msg string, args ...any) {
	l.original.Info(msg, args...)
}

func (l *Log) WithError(err error, msg string, args ...any) {
	passArgs := make([]any, len(args)+2)
	passArgs[0] = "error"
	passArgs[1] = err
	for i, arg := range args {
		passArgs[i+2] = arg
	}
	l.original.Error(msg, passArgs...)


}