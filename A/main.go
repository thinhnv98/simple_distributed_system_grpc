package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"A/controller"
	message "A/protobuffer/proto"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	message.RegisterGetMessageServer(grpcServer, controller.MessageController{})
	log.Println("Server ready on :8080")
	grpcServer.Serve(lis)
}
