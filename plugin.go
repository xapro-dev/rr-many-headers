package custom

import (
	"net/http"
	"regexp"

	"github.com/roadrunner-server/config"
)

const PluginName = "many_headers"

// Plugin serves headers files. Potentially convert into middleware?
type Plugin struct {
	// server configuration (location, forbidden files and etc)
	cfg *Config
}

func (s *Plugin) Init() error {
	return nil
}

// Middleware is HTTP plugin middleware to serve headers
func (s *Plugin) Middleware(next http.Handler) http.Handler {
	// Define the http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := s.cfg.routes {
			match, _ = regexp.MatchString(route.regex, r.URL.RawPath)
			if (match) {
				for k, v := range s.cfg.headers {
					w.Header().Set(k, v)				
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
