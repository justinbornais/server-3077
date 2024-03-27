package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/justinbornais/server-3077/utilities"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *fiber.Ctx) error {

	db := utilities.GetDB()

	var user utilities.User
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing the user details")
	}

	log.Printf("User contents:\t%+v\n", user)

	// Hash the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}

	result, err := db.Exec("INSERT INTO users (name,user_type,email,password) VALUES (?, ?, ?, ?)", user.Name, user.UserType, user.Email, string(hashedPassword))
	if err != nil {
		return err
	}

	user.ID, _ = result.LastInsertId()
	user.Password = ""
	// return c.Status(fiber.StatusOK).JSON(user)
	return c.Redirect("/login")
}

func GetUser(c *fiber.Ctx) error {

	db := utilities.GetDB()

	id := c.Params("id")
	row := db.QueryRow("SELECT id,name,user_type,email,password FROM users WHERE id = ?", id)

	var user utilities.User
	err := row.Scan(&user.ID, &user.Name, &user.UserType, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var user, db_user utilities.User
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	db := utilities.GetDB()

	err := db.QueryRow("SELECT id, name, user_type, email, password FROM users WHERE email = ?", user.Email).Scan(&db_user.ID, &db_user.Name, &db_user.UserType, &db_user.Email, &db_user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).SendString("No user found")
	}

	// Compare passwords.
	err = bcrypt.CompareHashAndPassword([]byte(db_user.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	sess, _ := utilities.GetStore().Get(c)
	sess.Set("loggedin", true)
	sess.Set("id", user.ID)
	sess.Set("usertype", user.UserType)
	sess.Save()

	return c.SendString("You are logged in")
}

func Logout(c *fiber.Ctx) error {
	sess, _ := utilities.GetStore().Get(c)
	sess.Destroy()
	return c.SendString("You are logged out")
}
