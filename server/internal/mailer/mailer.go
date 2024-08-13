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

// SendSimpleMessage sends an email using a Mailgun template and variables
// func SendMenuEmail(domain, apiKey string, menu Menu) (string, error) {
// 	// Initialize the Mailgun client
// 	mg := mailgun.NewMailgun(domain, apiKey)

// 	// Create a new email message with a template
// 	m := mg.NewMessage(
// 		"School Cafeteria <mailgun@sandbox314528bf85614e73b0a63061fb8c323a.mailgun.org>", // Sender's email
// 		"Today's Menu",        // Subject
// 		"",                    // Plain text body (not used if template body is defined)
// 		"ahdernasr@gmail.com", // Recipient's email
// 	)

// 	// Set the template to use
// 	m.SetTemplate("menu") // Replace 'menu_template' with the actual template name in Mailgun

// 	// Serialize the menu struct into JSON
// 	// menuData, err := json.Marshal(menu)
// 	// if err != nil {
// 	// 	return "", fmt.Errorf("failed to marshal menu: %v", err)
// 	// }

// 	// fmt.Printf(string(menuData))
// 	// m.AddHeader("X-Mailgun-Variables", string(menuData))

// 	grillBreakfast := []map[string]string{
// 		{"name": "Grilled Sausage", "description": "Savory sausage grilled to perfection.", "allergies": "Contains pork"},
// 		{"name": "Veggie Patties", "description": "Delicious patties made from fresh vegetables.", "allergies": "Vegetarian"},
// 	}

// 	grillBreakfastJson, err := json.Marshal(grillBreakfast)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set a 30-second timeout context
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	// Send the message
// 	_, id, err := mg.Send(ctx, m)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return id, nil
// }

// SendSimpleMessage sends an email using a Mailgun template and variables
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
