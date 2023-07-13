package belajargologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// cara manual untuk nunjukin dibungkus object Entry
func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("name", "Dicki").Info("this is info with entry")
}
