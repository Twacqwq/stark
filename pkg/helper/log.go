package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(formatter logrus.Formatter) *logrus.Logger {
	log := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		// Formatter: &logrus.TextFormatter{},
		Level: logrus.InfoLevel,
		// Level:        convertLogLevel(DefaultLogLevel),
		ReportCaller: true,
		ExitFunc:     os.Exit,
		Hooks:        logrus.LevelHooks{},
	}

	// log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
	// 	logrus.PanicLevel,
	// 	logrus.FatalLevel,
	// 	logrus.ErrorLevel,
	// 	logrus.WarnLevel,
	// )))

	return log
}
