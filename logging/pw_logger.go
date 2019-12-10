package logging

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// PWLogger is a custom logger which implements the LogFormatter interface in go-chi/middleware package
type PWLogger struct {
	Logger *logrus.Logger
}

func New(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&PWLogger{Logger: logger})
}

var Logger = logrus.New()

func (l *PWLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &PWLoggerEntry{Logger: l.Logger}
	logFields := logrus.Fields{}

	logFields["@timestamp"] = time.Now().Format(time.RFC3339Nano)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	// if r.TLS != nil {
	// 	scheme = "https"
	// }
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Infoln("request started")

	return entry
}
