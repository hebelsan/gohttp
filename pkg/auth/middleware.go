// Package auth contains all authentication checks and middleware
package auth

import (
	"fmt"
	"net/http"

	"github.com/hebelsan/gohttp/pkg/util"
)

type middleware struct {
	apiKey string
}

// NewMiddleware instantiates a new Middleware with defaults for not provided Options.
func NewMiddleware() *middleware {
	m := new(middleware)
	m.apiKey = util.PseudoUuid()
	logApiKey(m.apiKey)
	return m
}

// Handle verifies the auth header for each request
func (m *middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if api-key is correct
		apiKeyHeader := r.Header.Get("X-API-KEY")
		if apiKeyHeader == "" {
			http.Error(w, "X-API-KEY header missing", http.StatusBadRequest)
			return
		}
		if apiKeyHeader != m.apiKey {
			http.Error(w, "wrong api-key", http.StatusUnauthorized)
			return
		}

		// generate new api key
		m.apiKey = util.PseudoUuid()
		logApiKey(m.apiKey)

		next.ServeHTTP(w, r)
	}
}

func logApiKey(key string) {
	fmt.Println("New Api-Key: " + key)
}
