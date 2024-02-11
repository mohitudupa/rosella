package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"githib.com/mohitudupa/rosella/utils"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func loggerResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, statusCode: 200}
}

func (rw *responseWriter) Status() int {
	return rw.statusCode
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	log.Printf("INFO: initializing middelware: LoggerMiddleware")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					utils.ErrorJsonResponse(w, "something went wrong", http.StatusInternalServerError)
					log.Printf("ERROR: %s", debug.Stack())
				}
			}()

			start := time.Now()
			lw := loggerResponseWriter(w)

			next.ServeHTTP(lw, r)

			log.Printf("INFO: status=%d method=%s path=%s duration=%v", lw.statusCode, r.Method, r.URL.EscapedPath(), time.Since(start))
		},
	)
}
