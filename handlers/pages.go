package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/justinbornais/server-3077/utilities"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("base", fiber.Map{
		"ActivePage": "home",
	})
}

func SignupPage(c *fiber.Ctx) error {
	userTypes, err := GetUserTypes()
	if err != nil {
		log.Println(err)
		return c.Render("signup", fiber.Map{
			"ActivePage": nil,
			"UserTypes":  nil,
		})
	}
	return c.Render("signup", fiber.Map{
		"ActivePage": nil,
		"UserTypes":  userTypes,
	})
}

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func BookingPage(c *fiber.Ctx) error {
	floors, err := GetFloorNumbers()
	roomTypes, err2 := GetRoomTypes()
	if err != nil || err2 != nil {
		log.Println(err)
		return c.Render("booking", fiber.Map{
			"ActivePage":   "booking",
			"FloorNumbers": nil,
			"RoomTypes":    nil,
		})
	}
	return c.Render("booking", fiber.Map{
		"ActivePage":   "booking",
		"FloorNumbers": floors,
		"RoomTypes":    roomTypes,
	})
}

func BookingRoomsPage(c *fiber.Ctx, rooms []utilities.Room) error {
	floors, err := GetFloorNumbers()
	if err != nil {
		log.Println(err)
		return c.Render("booking", fiber.Map{
			"ActivePage":   "booking",
			"FloorNumbers": nil,
			"RoomNumbers":  nil,
		})
	}
	return c.Render("booking", fiber.Map{
		"ActivePage":   "booking",
		"FloorNumbers": floors,
		"RoomNumbers":  rooms,
	})
}
