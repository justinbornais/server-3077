package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/justinbornais/server-3077/handlers"
	"github.com/justinbornais/server-3077/utilities"
)

var store = session.New()

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./assets")
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("base", fiber.Map{
			"ContentTemplate": "home",
		})
	})

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

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
}
