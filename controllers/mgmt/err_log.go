package mgmt

// controllers/mgmt

import (
	// "github.com/gauryan/ssmon/store"
	"html/template"

	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"

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
	var start int = (current_page / PSIZE) * PSIZE
	pagination := "<nav><ul class='pagination justify-content-center'>"

	if start != 0 {
		pagination = pagination + "<li class='page-item'><a class='page-link' href='/mgmt/errlog/" + strconv.Itoa(start-PSIZE) + "'> Prev </a></li> "
	}

	for i := start; i < start+PSIZE; i++ {
		if i == current_page {
			pagination = pagination + " <li class='page-item active' aria-current='page'><a class='page-link' href='#'>" + strconv.Itoa(i+1) + "</a></li> "
		} else {
			pagination = pagination + " <li class='page-item'><a class='page-link' href='/mgmt/errlog/" + strconv.Itoa(i) + "'>" + strconv.Itoa(i+1) + "</a></li> "
		}

		if i == last_page {
			is_last_page = true
			break
		}
	}
	if is_last_page == false {
		pagination = pagination + "<li class='page-item'><a href='/mgmt/errlog/" + strconv.Itoa(start+PSIZE) + "'> Next </a></li> "
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
	if last_page < 0 {
		last_page = 0
	}
	pagination := pager(page, last_page)
	// fmt.Println("last_pate:", last_page)

	data := fiber.Map{"Errlogs": errlogs, "Pagination": template.HTML(pagination), "Menu": "errlog"}
	return c.Render("mgmt/errlog/index", data, "base")
}
