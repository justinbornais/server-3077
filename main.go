package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/justinbornais/server-3077/handlers"
	"github.com/justinbornais/server-3077/utilities"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./assets")
	utilities.SetupMiddleware(app)
	utilities.InitStore()
	db := utilities.InitDB()
	defer db.Close()

	// Routes for users.
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Post("/login-user", handlers.Login)
	app.Get("/logout", handlers.Logout)

	// Routes for hotel rooms.
	app.Get("/search-rooms", handlers.SearchRooms)

	// Routes for the different pages.
	app.Get("/", handlers.HomePage)
	app.Get("/signup", handlers.SignupPage)
	app.Get("/login", handlers.LoginPage)
	app.Get("/book-room", utilities.IsNotLoggedInMiddleware, handlers.BookingPage)

	// Start server
	log.Fatal(app.Listen(":1450"))
}
