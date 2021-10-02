package controllers

// controllers

import (
	"github.com/gauryan/ssmon/store"
	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"
)

// Login 화면
func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

// 로그인
func Login(c *fiber.Ctx) error {
	type Result struct {
		IsMember int
	}
	var result Result

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	userid := c.FormValue("userid")
	passwd := c.FormValue("passwd")

	db := database.DBConn
	db.Raw("SELECT isMember(?, ?) as is_member", userid, passwd).First(&result)

	if result.IsMember == 1 {
		session.Set("ssmon-login", true)
		session.Save()

		return c.Redirect("/admin")
	}
	return c.Redirect("/")
}


// 로그아웃
func Logout (c *fiber.Ctx) error {
	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}
	session.Destroy()
	return c.Redirect("/")
}
