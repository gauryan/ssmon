package routes

import (
	"github.com/gauryan/ssmon/controllers"
	"github.com/gauryan/ssmon/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// authSSMON 미들웨어
func authSSMON(c *fiber.Ctx) error {
	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	ssmom_login := session.Get("ssmon-login")
	if ssmon_login != true {
		return c.Redirect("/")
	}

	return c.Next()
}

func Router() *fiber.App {
	// App 생성과 템플릿 설정
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Route 설정
	App1 := app.Group("/") // 로그인전 접근가능
	App1.Get("/", controllers.Index)
	App1.Post("/login", controllers.Login)

	App2 := app.Group("/", authSSMON) // 로그인후에만 접근가능
	App2.Get("/logout", controllers.Logout)
	// App2.Get("/admin", controllers.ListAdmin)
	// App2.Get("/admin/insert_form", controllers.InsertForm)
	// App2.Post("/admin/insert", controllers.Insert)
	// App2.Get("/admin/chg_passwd_form/:id", controllers.ChgPasswdForm)
	// App2.Post("/admin/chg_passwd", controllers.ChgPasswd)
	// App2.Get("/admin/update_form/:id", controllers.UpdateForm)
	// App2.Post("/admin/update", controllers.Update)
	// App2.Get("/admin/delete/:id", controllers.Delete)

	return app
}
