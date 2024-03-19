package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/justinbornais/server-3077/handlers"
	"github.com/justinbornais/server-3077/utilities"
)

func main() {

	app := fiber.New()
	setupMiddleware(app)
	db := utilities.InitDB()
	defer db.Close()

	getUser := func(c *fiber.Ctx) error {
		return handlers.GetUser(c, db)
	}

	addUser := func(c *fiber.Ctx) error {
		return handlers.AddUser(c, db)
	}

	// Routes
	app.Get("/users/:id", getUser)
	app.Post("/users", addUser)

	// Start server
	log.Fatal(app.Listen(":1450"))
}

func setupMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	// app.Use(func(c *fiber.Ctx) error {
	// 	log.Printf("%s %s - %d", c.Method(), c.OriginalURL(), c.Response().StatusCode())
	// 	return c.Next()
	// })

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
}
