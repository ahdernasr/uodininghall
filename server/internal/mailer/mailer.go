package mailer

import (
	"fmt"
	"log"

	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// SendSimpleMessage sends a simple email using the Mailgun API
func SendSimpleMessage(domain, apiKey, publicApiKey string) (string, error) {
	// Initialize the Mailgun client
	mg := mailgun.NewMailgun(domain, apiKey)

	// Create a new email message
	m := mg.NewMessage(
		"Excited User <mailgun@sandboxe2100e417b624fa09860f37548c39173.mailgun.org>", // Sender's email
		"Hello",                             // Subject
		"Testing some Mailgun awesomeness!", // Plain text body
		"ahdernasr@gmail.com",               // Recipient's email
	)

	// Set a 10-second timeout context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message
	_, id, err := mg.Send(ctx, m)

	// Return the message ID and any error
	return id, err
}

func main() {
	// Replace these with your actual Mailgun domain and API key
	domain := "sandboxe2100e417b624fa09860f37548c39173.mailgun.org"
	apiKey := "9a24d4ff262d0431c070f3f82d43f9d2-0f1db83d-333ea510"
	domain := ""
	apiKey := ""
	publicApiKey := "" // Not needed for sending emails

	// Send an email
	messageID, err := SendSimpleMessage(domain, apiKey, publicApiKey)
	if err != nil {
		log.Fatalf("Could not send email: %v", err)
	}
	fmt.Printf("Email sent successfully! Message ID: %s\n", messageID)
}
