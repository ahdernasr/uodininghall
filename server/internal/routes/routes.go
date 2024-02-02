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

	req := new(queries.Subscriber)

	// Parse the body into the struct
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	fmt.Printf("Received email to subscribe: %s\n", req.Email)

	queries.AddSubscriber(req.Email)

	return c.SendString("Success")
}

func unsubscribeHandler(c *fiber.Ctx) error {
	//Todo check for duplicates

	req := new(queries.Subscriber)

	// Parse the body into the struct
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	fmt.Printf("Received email to unsubscribe: %s\n", req.Email)

	queries.RemoveSubscriber(req.Email)

	return c.SendString("Success")
}
