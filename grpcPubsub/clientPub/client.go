package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"waylon.com/demo/grpcPubsub/rpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &rpc.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &rpc.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(
		context.Background(), &rpc.String{Value: "golang: hello waylon"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
