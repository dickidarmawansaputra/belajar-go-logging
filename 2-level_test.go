package belajargologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	// yang keluar di console hanya yg diatas info, bisa mengubah level dengan logger.SetLevel()
	logger.Trace("this is trace")
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()

	// menampilkan log di console dari trace level
	// jika untuk development ok untuk debug level, saat production cukup dari info level ke atas
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("this is trace")
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
