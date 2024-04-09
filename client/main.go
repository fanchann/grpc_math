package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/fanchann/grpc_math/client/routes"
	"github.com/fanchann/grpc_math/proto"
)

func init() {
	// load env configuration
	if err := godotenv.Load(); err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	mathServiceUrl := fmt.Sprintf("localhost:%s", os.Getenv("GRPC_MATH_SERVICE_PORT"))

	conn, _ := grpc.Dial(mathServiceUrl, grpc.WithInsecure())

	client := proto.NewMathServiceClient(conn)

	urlRoute := routes.NewRoutes(client)

	r := http.NewServeMux()

	r.HandleFunc("GET /add/{a}/{b}", urlRoute.Add)
	r.HandleFunc("GET /mul/{a}/{b}", urlRoute.Multiply)
	r.HandleFunc("GET /div/{a}/{b}", urlRoute.Devide)
	r.HandleFunc("GET /red/{a}/{b}", urlRoute.Reduce)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("client started on port %v", os.Getenv("CLIENT_APP_PORT"))
		if err := http.ListenAndServe(":"+os.Getenv("CLIENT_APP_PORT"), r); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	<-stopCh
	log.Print("shutting down server")

}
