package main

import (
	"context"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
	"waylon.com/demo/grpcPubsub/rpc"
	"net"
	"strings"
	"time"
)

func main() {
	grpcServer := grpc.NewServer()
	rpc.RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *rpc.String, ) (*rpc.String, error) {
	p.pub.Publish(arg.GetValue())
	return &rpc.String{}, nil
}

func (p *PubsubService) Subscribe(arg *rpc.String, stream rpc.PubsubService_SubscribeServer, ) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&rpc.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
