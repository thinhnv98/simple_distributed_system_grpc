package main

import (
	"context"
	"flag"
	"fmt"
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

	client := message.NewGetMessageClient(conn)
	feature, err := client.Get(context.Background(), &message.ARequest{
		Name: "ThinhNguyenV",
	})
	if err != nil {
		fmt.Println("Got err: ", err.Error())
		return
	}

	fmt.Println(feature.Result)
}
