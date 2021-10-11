package mgmt 

// controllers/mgmt


import (
	"github.com/gauryan/ssmon/store"
	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"
	"fmt"
)


// SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN;
type Setting struct {
	Id    int
	Code  string
	Name  string
	Value string
	Memo  string
}


// Setting 메인화면
// /mgmt/setting
func ViewSetting(c *fiber.Ctx) error {
	var settings []Setting
	var flash string

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	db := database.DBConn
	db.Raw("CALL SP_LIST_SETTING()").Scan(&settings)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	if session.Get("flash") != nil {
		flash = fmt.Sprintf("%v", session.Get("flash"))
		session.Delete("flash")
		session.Save()
	}

	data := fiber.Map{"Settings": settings, "Flash": flash, "Menu": "setting"}
	return c.Render("mgmt/setting/index", data, "base")
}


// Setting 수정
// /mgmt/setting/update
func UpdateSetting(c *fiber.Ctx) error {
	aa := c.FormValue("ALARM_USE_YN")
	bb := c.FormValue("ERR_CNT_FOR_ALARM")
	cc := c.FormValue("SLACK_USE_YN")
	dd := c.FormValue("SLACK_CHANNEL")
	ee := c.FormValue("SLACK_TOKEN")
	ff := c.FormValue("SLACK_USERNAME")
		gg := c.FormValue("ERR_LOG_SAVE_DAYS")

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	db := database.DBConn
	db.Exec("CALL SP_UPDATE_SETTING(?, ?, ?, ?, ?, ?, ?)", aa, bb, cc, dd, ee, ff, gg)

	session.Set("flash", "설정이 저장되었습니다.")
    session.Save()

	return c.Redirect("/mgmt/setting")
}
