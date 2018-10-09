package logrusx

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New() *logrus.Logger {
	l := logrus.New()
	ll, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		ll = logrus.InfoLevel
	}
	l.Level = ll

	if viper.GetString("LOG_FORMAT") == "json" {
		l.Formatter = new(logrus.JSONFormatter)
	}

	return l
}

func HelpMessage() string {
	return `- LOG_LEVEL: Set the log level, supports "panic", "fatal", "error", "warn", "info" and "debug". Defaults to "info".

	Example: LOG_LEVEL=panic

- LOG_FORMAT: Leave empty for text based log format, or set to "json" for JSON formatting.

	Example: LOG_FORMAT=json`
}
