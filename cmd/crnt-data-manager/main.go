package main

import (
	"context"
	"log"
	"net"
	"net/http"

	taskService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/task"
	"github.com/Constantine27K/crnt-data-manager/pkg/task"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	runApp()
}

func runApp() {
	go runTaskRest()
	runTaskGrpc()
}

func runTaskGrpc() {
	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	task.RegisterTaskRegistryServer(s, taskService.NewService())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func runTaskRest() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := task.RegisterTaskRegistryHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
