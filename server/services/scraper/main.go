package main

import (
	"context"
	"log"
	"net"

	"github.com/ahdernasr/dailydininghall/internal/scraper"
	pb "github.com/ahdernasr/dailydininghall/server/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type scraperServer struct {
	pb.UnimplementedScraperServiceServer
}

func (s *scraperServer) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	m := scraper.Scraper()
	return &pb.GetMenuResponse{Menu: toProtoMenu(m)}, nil
}

func toProtoDish(d scraper.Dish) *pb.Dish {
	return &pb.Dish{Name: d.Name, Description: d.Description, Allergies: d.Allergies}
}

func toProtoMeal(meal scraper.Meal) *pb.Meal {
	p := &pb.Meal{}
	for _, x := range meal.Breakfast {
		p.Breakfast = append(p.Breakfast, toProtoDish(x))
	}
	for _, x := range meal.Lunch {
		p.Lunch = append(p.Lunch, toProtoDish(x))
	}
	for _, x := range meal.Dinner {
		p.Dinner = append(p.Dinner, toProtoDish(x))
	}
	for _, x := range meal.Other {
		p.Other = append(p.Other, toProtoDish(x))
	}
	return p
}

func toProtoMenu(m *scraper.Menu) *pb.Menu {
	return &pb.Menu{
		Grill:         toProtoMeal(m.Grill),
		MindBodySoul:  toProtoMeal(m.MindBodySoul),
		PlantBase:     toProtoMeal(m.PlantBase),
		ServiceMinute: toProtoMeal(m.ServiceMinute),
		Trattoria:     toProtoMeal(m.Trattoria),
		WorldFlavours: toProtoMeal(m.WorldFlavours),
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterScraperServiceServer(s, &scraperServer{})
	reflection.Register(s)

	log.Println("scraper service listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
