package logging

import (
	"context"
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func GetLogger(ctx context.Context) *PWLoggerEntry {
	logger := ctx.Value(middleware.LogEntryCtxKey)
	if logger != nil {
		logEntry := logger.(*PWLoggerEntry)
		logEntry.Logger = logEntry.Logger.WithField("@timestamp", time.Now().Format(time.RFC3339Nano))
		return logEntry
	}
	return &PWLoggerEntry{Logger: Logger}
}

func InitializeLogger(config *config.GeneralConfig) {

	l, e := logrus.ParseLevel("info")
	if e != nil {
		l = logrus.InfoLevel
	}

	logpath := config.LogFilePath
	var logFormatter = logrus.Logger{}
	logFormatter.Formatter = new(logrus.TextFormatter)

	logfile, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err == nil {
		*Logger = logrus.Logger{
			Out:       logfile,
			Formatter: logFormatter.Formatter,
			Hooks:     make(logrus.LevelHooks),
			Level:     l}
	}
	Logger.Info("initialized Logger successfully")
}

func SetLogLevel(loglevel string, logger *logrus.Logger) {
	l, e := logrus.ParseLevel(loglevel)
	if e != nil {
		l = logrus.InfoLevel
	}
	logger.Level = l
}
