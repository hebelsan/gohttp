// Package auth contains all authentication checks and middleware
package auth

import (
	"fmt"
	"net/http"

	"github.com/hebelsan/gohttp/pkg/util"
)

// ENV_KEY env key which defines the auth type
const ENV_KEY = "AUTH"

// TYPE_API_KEY defines the API KEY auth method
const TYPE_API_KEY = "API-KEY"

// Handler defines a general handler function
type Handler func(next http.HandlerFunc) http.HandlerFunc

type middleware struct {
	AuthHandler Handler
	apiKey      string
}

// NewMiddleware instantiates a new Middleware with defaults for not provided Options.
func NewMiddleware(authType string) *middleware {
	m := middleware{}
	switch authType {
	case TYPE_API_KEY:
		m.apiKey = util.PseudoUuid()
		m.AuthHandler = m.handleApiKey
	default:
		m.AuthHandler = handleNoAuth
	}
	return &m
}

func (m *middleware) handleApiKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if api-key is correct
		apiKeyHeader := r.Header.Get("X-API-KEY")
		if apiKeyHeader == "" {
			http.Error(w, "X-API-KEY header missing", http.StatusUnauthorized)
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

func handleNoAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

func logApiKey(key string) {
	fmt.Println("New Api-Key: " + key)
}
