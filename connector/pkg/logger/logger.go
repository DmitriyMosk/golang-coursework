package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger  *logrus.Logger
	logFile io.Writer
	errFile io.Writer
}

type LogLevel int

const (
	DEBUG   LogLevel = 0
	INFO    LogLevel = 1
	WARNING LogLevel = 2
	ERROR   LogLevel = 3
)

func CreateNewLogger() *Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logDir := "connector/logs"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		return nil
	}

	logsPath := filepath.Join(logDir, "logs.log")
	errorsPath := filepath.Join(logDir, "err_logs.log")

	logs, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return nil
	}

	errors, err := os.OpenFile(errorsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open error log file: %v\n", err)
		return nil
	}

	logFile := io.MultiWriter(logs)
	errFile := io.MultiWriter(os.Stdout, errors)

	return &Logger{
		logger:  logger,
		logFile: logFile,
		errFile: errFile,
	}
}

func (log *Logger) Log(logLevel LogLevel, logMessage string) {
	switch logLevel {
	case DEBUG:
		log.logger.Out = log.logFile
		log.logger.Debug(logMessage)
	case INFO:
		log.logger.Out = log.logFile
		log.logger.Infoln(logMessage)
	case WARNING:
		log.logger.Out = log.errFile
		log.logger.Warning(logMessage)
	case ERROR:
		log.logger.Out = log.errFile
		log.logger.Error(logMessage)
	}
}
