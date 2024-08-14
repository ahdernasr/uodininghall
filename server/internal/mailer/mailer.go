package mailer

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/ahdernasr/dailydininghall/internal/scraper"
	"github.com/mailgun/mailgun-go/v4"
)

func SendMenuEmail(domain, apiKey string, menu *scraper.Menu) (string, error) {
	// Initialize the Mailgun client
	mg := mailgun.NewMailgun(domain, apiKey)

	// Define the template
	tmpl := getMenuTemplate()

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

func main() {

	menu := scraper.Scraper()

	// Keys
	domain := "sandbox314528bf85614e73b0a63061fb8c323a.mailgun.org"
	apiKey := "b217988e98c92f971cfff1432c105353-afce6020-b5e3a061"

	// Send an email

	messageID, err := SendMenuEmail(domain, apiKey, menu)
	if err != nil {
		log.Fatalf("Could not send email: %v", err)
	}
	fmt.Printf("Email sent successfully! Message ID: %s\n", messageID)
}
