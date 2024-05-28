package server

import (
	"for_learning_2/src/infra/config"
	"for_learning_2/src/infra/persistence/redis"
	handler "for_learning_2/src/interface/grpc/handlers"

	bookProto "for_learning_2/src/app/proto/books"
	usecases "for_learning_2/src/app/usecase"
	bookUC "for_learning_2/src/app/usecase/books"
	pickUpUC "for_learning_2/src/app/usecase/pickup"
	"for_learning_2/src/infra/broker/nats"
	pickUpNats "for_learning_2/src/infra/broker/nats/consumer/pickup"
	natsPublisher "for_learning_2/src/infra/broker/nats/publisher"
	circuit_breaker_service "for_learning_2/src/infra/circuit_breaker"
	bookInteg "for_learning_2/src/infra/integration/books"
	ms_log "for_learning_2/src/infra/log"
	redisService "for_learning_2/src/infra/persistence/redis/service"
	// sentryClient "for_learning_2/src/infra/sentry_init"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is struct to hold any dependencies used for server
type Server struct {
	config *config.Config
}

type ServerGrpcOption func(*Server)

// NewGRPCServer is constructor
// func NewGRPCServer(conf *config.Config, repo *service.Repositories) *Server {
func NewGRPCServer(options ...ServerGrpcOption) *Server {
	server := &Server{}

	for _, option := range options {
		option(server)
	}

	return server
}

func (s *Server) Run(port int) error {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	m := make(map[string]interface{})
	m["env"] = s.config.App.Environment
	m["service"] = s.config.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(s.config.Log.Name),
		ms_log.IsProduction(false),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(s.config.Redis, logger)
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to initialize redis client")
	}
	redisSvc := redisService.NewServRedis(redisClient)
	// err = sentryClient.NewSentryClient(s.config.Sentry, logger)
	// if err != nil {
	// 	logger.WithField("error", err).Fatal("Failed to initialize Sentry")
	// }

	circuitBreaker := circuit_breaker_service.NewCircuitBreakerInstance()
	bookIntegration := bookInteg.NewIntegOpenLibrary(circuitBreaker)

	Nats := nats.NewNats(s.config.Nats, logger)
	publisher := natsPublisher.NewPushWorker(Nats)
	// HTTP Handler
	// the server already implements a graceful shutdown.

	allUC := usecases.AllUseCases{
		BookUC:   bookUC.NewBooksUseCase(bookIntegration, redisSvc),
		PickUpUC: pickUpUC.NewPickUpUseCase(publisher),
	}

	pickUpNats.NewPickUpWorker(Nats, allUC.PickUpUC)

	handlers := handler.NewHandler(s.config, allUC)

	// register from proto
	bookProto.RegisterBookServiceServer(server, handlers)

	// register reflection
	reflection.Register(server)

	return RunGRPCServer(server, port)
}
