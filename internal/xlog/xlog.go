package xlog

import (
	"context"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/config"
	"github.com/FarmerChillax/stark/pkg/helper"
	"github.com/sirupsen/logrus"
)

type logger struct {
	log *logrus.Logger
}

func InitGlobalLogger(loggerConf *config.LoggerConfig) {
	if loggerConf.Formatter == nil {
		loggerConf.Formatter = NewFormatter()
	}
	l := helper.NewLogger(loggerConf.Formatter)
	stark.Logger = &logger{
		log: l,
	}
}

func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.log.WithContext(ctx).Info(args...)
}

func (l *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Infof(format, args...)
}

func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	l.log.WithContext(ctx).Warn(args...)
}

func (l *logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Warnf(format, args...)
}

func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.log.WithContext(ctx).Error(args...)
}

func (l *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.log.WithContext(ctx).Errorf(format, args...)
}
