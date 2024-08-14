package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahdernasr/dailydininghall/internal/db"
	"github.com/ahdernasr/dailydininghall/internal/routes"
	"github.com/robfig/cron/v3"

	// "github.com/ahdernasr/dailydininghall/internal/scraper"
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
	_, err = c.AddFunc("0 6 * * *", test)
	if err != nil {
		log.Fatal("Failed to schedule the task: ", err)
	}

	// Start the cron scheduler
	c.Start()

	log.Fatal(app.Listen(":4000"))
}

func test() {
	fmt.Println("running cron")
}
