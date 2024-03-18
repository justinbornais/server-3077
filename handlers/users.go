package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinbornais/server-3077/utilities"
)

func GetUser(c *gin.Context, db *sql.DB) {

	var user utilities.User
	id := c.Param("id")

	row := db.QueryRow("SELECT id, name, user_type, email, password FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.UserType, &user.Email, &user.Password)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context, db *sql.DB) {

	var user utilities.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("INSERT INTO users(name, user_type, email, password) VALUES(?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(user.Name, user.UserType, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
