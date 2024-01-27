package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(formatter logrus.Formatter) *logrus.Logger {
	log := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		Level:     logrus.InfoLevel,
		// Level:        convertLogLevel(DefaultLogLevel),
		ReportCaller: true,
		ExitFunc:     os.Exit,
		Hooks:        logrus.LevelHooks{},
	}

	return log
}
