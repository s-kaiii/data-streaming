package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"real-time-streaming/internal/utils/errList"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{ // JSONFormatter for structured logging
		TimestampFormat: time.RFC3339,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel) // Set log level to debug; change as needed
}

func Info(msg string) {
	Logger.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"msg":       msg,
	}).Info()
}

func Request(method, path, clientIP string) {
	Logger.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"type":      "request",
		"method":    method,
		"path":      path,
		"client_ip": clientIP,
	}).Info("Incoming request")
}

func KafkaInteraction(topic, key, action string, success bool) {
	status := "failed"
	if success {
		status = "succeeded"
	}

	Logger.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"type":      "kafka",
		"topic":     topic,
		"key":       key,
		"action":    action,
		"status":    status,
	}).Info("Kafka interaction")
}

func Error(errMap errList.CoreError, context string) {
	Logger.WithFields(logrus.Fields{
		"type":           "error",
		"error_code":     errMap.Code,
		"error_category": errMap.Category,
		"error_type":     errMap.EventErrorType,
		"error_message":  errMap.Description,
		"context":        context,
		"timestamp":      time.Now().Format(time.RFC3339),
	}).Error("Error occurred")
}
