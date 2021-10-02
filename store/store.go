package store

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore *session.Store


func Init() {
	SessionStore = session.New()
}
