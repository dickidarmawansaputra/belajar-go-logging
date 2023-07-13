package belajargologging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	// penggunaan hook jika ada level log tertentu bisa dieksekusi misal kirim notif dll
	logger.AddHook(&SampleHook{})

	logger.Info("hello info")
	logger.Warn("hello info")
	logger.Error("hello info")

}
