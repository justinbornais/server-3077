package handlers

import (
	"github.com/justinbornais/server-3077/utilities"
)

func GetRoomTypes() ([]utilities.RoomType, error) {

	db := utilities.GetDB()

	rows, err := db.Query("SELECT id,name,beds,beach,price FROM room_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomTypes []utilities.RoomType
	for rows.Next() {
		var rt utilities.RoomType
		if err := rows.Scan(&rt.ID, &rt.Name, &rt.Beds, &rt.Beach, &rt.Price); err != nil {
			return nil, err
		}
		roomTypes = append(roomTypes, rt)
	}
	return roomTypes, nil
}
