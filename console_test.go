package consolelogrus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// TestPrintAllLevels is a test function to print all log levels
func TestPrintAllLevels(t *testing.T) {

	logger := InitNewLogger(&CustomFormatter{
		PaddingEnabled: true,
	})

	logger.SetLevel(logrus.DebugLevel)
	logger.Debug("Debug level")
	logger.Info("Info level")
	logger.Warn("Warning level")
	logger.Error("Error level")
	logger.Fatal("Fatal level")
	logger.Panic("Panic level")
}
