package belajargologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// ini untuk versi cepat, tapi sebaiknya gunakan object baru seperti sebelumnya
func TestSingleton(t *testing.T) {
	logrus.Info("hello world")
	logrus.Warn("hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("hello world")
	logrus.Warn("hello world")
}
