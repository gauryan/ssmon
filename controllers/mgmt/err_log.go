package mgmt 

// controllers/mgmt


import (
	// "github.com/gauryan/ssmon/store"
	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"
	"html/template"
	// "fmt"
	"strconv"
)


type Errlog struct {
	Id            int
	Err_time      string
	Err_rec_gubun string
	Service       string
	Name          string
	Ip_addr       string
	Port          string
	Url           string
}


const PSIZE = 20

func pager(current_page int, last_page int) string {
	is_last_page := false
	var start int = (current_page / PSIZE ) * PSIZE
	pagination := "<nav><ul class='pagination justify-content-center'>"

	if start != 0 {
		pagination = pagination + "<li class='page-item'><a class='page-link' href='/mgmt/errlog/"+strconv.Itoa(start-PSIZE)+"'> Prev </a></li> "
	}

	for i := start; i < start+PSIZE; i ++ {
		if i == current_page {
			pagination = pagination + " <li class='page-item active' aria-current='page'><a class='page-link' href='#'>"+strconv.Itoa(i+1)+"</a></li> "
		} else {
			pagination = pagination + " <li class='page-item'><a class='page-link' href='/mgmt/errlog/"+strconv.Itoa(i)+"'>"+strconv.Itoa(i+1)+"</a></li> "
		}

		if i == last_page {
			is_last_page = true
			break
		}
	}
	if is_last_page == false {
		pagination = pagination + "<li class='page-item'><a href='/mgmt/errlog/"+strconv.Itoa(start+PSIZE)+"'> Next </a></li> "
	}

	pagination = pagination + "</ul></nav>"

	return pagination
}

// 장애/복구 로그
// /mgmt/errlog/:page
func ListErrRecLog(c *fiber.Ctx) error {
	type Result struct {
		Total_cnt int
	}
	var errlogs []Errlog
	var result Result

	page, _ := strconv.Atoi(c.Params("page"))

	db := database.DBConn
	db.Raw("CALL SP_LIST_ERR_LOG(?, ?)", page, PSIZE).Scan(&errlogs)
	db.Raw("CALL SP_GET_TOTAL_CNT_ERR_LOG()").First(&result)
	total_cnt := result.Total_cnt
	last_page := (total_cnt / 10) - 1
	pagination := pager(page, last_page)

	data := fiber.Map{"Errlogs": errlogs, "Pagination": template.HTML(pagination), "Menu": "errlog"}
	return c.Render("mgmt/errlog/index", data, "base")
}


// TCP Server 추가 폼
// /mgmt/tcp_server/insert_form
/*
func InsertFormTCPServer(c *fiber.Ctx) error {
	return c.Render("mgmt/tcp_server/insert_form", fiber.Map{})
}
*/


// TCP Server 추가
// /mgmt/tcp_server/insert
/*
func InsertTCPServer(c *fiber.Ctx) error {
	name    := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	port    := c.FormValue("port")
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
*/

// TCP Server 수정 폼
// /mgmt/tcp_server/update_form/{id}
/*
func UpdateFormTCPServer(c *fiber.Ctx) error {
	var tcp_server TcpServer

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_TCPSERVER(?)", id).First(&tcp_server)
	data := fiber.Map{"TCPServer": tcp_server}
	return c.Render("mgmt/tcp_server/update_form", data)
}
*/


// TCP Server Enabled 상태 토글하기
// /mgmt/tcp_server/toggle_enabled/{id}
/*
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
*/


// TCP Server 수정
// /mgmt/tcp_server/update
/*
func UpdateTCPServer(c *fiber.Ctx) error {
	id      := c.FormValue("id")
	name    := c.FormValue("name")
	ip_addr := c.FormValue("ip_addr")
	port    := c.FormValue("port")
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
*/


// TCP Server 삭제
// /mgmt/tcp_server/delete/{id}
/*
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
*/

