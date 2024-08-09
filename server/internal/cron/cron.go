package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Your task function
func fetchScrapeAndEmail() {
	// Your code to fetch email addresses, scrape the menu, and send emails
	log.Println("Fetching email addresses, scraping the menu, and sending emails...")
	// Example function calls
	fetchEmailAddresses()
	scrapeMenu()
	sendEmails()
}

func fetchEmailAddresses() {
	// Placeholder function to fetch email addresses from the database
	log.Println("Fetching email addresses from the database...")
}

func scrapeMenu() {
	// Placeholder function to scrape the menu
	log.Println("Scraping the menu...")
}

func sendEmails() {
	// Placeholder function to send emails
	log.Println("Sending emails...")
}

func main() {
	// Define the timezone
	location, err := time.LoadLocation("Africa/Cairo") // Use the appropriate timezone
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	// Create a new cron instance with the specified timezone
	c := cron.New(cron.WithLocation(location))

	// Schedule the task to run at 6 AM every day in the specified timezone
	_, err = c.AddFunc("0 20 * * *", fetchScrapeAndEmail)
	if err != nil {
		log.Fatal("Failed to schedule the task: ", err)
	}

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}
