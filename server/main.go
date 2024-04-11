package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/fanchann/grpc_math/proto"
	"github.com/fanchann/grpc_math/server/service"
)

func init() {
	// load env configuration
	if err := godotenv.Load(); err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_MATH_SERVICE_PORT"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	// grpc server
	grpcSvr := grpc.NewServer()

	// reflection API
	reflection.Register(grpcSvr)

	// math service
	mathSvc := service.NewMathServiceImpl()

	// grpc register
	proto.RegisterMathServiceServer(grpcSvr, mathSvc)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := grpcSvr.Serve(listener); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	<-stopCh
	log.Printf("shutting down server \n")
}
