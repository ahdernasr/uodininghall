package main

import (
	"context"
	"log"
	"net"

	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/ahdernasr/dailydininghall/internal/mailer"
	"github.com/ahdernasr/dailydininghall/internal/scraper"
	pb "github.com/ahdernasr/dailydininghall/server/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type mailerServer struct {
	pb.UnimplementedMailerServiceServer
}

func (s *mailerServer) SendMenuEmail(ctx context.Context, req *pb.SendMenuEmailRequest) (*pb.SendMenuEmailResponse, error) {
	// Convert proto menu to internal types
	in := req.Menu
	toMeal := func(m *pb.Meal) scraper.Meal {
		conv := func(ds []*pb.Dish) []scraper.Dish {
			out := make([]scraper.Dish, 0, len(ds))
			for _, d := range ds {
				out = append(out, scraper.Dish{Name: d.Name, Description: d.Description, Allergies: d.Allergies})
			}
			return out
		}
		return scraper.Meal{Breakfast: conv(m.Breakfast), Lunch: conv(m.Lunch), Dinner: conv(m.Dinner), Other: conv(m.Other)}
	}
	menu := &scraper.Menu{
		Grill:         toMeal(in.Grill),
		MindBodySoul:  toMeal(in.MindBodySoul),
		PlantBase:     toMeal(in.PlantBase),
		ServiceMinute: toMeal(in.ServiceMinute),
		Trattoria:     toMeal(in.Trattoria),
		WorldFlavours: toMeal(in.WorldFlavours),
	}

	subs := make([]queries.Subscriber, 0, len(req.Subscribers))
	for _, s := range req.Subscribers {
		subs = append(subs, queries.Subscriber{Email: s.Email})
	}

	if err := mailer.SendMenuEmail(menu, subs); err != nil {
		return nil, err
	}
	return &pb.SendMenuEmailResponse{}, nil
}

func (s *mailerServer) SendSubscribeEmail(ctx context.Context, req *pb.SendSubscribeEmailRequest) (*pb.SendSubscribeEmailResponse, error) {
	if err := mailer.SendSubscribeEmail(req.Email); err != nil {
		return nil, err
	}
	return &pb.SendSubscribeEmailResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMailerServiceServer(s, &mailerServer{})
	reflection.Register(s)
	log.Println("mailer service listening on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
