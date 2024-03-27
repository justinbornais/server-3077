package utilities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddleware(app *fiber.App) {
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
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
}

// Checks if the user is logged in by looking at the session.
func IsLoggedInMiddleware(c *fiber.Ctx) error {
	sess, err := GetStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if loggedIn, ok := sess.Get("loggedin").(bool); ok || loggedIn {
		return c.Redirect("/")
	}

	return c.Next()
}

func IsNotLoggedInMiddleware(c *fiber.Ctx) error {
	sess, err := GetStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if loggedIn, ok := sess.Get("loggedin").(bool); !ok || !loggedIn {
		return c.Redirect("/login")
	}

	return c.Next()
}
