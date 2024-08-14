package mailer

import (
	"context"
	"fmt"
	"html/template"
	"strings"
	"sync"
	"time"

	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/ahdernasr/dailydininghall/internal/scraper"
	"github.com/mailgun/mailgun-go/v4"
)

func SendMenuEmail(domain, apiKey string, menu *scraper.Menu, subscribers []queries.Subscriber) error {
	// Initialize the Mailgun client
	mg := mailgun.NewMailgun(domain, apiKey)

	// Define the template
	tmpl := getMenuTemplate()

	// Parse the template
	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Create a string builder to capture the output
	var bodyBuilder strings.Builder
	err = t.Execute(&bodyBuilder, menu)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	// Get the final email body from the builder
	emailBody := bodyBuilder.String()

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to receive errors from goroutines
	errChan := make(chan error, len(subscribers))

	// Iterate over the list of subscribers and send the email concurrently
	for _, subscriber := range subscribers {
		wg.Add(1)
		go func(subscriber queries.Subscriber) {
			defer wg.Done()

			m := mg.NewMessage(
				"School Cafeteria <mailgun@sandbox314528bf85614e73b0a63061fb8c323a.mailgun.org>", // Sender's email
				"Today's Menu", // Subject
				"Hello, please view this email in HTML format.", // Plain-text body
				subscriber.Email, // Recipient's email
			)

			m.SetHtml(emailBody)

			// Set a 30-second timeout context
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			// Send the message
			_, _, err := mg.Send(ctx, m)
			if err != nil {
				errChan <- fmt.Errorf("failed to send email to %s: %v", subscriber.Email, err)
			}
		}(subscriber)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}

func SendSubscribeEmail(domain, apiKey string, menu *scraper.Menu) (string, error) {
	// Initialize the Mailgun client
	mg := mailgun.NewMailgun(domain, apiKey)

	// Define the template
	tmpl := GetSubscribeTemplate()

	// Parse the template
	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	// Create a string builder to capture the output
	var bodyBuilder strings.Builder
	err = t.Execute(&bodyBuilder, menu)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	// Get the final email body from the builder
	emailBody := bodyBuilder.String()

	// Create a new email message with the custom body
	m := mg.NewMessage(
		"School Cafeteria <mailgun@sandbox314528bf85614e73b0a63061fb8c323a.mailgun.org>", // Sender's email
		"Today's Menu", // Subject
		"Hello, please view this email in HTML format.", // HTML body
		"ahdernasr@gmail.com",                           // Recipient's email
	)

	m.SetHtml(emailBody)

	// Set a 30-second timeout context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Send the message
	_, id, err := mg.Send(ctx, m)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %v", err)
	}

	return id, nil
}
