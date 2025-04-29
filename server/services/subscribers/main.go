package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ahdernasr/dailydininghall/internal/db"
	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	pb "github.com/ahdernasr/dailydininghall/server/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type subscribersServer struct {
	pb.UnimplementedSubscribersServiceServer
}

func (s *subscribersServer) AddSubscriber(ctx context.Context, req *pb.AddSubscriberRequest) (*pb.Empty, error) {
	if err := queries.AddSubscriber(req.Email); err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func (s *subscribersServer) RemoveSubscriber(ctx context.Context, req *pb.RemoveSubscriberRequest) (*pb.Empty, error) {
	if err := queries.RemoveSubscriber(req.Email); err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func (s *subscribersServer) ListSubscribers(ctx context.Context, req *pb.ListSubscribersRequest) (*pb.ListSubscribersResponse, error) {
	subs, err := queries.GetAllSubscribers()
	if err != nil {
		return nil, err
	}
	out := &pb.ListSubscribersResponse{}
	for _, v := range subs {
		out.Subscribers = append(out.Subscribers, &pb.Subscriber{Email: v.Email})
	}
	return out, nil
}

func main() {
	connStr := os.Getenv("CONNECTION_STRING")
	if err := db.Connect(connStr); err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSubscribersServiceServer(s, &subscribersServer{})
	reflection.Register(s)
	log.Println("subscribers service listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
