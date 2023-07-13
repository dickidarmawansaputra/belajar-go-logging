package belajargologging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

// default: text formatter
// mengubah formatter logger.SetFormatter()
// 2 format text & json
func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
