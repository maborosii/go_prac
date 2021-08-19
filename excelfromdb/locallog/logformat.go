package locallog

import (
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	Log = logrus.New()
	Log.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})
	return Log
}
func init() {
	Log := NewLogger()

	Log.SetLevel(logrus.InfoLevel)
}
