package belajargologging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.WithField("name", "Dicki Darmawan Saputra with field").WithField("subject", "programmer").Info("this is info with field")
}
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.WithFields(logrus.Fields{
		"name":    "Dicki with fields",
		"subject": "programmer",
	}).Info("this is info with field")
}
