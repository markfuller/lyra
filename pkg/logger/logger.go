package logger

import (
	"io"
	"sync"

	hclog "github.com/hashicorp/go-hclog"
)

const (
	logsDisabled = 100
)

var logger hclog.Logger
var once sync.Once

// Spec describes the logger to be created
type Spec struct {
	Name            string
	Level           string
	Output          io.Writer
	JSON            bool
	IncludeLocation bool
}

// Get returns the initialised Logger
func Get() hclog.Logger {
	return logger
}

// Initialise the Logger
func Initialise(spec Spec) hclog.Logger {
	once.Do(func() {

		hclog.DefaultOptions = &hclog.LoggerOptions{
			Name:            spec.Name,
			Level:           hclog.Level(logsDisabled),
			JSONFormat:      spec.JSON,
			IncludeLocation: spec.IncludeLocation,
		}
		if len(spec.Level) > 0 {
			ll := hclog.LevelFromString(spec.Level)
			if ll != hclog.NoLevel {
				hclog.DefaultOptions.Level = ll
				level = spec.Level
			}
		}
		if spec.Output != nil {
			hclog.DefaultOptions.Output = spec.Output
		}
		l := hclog.Default()
		logger = l
	})
	return logger
}

var level string

// Level returns the log level as a string
func Level() string {
	return level
}
