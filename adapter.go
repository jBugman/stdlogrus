// Package stdlogrus provides some glue for using logrus from std logger.
// It is based on code from https://github.com/sirupsen/logrus/issues/118#issuecomment-345475880
package stdlogrus

import (
	"io"
	"log"

	"github.com/sirupsen/logrus"
)

// Setup forwards std logger to logrus.
func Setup() {
	log.SetOutput(New())
}

// New returns new adapter.
func New() io.Writer {
	return &log2LogrusWriter{
		entry: logrus.StandardLogger().WithField("logger", "std"),
	}
}

// log2LogrusWriter exploits the documented fact that the standard
// log pkg sends each log entry as a single io.Writer.Write call:
// https://golang.org/pkg/log/#Logger
type log2LogrusWriter struct {
	entry *logrus.Entry
}

func (w *log2LogrusWriter) Write(b []byte) (int, error) {
	n := len(b)
	if n > 0 && b[n-1] == '\n' {
		b = b[:n-1]
	}
	w.entry.Info(string(b))
	return n, nil
}
