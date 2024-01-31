package routes

import (
	"fmt"

	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check the API
	app.Get("/healthcheck", healthcheckHandler)

	// Register new subsciber
	app.Post("/api/subscribe", subscribeHandler)

	// Tracks unique visits
	// app.Get("/api/track-visit", trackVisitHandler)

}

func healthcheckHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

// Incorporate a faster DB like Redis to do this
// func trackVisitHandler(c *fiber.Ctx) error {
// 	// Check if the visitor has a specific cookie
// 	cookie := c.Cookies("unique_visitor")
// 	if cookie == "" {
// 		// Set a new cookie
// 		newCookie := new(fiber.Cookie)
// 		newCookie.Name = "unique_visitor"
// 		newCookie.Value = "1"
// 		newCookie.Expires = time.Now().Add(24 * time.Hour) // Expires in 1 day
// 		c.Cookie(newCookie)

// 		// subscriptions.TotalVisits = subscriptions.TotalVisits + 1
// 	}

// 	// fmt.Println("Tracked visitors. Total: ", subscriptions.TotalVisits)

// 	// return c.SendString(fmt.Sprint(subscriptions.TotalSubscriptions))
// }

func subscribeHandler(c *fiber.Ctx) error {
	//Todo check for duplicates

	req := new(queries.User)

	// Parse the body into the struct
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Use the email from the request
	fmt.Printf("Received email: %s\n", req.Email)

	// Here you can add logic to handle the email, like storing it in a database
	return c.SendString("Success")
}
