package mgmt

// controllers/mgmt

import (
	"fmt"
	"strconv"

	"github.com/gauryan/ssmon/database"
	"github.com/gauryan/ssmon/store"
	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	Id      int
	Enabled int
	Name    string
	Url     string
	Chk_str string
	Timeout int
	Err_cnt int
}

// HTTP Server 모니터링
// /mgmt/http_monitor
func MonitorHTTPServer(c *fiber.Ctx) error {
	type Result struct {
		Value string
	}
	var result Result
	var http_servers []HttpServer

	db := database.DBConn
	db.Raw("CALL SP_MONITOR_HTTPSERVER()").Scan(&http_servers)
	db.Raw("CALL SP_GET_ERR_CNT_FOR_ALARM()").First(&result)
	err_cnt_for_alarm, _ := strconv.Atoi(result.Value)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	data := fiber.Map{"Httpservers": http_servers, "Errcnt4alarm": err_cnt_for_alarm, "Menu": "http_monitor"}
	return c.Render("mgmt/http_server/monitor", data, "base")
}

// HTTP Server 목록
// /mgmt/http_server
func ListHTTPServer(c *fiber.Ctx) error {
	var http_servers []HttpServer
	var flash string

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Raw("CALL SP_LIST_HTTPSERVER()").Scan(&http_servers)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	if session.Get("flash") != nil {
		flash = fmt.Sprintf("%v", session.Get("flash"))
		session.Delete("flash")
		session.Save()
	}

	data := fiber.Map{"Httpservers": http_servers, "Flash": flash, "Menu": "http_server"}
	return c.Render("mgmt/http_server/index", data, "base")
}

// HTTP Server 추가 폼
// /mgmt/http_server/insert_form
func InsertFormHTTPServer(c *fiber.Ctx) error {
	return c.Render("mgmt/http_server/insert_form", fiber.Map{})
}

// HTTP Server 추가
// /mgmt/http_server/insert
func InsertHTTPServer(c *fiber.Ctx) error {
	name := c.FormValue("name")
	url := c.FormValue("url")
	chk_str := c.FormValue("chk_str")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_INSERT_HTTPSERVER(?, ?, ?, ?)", name, url, chk_str, timeout)

	session.Set("flash", "새로운 HTTP Server("+name+")이 추가되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/http_server")
}

// HTTP Server 수정 폼
// /mgmt/http_server/update_form/{id}
func UpdateFormHTTPServer(c *fiber.Ctx) error {
	var http_server HttpServer

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_HTTPSERVER(?)", id).First(&http_server)
	data := fiber.Map{"HTTPServer": http_server}
	return c.Render("mgmt/http_server/update_form", data)
}

// HTTP Server 수정
// /mgmt/http_server/update
func UpdateHTTPServer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	url := c.FormValue("url")
	chk_str := c.FormValue("chk_str")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_UPDATE_HTTPSERVER(?, ?, ?, ?, ?)", id, name, url, chk_str, timeout)

	session.Set("flash", "HTTP서버("+name+")가 수정되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/http_server")
}

// HTTP Server Enabled 상태 토글하기
// /mgmt/http_server/toggle_enabled/{id}
func ToggleEnabledHTTPServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_TOGGLE_ENABLED_HTTPSERVER(?)", id)

	session.Set("flash", "HTTP서버상태가 토글되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/http_server")
}

// HTTP Server 삭제
// /mgmt/http_server/delete/{id}
func DeleteHTTPServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_DELETE_HTTPSERVER(?)", id)

	session.Set("flash", "HTTP서버가 삭제되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/http_server")
}
