package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Analytics struct {
	TotalSubscriptions int
	TotalVisits        int
}

type EmailRequest struct {
	Email string `json:"email"`
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// To be removed when database is integrated
	subscriptions := Analytics{0, 0}

	fmt.Println(subscriptions.TotalSubscriptions)

	// Health check the API
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Tracks unique visits
	app.Get("/api/track-visit", func(c *fiber.Ctx) error {
		// Check if the visitor has a specific cookie
		cookie := c.Cookies("unique_visitor")
		if cookie == "" {
			// Set a new cookie
			newCookie := new(fiber.Cookie)
			newCookie.Name = "unique_visitor"
			newCookie.Value = "1"
			newCookie.Expires = time.Now().Add(24 * time.Hour) // Expires in 1 day
			c.Cookie(newCookie)

			subscriptions.TotalVisits = subscriptions.TotalVisits + 1
		}

		fmt.Println("Tracked visitors. Total: ", subscriptions.TotalVisits)

		return c.SendString(fmt.Sprint(subscriptions.TotalSubscriptions))
	})

	app.Post("/api/subscribe", func(c *fiber.Ctx) error {
		//Todo check for duplicates

		req := new(EmailRequest)

		// Parse the body into the struct
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		//Record a new subscription
		subscriptions.TotalSubscriptions = subscriptions.TotalSubscriptions + 1

		// Use the email from the request
		fmt.Println("Total subscriptions: ", subscriptions.TotalSubscriptions)
		fmt.Printf("Received email: %s\n", req.Email)

		// Here you can add logic to handle the email, like storing it in a database
		return c.SendString(fmt.Sprint(subscriptions.TotalSubscriptions))
	})

	log.Fatal(app.Listen(":4000"))
}
