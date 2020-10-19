package log

import (
	"github.com/sirupsen/logrus"
)

// ConfigLog - Configuration for log app
func ConfigLog(logLevel *string) {
	switch *logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
		break
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
		break
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
