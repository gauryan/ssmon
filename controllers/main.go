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
		Is_admin string
	}
	var result Result

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	userid := c.FormValue("userid")
	passwd := c.FormValue("passwd")

	db := database.DBConn
	db.Raw("CALL SP_IS_ADMIN(?, ?)", userid, passwd).First(&result)

	if result.Is_admin == "Y" {
		session.Set("ssmon-login", true)
		session.Save()

		return c.Redirect("/mgmt/tcp_server")
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
