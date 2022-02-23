package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"waylon.com/demo/grpc/rpc"
	"waylon.com/demo/grpcToken/auth"
	"time"
)

func main() {
	auth := auth.Authentication{
		User:     "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rpc.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &rpc.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&rpc.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}

}
