package main

import (
	"context"
	"github.com/ananasjesus/fibonacci"
	"github.com/ananasjesus/fibonacci/pkg/cache"
	"github.com/ananasjesus/fibonacci/pkg/fibonacci/handler"
	"github.com/ananasjesus/fibonacci/proto"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	redisCache := cache.NewRedisCache(redisClient)

	handlers := handler.NewHandler(redisCache)

	srv := new(fibonacci.Server)

	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	grpcRegistrar := grpc.NewServer()
	grpcFibonacciServer := &fibonacci.GRPCServer{Cache: redisCache}
	proto.RegisterFibonacciServer(grpcRegistrar, grpcFibonacciServer)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error occured while running grpc server: %s", err.Error())
	}
	go func() {
		if err := grpcRegistrar.Serve(listener); err != nil {
			log.Fatalf("error occured while running grpc server: %s", err.Error())
		}
	}()

	log.Print("Fibonacci service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Fibonacci service shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

}
