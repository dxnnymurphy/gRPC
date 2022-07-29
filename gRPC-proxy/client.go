package main

import (
	"context"
	"log"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	pb "dxnnymurphy/gRPC/pb"
	"google.golang.org/grpc"
	"github.com/rs/cors"
)

func main() {
	var grpcServerEndpoint = "localhost:58883"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAnomalyDetectionHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Listening on port 8081")
	port := ":8081"
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(port, handler)
}