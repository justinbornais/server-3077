package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/justinbornais/server-3077/utilities"
)

func AddUser(c *fiber.Ctx, db *sql.DB) error {

	log.Println("Adding user....", string(c.Body()))
	var user utilities.User
	if err := c.BodyParser(&user); err != nil {
		log.Printf("Partially parsed contents:\t%+v\n", user)
		log.Println(err)
		return err
	}

	log.Printf("User contents:\t%+v\n", user)

	result, err := db.Exec("INSERT INTO users (name,user_type,email,password) VALUES (?, ?, ?, ?)", user.Name, user.UserType, user.Email, user.Password)
	if err != nil {
		return err
	}

	user.ID, _ = result.LastInsertId()
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")

	row := db.QueryRow("SELECT id,name,user_type,email,password FROM users WHERE id = ?", id)

	var user utilities.User
	err := row.Scan(&user.ID, &user.Name, &user.UserType, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
