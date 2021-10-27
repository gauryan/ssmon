package mgmt

// controllers/mgmt

import (
	"fmt"
	"strconv"

	"github.com/gauryan/ssmon/database"
	"github.com/gauryan/ssmon/store"
	"github.com/gofiber/fiber/v2"
)

type PingServer struct {
	Id      int
	Enabled int
	Name    string
	Ip_addr string
	Timeout int
	Err_cnt int
}

// PING Server 모니터링
// /mgmt/ping_monitor
func MonitorPINGServer(c *fiber.Ctx) error {
	type Result struct {
		Value string
	}
	var result Result
	var ping_servers []PingServer

	db := database.DBConn
	db.Raw("CALL SP_MONITOR_PINGSERVER()").Scan(&ping_servers)
	db.Raw("CALL SP_GET_ERR_CNT_FOR_ALARM()").First(&result)
	err_cnt_for_alarm, _ := strconv.Atoi(result.Value)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	data := fiber.Map{"Pingservers": ping_servers, "Errcnt4alarm": err_cnt_for_alarm, "Menu": "ping_monitor"}
	return c.Render("mgmt/ping_server/monitor", data, "base")
}

// PING Server 목록
// /mgmt/ping_server
func ListPINGServer(c *fiber.Ctx) error {
	var ping_servers []PingServer
	var flash string

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Raw("CALL SP_LIST_PINGSERVER()").Scan(&ping_servers)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	if session.Get("flash") != nil {
		flash = fmt.Sprintf("%v", session.Get("flash"))
		session.Delete("flash")
		session.Save()
	}

	data := fiber.Map{"Pingservers": ping_servers, "Flash": flash, "Menu": "ping_server"}
	return c.Render("mgmt/ping_server/index", data, "base")
}

// PING Server 추가 폼
// /mgmt/ping_server/insert_form
func InsertFormPINGServer(c *fiber.Ctx) error {
	return c.Render("mgmt/ping_server/insert_form", fiber.Map{})
}

// PING Server 추가
// /mgmt/ping_server/insert
func InsertPINGServer(c *fiber.Ctx) error {
	name := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_INSERT_PINGSERVER(?, ?, ?)", name, ip_addr, timeout)

	session.Set("flash", "새로운 PING Server("+name+")이 추가되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/ping_server")
}

// PING Server 수정 폼
// /mgmt/ping_server/update_form/{id}
func UpdateFormPINGServer(c *fiber.Ctx) error {
	var ping_server PingServer

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_PINGSERVER(?)", id).First(&ping_server)
	data := fiber.Map{"PINGServer": ping_server}
	return c.Render("mgmt/ping_server/update_form", data)
}

// PING Server 수정
// /mgmt/ping_server/update
func UpdatePINGServer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_UPDATE_PINGSERVER(?, ?, ?, ?)", id, name, ip_addr, timeout)

	session.Set("flash", "PING서버("+name+")가 수정되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/ping_server")
}

// PING Server Enabled 상태 토글하기
// /mgmt/ping_server/toggle_enabled/{id}
func ToggleEnabledPINGServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_TOGGLE_ENABLED_PINGSERVER(?)", id)

	session.Set("flash", "PING서버상태가 토글되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/ping_server")
}

// PING Server 삭제
// /mgmt/ping_server/delete/{id}
func DeletePINGServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_DELETE_PINGSERVER(?)", id)

	session.Set("flash", "PING서버가 삭제되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/ping_server")
}
