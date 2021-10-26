package mgmt

// controllers/mgmt

import (
	"fmt"
	"strconv"

	"github.com/gauryan/ssmon/database"
	"github.com/gauryan/ssmon/store"
	"github.com/gofiber/fiber/v2"
)

type TcpServer struct {
	Id      int
	Enabled int
	Name    string
	Ip_addr string
	Port    int
	Timeout int
	Err_cnt int
}

// TCP Server 모니터링
// /mgmt/tcp_monitor
func MonitorTCPServer(c *fiber.Ctx) error {
	type Result struct {
		Value string
	}
	var result Result
	var tcp_servers []TcpServer

	db := database.DBConn
	db.Raw("CALL SP_MONITOR_TCPSERVER()").Scan(&tcp_servers)
	db.Raw("CALL SP_GET_ERR_CNT_FOR_ALARM()").First(&result)
	err_cnt_for_alarm, _ := strconv.Atoi(result.Value)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	data := fiber.Map{"Tcpservers": tcp_servers, "Errcnt4alarm": err_cnt_for_alarm, "Menu": "tcp_monitor"}
	return c.Render("mgmt/tcp_server/monitor", data, "base")
}

// TCP Server 목록
// /mgmt/tcp_server
func ListTCPServer(c *fiber.Ctx) error {
	var tcp_servers []TcpServer
	var flash string

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Raw("CALL SP_LIST_TCPSERVER()").Scan(&tcp_servers)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	if session.Get("flash") != nil {
		flash = fmt.Sprintf("%v", session.Get("flash"))
		session.Delete("flash")
		session.Save()
	}

	data := fiber.Map{"Tcpservers": tcp_servers, "Flash": flash, "Menu": "tcp_server"}
	return c.Render("mgmt/tcp_server/index", data, "base")
}

// TCP Server 추가 폼
// /mgmt/tcp_server/insert_form
func InsertFormTCPServer(c *fiber.Ctx) error {
	return c.Render("mgmt/tcp_server/insert_form", fiber.Map{})
}

// TCP Server 추가
// /mgmt/tcp_server/insert
func InsertTCPServer(c *fiber.Ctx) error {
	name := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	port := c.FormValue("port")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_INSERT_TCPSERVER(?, ?, ?, ?)", name, ip_addr, port, timeout)

	session.Set("flash", "새로운 TCP Server("+name+")이 추가되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/tcp_server")
}

// TCP Server 수정 폼
// /mgmt/tcp_server/update_form/{id}
func UpdateFormTCPServer(c *fiber.Ctx) error {
	var tcp_server TcpServer

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_TCPSERVER(?)", id).First(&tcp_server)
	data := fiber.Map{"TCPServer": tcp_server}
	return c.Render("mgmt/tcp_server/update_form", data)
}

// TCP Server 수정
// /mgmt/tcp_server/update
func UpdateTCPServer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	port := c.FormValue("port")
	timeout := c.FormValue("timeout")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_UPDATE_TCPSERVER(?, ?, ?, ?, ?)", id, name, ip_addr, port, timeout)

	session.Set("flash", "TCP서버("+name+")가 수정되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/tcp_server")
}

// TCP Server Enabled 상태 토글하기
// /mgmt/tcp_server/toggle_enabled/{id}
func ToggleEnabledTCPServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_TOGGLE_ENABLED_TCPSERVER(?)", id)

	session.Set("flash", "TCP서버상태가 토글되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/tcp_server")
}

// TCP Server 삭제
// /mgmt/tcp_server/delete/{id}
func DeleteTCPServer(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	db := database.DBConn
	db.Exec("CALL SP_DELETE_TCPSERVER(?)", id)

	session.Set("flash", "TCP서버가 삭제되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/tcp_server")
}
