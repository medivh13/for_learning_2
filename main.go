package main

import (
	"for_learning_2/src/infra/config"
	"for_learning_2/src/server"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conf := config.Make()
	grpcServer := server.NewGRPCServer(

		server.WithConfig(&conf),
	)
	num, _ := strconv.Atoi(conf.Http.Port)
	grpcServer.Run(num)
}
