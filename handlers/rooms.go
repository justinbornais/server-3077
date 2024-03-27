package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/justinbornais/server-3077/utilities"
)

type RoomSearch struct {
	Floor     int    `form:"floor" json:"floor"`
	RoomType  int    `form:"room_type" json:"room_type"`
	StartDate string `form:"start_date" json:"start_date"`
	EndDate   string `form:"end_date" json:"end_date"`
}

func GetFloorNumbers() ([]int, error) {

	db := utilities.GetDB()

	rows, err := db.Query("SELECT DISTINCT floor FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var floors []int
	for rows.Next() {
		var floor int
		if err := rows.Scan(&floor); err != nil {
			return nil, err
		}
		floors = append(floors, floor)
	}
	return floors, nil
}

func SearchRooms(c *fiber.Ctx) error {

	db := utilities.GetDB()
	var search RoomSearch
	search.Floor = c.QueryInt("floor", -1)
	search.RoomType = c.QueryInt("room_type", -1)
	search.StartDate = c.Query("start_date", "no-data")
	search.EndDate = c.Query("end_date", "no-data")

	if search.Floor == -1 || search.RoomType == -1 || search.StartDate == "no-data" || search.EndDate == "no-data" {
		log.Println(search)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	rows, err := db.Query(`SELECT rooms.id, rooms.number, rooms.room_type, rooms.floor, rooms.status
			FROM rooms
			LEFT JOIN bookings ON rooms.id = bookings.room_id 
				AND NOT (
					bookings.end <= ? OR 
					bookings.start >= ?
				)
			WHERE rooms.floor=?
			AND rooms.room_type=?
			AND rooms.status="available"
			AND bookings.id IS NULL
			GROUP BY rooms.id;`,
		search.EndDate, search.StartDate, search.Floor, search.RoomType)
	if err != nil {
		log.Println(err)
		return c.Redirect("/book-room", fiber.StatusNotFound)
	}
	defer rows.Close()

	var rooms []utilities.Room
	for rows.Next() {
		var room utilities.Room
		if err := rows.Scan(&room.ID, &room.Number, &room.RoomType, &room.Floor, &room.Status); err != nil {
			return c.Redirect("/book-room", fiber.StatusNotFound)
		}
		rooms = append(rooms, room)
	}

	return BookingRoomsPage(c, rooms)
}
