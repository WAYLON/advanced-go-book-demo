package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"waylon.com/demo/grpc/rpc"
	"waylon.com/demo/grpcToken/auth"
	"net"
)

type HelloServiceImpl struct{auth *auth.Authentication}

func (p *HelloServiceImpl) Channel(stream rpc.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &rpc.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *rpc.String) (*rpc.String, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	reply := &rpc.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func filter(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))
	a := &HelloServiceImpl{
		&auth.Authentication{User: "gopher",Password: "password"},
	}

	rpc.RegisterHelloServiceServer(grpcServer, a)
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}