package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahdernasr/dailydininghall/internal/db"
	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/ahdernasr/dailydininghall/internal/mailer"
	"github.com/ahdernasr/dailydininghall/internal/routes"
	"github.com/ahdernasr/dailydininghall/internal/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	// "github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {

	// ADD THIS IF IN DEVELOPMENT
	// Load connection string frome .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(".env file could not be loaded.", err)
	// }

	/* SERVER */

	app := fiber.New()

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_URL"),
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Rate limiter
	app.Use(limiter.New(limiter.Config{
		Max:        10,               // Max requests per duration
		Expiration: 30 * time.Second, // Duration before the rate limit resets
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP address
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		},
	}))

	// Setup routes
	routes.SetupRoutes(app)

	connectionString := os.Getenv("CONNECTION_STRING")

	// Connect to db
	if err := db.Connect(connectionString); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} else {
		fmt.Println("Connected to db!")
	}

	/* CRON SCHEDULER */

	location, err := time.LoadLocation("Africa/Cairo") // Use the appropriate timezone
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	// Create a new cron instance with the specified timezone
	c := cron.New(cron.WithLocation(location))

	// Schedule the task to run at 6 AM every day in the specified timezone
	_, err = c.AddFunc("0 6 * * *", daily)
	if err != nil {
		log.Fatal("Failed to schedule the task: ", err)
	}

	// Start the cron scheduler
	c.Start()

	log.Fatal(app.Listen(":4000"))
}

func daily() {

	// ADD THIS IF IN DEVELOPMENT

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(".env file could not be loaded.", err)
	// }

	domain := os.Getenv("EMAIL_DOMAIN")
	apiKey := os.Getenv("EMAIL_API_KEY")

	// Run the scraper to get the menu
	/* TODO Add error checking here */
	menu := scraper.Scraper()

	// Get the mailing list
	subscribers, err1 := queries.GetAllSubscribers()
	if err1 != nil {
		log.Fatalf("Could not get all subscribers: %v", err1)
	}

	// Send the menu to all subscribers
	err2 := mailer.SendMenuEmail(domain, apiKey, menu, subscribers)
	if err2 != nil {
		log.Fatalf("Could not send email: %v", err2)
	}

	fmt.Printf("Emails sent successfully!")

}
