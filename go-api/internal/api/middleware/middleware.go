package middleware

import (
    "log"
    "net/http"
    "time"
)


// In middlewares, a function is returned to promote function chaining og http.HandleFunc so that each middleware chain returns a function chain, for easier wrapping and functionality
// LoggingMiddleware logs the incoming HTTP request & its duration
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
        log.Printf("Completed in %v", time.Since(start))
    })
}
