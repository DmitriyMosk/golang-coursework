package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() *log.Logger {
	logFile, err := os.OpenFile("../../../logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	errLogFile, err := os.OpenFile("../../../logs/err_logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New()
	logger.SetOutput(logFile)
	logger.SetOutput(errLogFile)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
