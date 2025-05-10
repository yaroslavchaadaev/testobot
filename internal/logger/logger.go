package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// New создаёт и настраивает логгер
func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log
}
