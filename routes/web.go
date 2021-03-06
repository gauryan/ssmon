package routes

import (
	"github.com/gauryan/ssmon/controllers"
	"github.com/gauryan/ssmon/controllers/mgmt"
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

	ssmon_login := session.Get("ssmon-login")
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
	App1.Get("/logout", controllers.Logout)

	App2 := app.Group("/mgmt", authSSMON) // 로그인후에만 접근가능
	// App2 := app.Group("/mgmt")

	App2.Get("/errlog/:page", mgmt.ListErrRecLog)

	App2.Get("/setting", mgmt.ViewSetting)
	App2.Post("/setting/update", mgmt.UpdateSetting)

	App2.Get("/admin", mgmt.ListAdmin)
	App2.Get("/admin/insert_form", mgmt.InsertFormAdmin)
	App2.Post("/admin/insert", mgmt.InsertAdmin)
	App2.Get("/admin/chg_passwd_form/:id", mgmt.ChgPasswdFormAdmin)
	App2.Post("/admin/chg_passwd", mgmt.ChgPasswdAdmin)
	App2.Get("/admin/update_form/:id", mgmt.UpdateFormAdmin)
	App2.Post("/admin/update", mgmt.UpdateAdmin)
	App2.Get("/admin/delete/:id", mgmt.DeleteAdmin)

	App2.Get("/tcp_monitor", mgmt.MonitorTCPServer)
	App2.Get("/tcp_server", mgmt.ListTCPServer)
	App2.Get("/tcp_server/insert_form", mgmt.InsertFormTCPServer)
	App2.Post("/tcp_server/insert", mgmt.InsertTCPServer)
	App2.Get("/tcp_server/update_form/:id", mgmt.UpdateFormTCPServer)
	App2.Post("/tcp_server/update", mgmt.UpdateTCPServer)
	App2.Get("/tcp_server/toggle_enabled/:id", mgmt.ToggleEnabledTCPServer)
	App2.Get("/tcp_server/delete/:id", mgmt.DeleteTCPServer)

	App2.Get("/http_monitor", mgmt.MonitorHTTPServer)
	App2.Get("/http_server", mgmt.ListHTTPServer)
	App2.Get("/http_server/insert_form", mgmt.InsertFormHTTPServer)
	App2.Post("/http_server/insert", mgmt.InsertHTTPServer)
	App2.Get("/http_server/update_form/:id", mgmt.UpdateFormHTTPServer)
	App2.Post("/http_server/update", mgmt.UpdateHTTPServer)
	App2.Get("/http_server/toggle_enabled/:id", mgmt.ToggleEnabledHTTPServer)
	App2.Get("/http_server/delete/:id", mgmt.DeleteHTTPServer)

	App2.Get("/ping_monitor", mgmt.MonitorPINGServer)
	App2.Get("/ping_server", mgmt.ListPINGServer)
	App2.Get("/ping_server/insert_form", mgmt.InsertFormPINGServer)
	App2.Post("/ping_server/insert", mgmt.InsertPINGServer)
	App2.Get("/ping_server/update_form/:id", mgmt.UpdateFormPINGServer)
	App2.Post("/ping_server/update", mgmt.UpdatePINGServer)
	App2.Get("/ping_server/toggle_enabled/:id", mgmt.ToggleEnabledPINGServer)
	App2.Get("/ping_server/delete/:id", mgmt.DeletePINGServer)

	return app
}
