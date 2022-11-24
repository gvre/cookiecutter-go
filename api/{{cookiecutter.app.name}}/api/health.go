package api

import (
	"net/http"
	"runtime/debug"
)

var healthCheck struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
}

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	healthCheck.GoVersion = info.GoVersion
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			healthCheck.Version = kv.Value
		case "vcs.time":
			healthCheck.BuildDate = kv.Value
		}
	}
}

// HandleCheckLive is used for checking if the service is up.
func (s *Server) HandleCheckLive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		healthCheck.Status = "ok"
		_ = Ok(w, healthCheck)
	}
}
