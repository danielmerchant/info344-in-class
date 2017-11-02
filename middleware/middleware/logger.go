package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	Handler http.Handler
}

type LoggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *LoggerResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *LoggerResponseWriter) Write(data []byte) (int, error) {
	log.Println("writing something")
	return lrw.ResponseWriter.Write(data)
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{Handler: handler}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lrw := &LoggerResponseWriter{w, http.StatusOK}
	start := time.Now()
	l.Handler.ServeHTTP(lrw, r)
	log.Printf("%s, %s, %d, %d", r.Method, r.URL.Path, lrw.statusCode, time.Since(start))
}
