package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	message "B/protobuffer/proto"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("addr", "localhost:8080", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	log.Println("Run test client for method: Simple RPC")
	SimpleClient(conn)

	fmt.Println("")

	log.Println("Run test client for method: Server-side streaming RPC")
	ServersideStreamingRPC(conn)
}

func SimpleClient(conn *grpc.ClientConn) {
	client := message.NewGetMessageClient(conn)
	request := &message.ARequest{
		Name: "ThinhNguyenV",
	}

	resp, err := client.Get(context.Background(), request)
	if err != nil {
		fmt.Println("Got err: ", err.Error())
		return
	}

	fmt.Println(resp.Result)
}

func ServersideStreamingRPC(conn *grpc.ClientConn) {
	client := message.NewGetMessageClient(conn)
	request := &message.ARequests{
		Names: []string{"Name A", "Name B", "Name C"},
	}
	stream, err := client.GetStreaming(context.Background(), request)
	if err != nil {
		fmt.Println("Request got err: ", err.Error())
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("No more stream")
			break
		}

		if err != nil {
			fmt.Println("Receive stream got err: ", err.Error())
		}

		fmt.Println(resp.Result)
	}
}
