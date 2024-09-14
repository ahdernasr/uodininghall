package routes

import (
	"fmt"

	"github.com/ahdernasr/dailydininghall/internal/db/queries"
	"github.com/ahdernasr/dailydininghall/internal/mailer"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", healthcheckHandler)

	app.Post("/api/subscribe", subscribeHandler)
	app.Post("/api/unsubscribe", unsubscribeHandler)
}

func healthcheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func subscribeHandler(c *fiber.Ctx) error {
	//Todo check for duplicates

	req := new(queries.Subscriber)

	// Parse the body into the struct
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	fmt.Printf("Received email to subscribe: %s\n", req.Email)

	err := queries.AddSubscriber(req.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Email is already subscribed.")
	}

	mailError := mailer.SendSubscribeEmail(req.Email)

	if mailError != nil {
		fmt.Println("Error sending email to: ", req.Email)
	}

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
