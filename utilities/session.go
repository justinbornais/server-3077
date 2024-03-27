package utilities

import "github.com/gofiber/fiber/v2/middleware/session"

var store *session.Store

func InitStore() {
	store = session.New()
}

func GetStore() *session.Store {
	return store
}
