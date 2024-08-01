package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type Dish struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Allergies   string `json:"allergies"`
}

type Meal struct {
	Breakfast []Dish `json:"breakfast"`
	Lunch     []Dish `json:"lunch"`
	Dinner    []Dish `json:"dinner"`
	Other     []Dish `json:"other"`
}

type Menu struct {
	Grill         Meal `json:"grill"`
	MindBodySoul  Meal `json:"mindBodySoul"`
	PlantBase     Meal `json:"plantBase"`
	ServiceMinute Meal `json:"serviceMinute"`
	Trattoria     Meal `json:"trattoria"`
	WorldFlavours Meal `json:"worldFlavours"`
}

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
func SendMenuEmail(domain, apiKey string, menu Menu) (string, error) {
	// Initialize the Mailgun client
	mg := mailgun.NewMailgun(domain, apiKey)

	// Define the template
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Today's Menu</title>
	</head>
	<body>
		<h1>Today's Menu</h1>
		<h2>Grill</h2>
		<h3>Breakfast</h3>
		{{range .Grill.Breakfast}}
			<p><strong>{{.Name}}</strong>: {{.Description}} (Allergies: {{.Allergies}})</p>
		{{end}}
		<!-- Add similar blocks for Lunch, Dinner, and other meal categories -->
	</body>
	</html>`

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

	menu := Menu{
		Grill: Meal{
			Breakfast: []Dish{
				{Name: "Grilled Sausage", Description: "Savory sausage grilled to perfection.", Allergies: "Contains pork"},
				{Name: "Veggie Patties", Description: "Delicious patties made from fresh vegetables.", Allergies: "Vegetarian"},
			},
			Lunch: []Dish{
				{Name: "BBQ Chicken Sandwich", Description: "Grilled chicken breast with BBQ sauce, lettuce, and tomato.", Allergies: "Contains gluten"},
				{Name: "Steak and Cheese", Description: "Grilled steak slices with melted cheese on a hoagie roll.", Allergies: ""},
			},
			Dinner: []Dish{
				{Name: "Grilled Salmon", Description: "Fresh salmon grilled with herbs and spices.", Allergies: "Contains fish"},
				{Name: "Grilled Portobello Mushrooms", Description: "Marinated portobello mushrooms grilled to perfection.", Allergies: "Vegan"},
			},
			Other: []Dish{
				{Name: "Grilled Corn on the Cob", Description: "Sweet corn grilled with butter.", Allergies: "Vegetarian"},
			},
		},
		MindBodySoul: Meal{
			Breakfast: []Dish{
				{Name: "Fruit Smoothie Bowl", Description: "A blend of fresh fruits topped with granola and honey.", Allergies: "Contains nuts"},
				{Name: "Overnight Oats", Description: "Oats soaked in almond milk with chia seeds and berries.", Allergies: "Vegan"},
			},
			Lunch: []Dish{
				{Name: "Quinoa Salad", Description: "Quinoa mixed with vegetables and a lemon vinaigrette.", Allergies: "Gluten-free"},
				{Name: "Hummus and Veggie Wrap", Description: "Fresh veggies and hummus wrapped in a whole wheat tortilla.", Allergies: "Vegetarian"},
			},
			Dinner: []Dish{
				{Name: "Baked Tofu Stir-fry", Description: "Tofu stir-fried with mixed vegetables in a soy-ginger sauce.", Allergies: "Vegan"},
			},
			Other: []Dish{
				{Name: "Energy Bars", Description: "Homemade energy bars with oats, nuts, and dried fruit.", Allergies: "Contains nuts"},
			},
		},
		PlantBase: Meal{
			Breakfast: []Dish{
				{Name: "Avocado Toast", Description: "Whole grain toast topped with mashed avocado and cherry tomatoes.", Allergies: "Vegan"},
				{Name: "Chia Seed Pudding", Description: "Chia seeds soaked in coconut milk, topped with fresh fruit.", Allergies: "Vegan"},
			},
			Lunch: []Dish{
				{Name: "Vegan Buddha Bowl", Description: "A mix of quinoa, chickpeas, avocado, and roasted vegetables.", Allergies: "Vegan"},
			},
			Dinner: []Dish{
				{Name: "Stuffed Bell Peppers", Description: "Bell peppers stuffed with a mixture of rice, beans, and corn.", Allergies: "Vegan"},
				{Name: "Lentil Soup", Description: "Hearty soup made with lentils and a variety of vegetables.", Allergies: "Vegan"},
			},
			Other: []Dish{
				{Name: "Fruit Salad", Description: "A mix of seasonal fruits, served fresh.", Allergies: "Vegan"},
			},
		},
		ServiceMinute: Meal{
			Breakfast: []Dish{
				{Name: "Scrambled Eggs", Description: "Fluffy scrambled eggs, served with toast.", Allergies: ""},
			},
			Lunch: []Dish{
				{Name: "Chicken Caesar Salad", Description: "Grilled chicken breast with romaine lettuce and Caesar dressing.", Allergies: "Contains gluten"},
			},
			Dinner: []Dish{
				{Name: "Spaghetti Bolognese", Description: "Classic spaghetti with a rich meat sauce.", Allergies: "Contains gluten"},
			},
			Other: []Dish{
				{Name: "Garlic Bread", Description: "Toasted bread with garlic butter.", Allergies: "Contains gluten"},
			},
		},
		Trattoria: Meal{
			Breakfast: []Dish{
				{Name: "Breakfast Pizza", Description: "Pizza topped with scrambled eggs, bacon, and cheese.", Allergies: "Contains gluten"},
			},
			Lunch: []Dish{
				{Name: "Margherita Pizza", Description: "Classic pizza with tomatoes, mozzarella, and basil.", Allergies: "Contains gluten"},
			},
			Dinner: []Dish{
				{Name: "Fettuccine Alfredo", Description: "Fettuccine pasta in a creamy Alfredo sauce.", Allergies: "Contains gluten"},
				{Name: "Pasta Primavera", Description: "Pasta with fresh vegetables in a light garlic sauce.", Allergies: "Contains gluten"},
			},
			Other: []Dish{
				{Name: "Tiramisu", Description: "Traditional Italian dessert with layers of coffee-soaked ladyfingers and mascarpone cream.", Allergies: "Contains gluten"},
			},
		},
		WorldFlavours: Meal{
			Breakfast: []Dish{
				{Name: "Shakshuka", Description: "Poached eggs in a spicy tomato sauce.", Allergies: "Vegetarian"},
			},
			Lunch: []Dish{
				{Name: "Chicken Tikka Masala", Description: "Grilled chicken in a creamy spiced tomato sauce.", Allergies: "Halal"},
			},
			Dinner: []Dish{
				{Name: "Beef Teriyaki", Description: "Beef strips in a sweet and savory teriyaki sauce.", Allergies: "Contains soy"},
				{Name: "Falafel Platter", Description: "Crispy falafel served with hummus, pita, and salad.", Allergies: "Vegan"},
			},
			Other: []Dish{
				{Name: "Mango Sticky Rice", Description: "Sweet sticky rice with mango slices and coconut sauce.", Allergies: "Vegan"},
			},
		},
	}

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
