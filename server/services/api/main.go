package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/ahdernasr/dailydininghall/server/proto/gen"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"google.golang.org/grpc"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_URL"),
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(limiter.New(limiter.Config{Max: 10, Expiration: 30 * time.Second, KeyGenerator: func(c *fiber.Ctx) string { return c.IP() }}))

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(200) })

	subsAddr := getenv("SUBSCRIBERS_ADDR", "localhost:50052")
	mailerAddr := getenv("MAILER_ADDR", "localhost:50053")

	// Create gRPC clients per request to avoid global state; connection pooling can be added later
	app.Post("/api/subscribe", func(c *fiber.Ctx) error {
		var req struct {
			Email string `json:"email"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		subsConn, err := grpc.Dial(subsAddr, grpc.WithInsecure())
		if err != nil {
			return c.Status(500).SendString("failed to connect to subscribers service")
		}
		defer subsConn.Close()
		subsClient := pb.NewSubscribersServiceClient(subsConn)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := subsClient.AddSubscriber(ctx, &pb.AddSubscriberRequest{Email: req.Email}); err != nil {
			return c.Status(400).SendString("Email is already subscribed.")
		}

		mailConn, err := grpc.Dial(mailerAddr, grpc.WithInsecure())
		if err == nil {
			defer mailConn.Close()
			mailClient := pb.NewMailerServiceClient(mailConn)
			_, _ = mailClient.SendSubscribeEmail(ctx, &pb.SendSubscribeEmailRequest{Email: req.Email})
		}

		return c.SendString("Success")
	})

	app.Post("/api/unsubscribe", func(c *fiber.Ctx) error {
		var req struct {
			Email string `json:"email"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		subsConn, err := grpc.Dial(subsAddr, grpc.WithInsecure())
		if err != nil {
			return c.Status(500).SendString("failed to connect to subscribers service")
		}
		defer subsConn.Close()
		subsClient := pb.NewSubscribersServiceClient(subsConn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := subsClient.RemoveSubscriber(ctx, &pb.RemoveSubscriberRequest{Email: req.Email}); err != nil {
			return c.Status(500).SendString("failed to unsubscribe")
		}
		return c.SendString("Success")
	})

	log.Fatal(app.Listen(":4000"))
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
