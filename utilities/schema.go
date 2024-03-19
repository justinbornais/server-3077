package utilities

type UserType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RoomType struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Beds  int     `json:"beds"`
	Beach bool    `json:"beach"`
	Price float64 `json:"price"`
}

type Room struct {
	ID       int64  `json:"id"`
	Number   int    `json:"number"`
	RoomType int    `json:"room_type"`
	Floor    int    `json:"floor"`
	Status   string `json:"status"`
}

type Booking struct {
	ID     int64  `json:"id"`
	RoomID int    `json:"room_id"`
	UserID int    `json:"user_id"`
	Start  string `json:"start"`
	End    string `json:"end"`
}
