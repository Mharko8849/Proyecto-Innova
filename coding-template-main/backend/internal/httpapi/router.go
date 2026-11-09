package httpapi

import (
	"net/http"

	"indicepartidas/backend/internal/health"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Handler)
	return mux
}
