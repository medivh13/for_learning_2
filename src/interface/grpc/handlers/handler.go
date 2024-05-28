package handler

import (
	book "for_learning_2/src/app/proto/books"
	"for_learning_2/src/infra/config"
	usecases "for_learning_2/src/app/usecase"
)

// Interface is an interface
type Interface interface {
	// interface of grpc handler
	// book.BookServiceServer
	book.BookServiceServer
}

// Handler is struct
type Handler struct {
	config   *config.Config
	useCases usecases.AllUseCases
	book.UnimplementedBookServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, uc usecases.AllUseCases) *Handler {
	return &Handler{
		config: conf,
		// repo:       repo,
		// grpcClient: grpcClient,
		useCases: uc,
	}
}

var _ Interface = &Handler{}
