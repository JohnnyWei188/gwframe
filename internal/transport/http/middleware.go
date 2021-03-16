package http

import (
	"net/http"
)

// Middleware is HTTP transport middleware
type Middleware func(next http.Handler) http.Handler

// Chain returns a Middleware that specifies the chained handler for endpoint.
func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
