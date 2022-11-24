package api

import (
	"github.com/gorilla/mux"
)

// Server is used as a container for the most important dependencies.
type Server struct {
	Router *mux.Router
}

// NewServer returns a pointer to a new Server.
func NewServer() *Server {
	server := Server{
		Router: mux.NewRouter().StrictSlash(true),
	}
	server.registerRoutes()

	return &server
}
