package interceptor

import (
	"net/http"
	"time"

	"github.com/SangBejoo/parking-space-monitor/init/logger"
)

type Middleware func(http.Handler) http.Handler

type BaseInterceptor struct {
	middlewares []Middleware
}

func NewBaseInterceptor() *BaseInterceptor {
	return &BaseInterceptor{
		middlewares: make([]Middleware, 0),
	}
}

func (b *BaseInterceptor) Use(middleware Middleware) {
	b.middlewares = append(b.middlewares, middleware)
}

func (b *BaseInterceptor) Wrap(handler http.Handler) http.Handler {
	for i := len(b.middlewares) - 1; i >= 0; i-- {
		handler = b.middlewares[i](handler)
	}
	return handler
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info("Request: %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		logger.Info("Response: %s %s - Completed in %v",
			r.Method, r.URL.Path, time.Since(start))
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
