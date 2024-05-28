package server

import (
	"for_learning_2/src/infra/config"
	
)

// WithConfig is function
func WithConfig(config *config.Config) ServerGrpcOption {
	return func(r *Server) {
		r.config = config
	}
}

