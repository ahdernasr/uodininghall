package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahdernasr/dailydininghall/internal/db"
	"github.com/ahdernasr/dailydininghall/internal/routes"
	"github.com/robfig/cron/v3"

	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/ahdernasr/dailydininghall/internal/mailer"
	"github.com/ahdernasr/dailydininghall/internal/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	/* SERVER */

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Load connection string frome .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded.")
	}

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

	// Keys
	domain := "sandbox314528bf85614e73b0a63061fb8c323a.mailgun.org"
	apiKey := "b217988e98c92f971cfff1432c105353-afce6020-b5e3a061"

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
