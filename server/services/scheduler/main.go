package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/ahdernasr/dailydininghall/server/proto/gen"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

func main() {
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	c := cron.New(cron.WithLocation(location))
	_, err = c.AddFunc("0 6 * * *", daily)
	if err != nil {
		log.Fatalf("Failed to schedule the task: %v", err)
	}
	c.Start()

	select {} // block forever
}

func daily() {
	scraperAddr := getenv("SCRAPER_ADDR", "localhost:50051")
	subsAddr := getenv("SUBSCRIBERS_ADDR", "localhost:50052")
	mailerAddr := getenv("MAILER_ADDR", "localhost:50053")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get menu
	sconn, err := grpc.Dial(scraperAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("scraper dial err: %v", err)
		return
	}
	defer sconn.Close()
	sm := pb.NewScraperServiceClient(sconn)
	menuRes, err := sm.GetMenu(ctx, &pb.GetMenuRequest{})
	if err != nil {
		log.Printf("scraper GetMenu err: %v", err)
		return
	}

	// Get subscribers
	subConn, err := grpc.Dial(subsAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("subs dial err: %v", err)
		return
	}
	defer subConn.Close()
	subc := pb.NewSubscribersServiceClient(subConn)
	subsRes, err := subc.ListSubscribers(ctx, &pb.ListSubscribersRequest{})
	if err != nil {
		log.Printf("subs List err: %v", err)
		return
	}

	// Send emails
	mconn, err := grpc.Dial(mailerAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("mailer dial err: %v", err)
		return
	}
	defer mconn.Close()
	mc := pb.NewMailerServiceClient(mconn)
	_, err = mc.SendMenuEmail(ctx, &pb.SendMenuEmailRequest{Menu: menuRes.Menu, Subscribers: subsRes.Subscribers})
	if err != nil {
		log.Printf("send email err: %v", err)
		return
	}

	log.Printf("Emails sent successfully!")
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
