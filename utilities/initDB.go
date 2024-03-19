package utilities

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// This will
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.sqlite3")
	PanicError(err, "Error opening the database")

	user_type_stmt := `
		CREATE TABLE IF NOT EXISTS user_type (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)`

	users_stmt := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			user_type INTEGER NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			FOREIGN KEY (user_type) REFERENCES user_type(id)
		)`

	room_type_stmt := `
		CREATE TABLE IF NOT EXISTS room_type (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			beds INTEGER DEFAULT 1,
			beach BOOLEAN NOT NULL,
			price FLOAT NOT NULL
		)`

	rooms_stmt := `
		CREATE TABLE IF NOT EXISTS rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			number INTEGER NOT NULL,
			room_type INTEGER NOT NULL,
			floor INTEGER NOT NULL,
			status TEXT NOT NULL,
			FOREIGN KEY (room_type) REFERENCES room_type(id)
		)`

	bookings_stmt := `
		CREATE TABLE IF NOT EXISTS bookings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			room_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			start DATETIME NOT NULL,
			end DATETIME NOT NULL,
			FOREIGN KEY (room_id) REFERENCES rooms(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`

	_, err = db.Exec(user_type_stmt)
	PanicError(err, "Error creating user_type table")

	err = insertUserType(db, 1, "customer")
	PanicError(err, "Error adding user_type customer")
	err = insertUserType(db, 2, "staff")
	PanicError(err, "Error adding user_type staff")
	err = insertUserType(db, 3, "admin")
	PanicError(err, "Error adding user_type admin")

	_, err = db.Exec(users_stmt)
	PanicError(err, "Error creating users table")

	_, err = db.Exec(room_type_stmt)
	PanicError(err, "Error creating room_type table")

	_, err = db.Exec(rooms_stmt)
	PanicError(err, "Error creating rooms table")

	_, err = db.Exec(bookings_stmt)
	PanicError(err, "Error creating bookings table")

	return db
}

func insertUserType(db *sql.DB, id int, name string) error {
	// Check if the user type already exists
	query := "SELECT id FROM user_type WHERE id = ?"
	var existingID int
	err := db.QueryRow(query, id).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingID != 0 {
		return nil // User type already exists, do nothing.
	}

	// Insert the user type if it doesn't exist
	stmt, err := db.Prepare("INSERT INTO user_type (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name)
	if err != nil {
		return err
	}

	return nil
}
