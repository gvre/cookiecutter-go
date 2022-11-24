package api

import (
	_ "embed"
	"net/http"
)

//go:embed spec/openapi.yaml
var spec []byte

// registerRoutes registers all app routes.
func (s *Server) registerRoutes() {
	// Health check
	s.Router.HandleFunc("/", s.HandleCheckLive()).Methods("GET")

	// API spec
	s.Router.HandleFunc("/spec", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		_, _ = w.Write(spec)
	}).Methods("GET")
}
