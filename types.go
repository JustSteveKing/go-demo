package main

import "net/http"

// Build info set via -ldflags at build time (optional).
var (
	buildVersion = "dev"
	buildCommit  = ""
	buildTime    = ""
)

// Info holds build metadata for the binary.
type Info struct {
	Version string
	Commit  string
	Built   string
}

// BuildInfo returns the build metadata.
func BuildInfo() Info {
	return Info{
		Version: buildVersion,
		Commit:  buildCommit,
		Built:   buildTime,
	}
}

// responseWriter wraps http.ResponseWriter to record the status code.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader records the status code and writes the header.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
