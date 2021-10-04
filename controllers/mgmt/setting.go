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
	Type  string
}


// Setting 메인화면
// /mgmt/settingn
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


// 관리자 추가 폼
// /mgmt/admin/insert_form
/*
func InsertFormAdmin(c *fiber.Ctx) error {
	return c.Render("mgmt/admin/insert_form", fiber.Map{})
}
*/


// 관리자 추가
// /mgmt/admin/insert
/*
func InsertAdmin(c *fiber.Ctx) error {
	userid  := c.FormValue("userid")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")
	nick    := c.FormValue("nick")
	phone   := c.FormValue("phone")

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	if passwd1 != passwd2 {
		return c.Redirect("/mgmt/admin")
	}
	db := database.DBConn
	db.Exec("CALL SP_INSERT_ADMIN(?, ?, ?, ?)", userid, passwd1, nick, phone)

	session.Set("flash", "새로운 관리자("+nick+")이 추가되었습니다.")
	session.Save()

	return c.Redirect("/mgmt/admin")
}
*/


// 관리자 비밀번호변경 폼
// /mgmt/admin/chg_passwd_form/:id
/*
func ChgPasswdFormAdmin(c *fiber.Ctx) error {
	var admin Admin

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_ADMIN(?)", id).First(&admin)
	data := fiber.Map{"Admin": admin}
	return c.Render("mgmt/admin/chg_passwd_form", data)
}
*/


// 관리자 비밀번호변경
// /mgmt/admin/chg_passwd
/*
func ChgPasswdAdmin(c *fiber.Ctx) error {
	id := c.FormValue("id")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	if passwd1 != passwd2 {
		return c.Redirect("/mgmt/admin")
	}
	db := database.DBConn
	db.Exec("CALL SP_UPDATE_ADMIN_PASSWD(?, ?)", id, passwd1)

	session.Set("flash", "관리자 비밀번호가 변경되었습니다.")
    session.Save()

	return c.Redirect("/mgmt/admin")
}
*/


// 관리자 수정 폼
// /mgmt/admin/update_form/{id}
/*
func UpdateFormAdmin(c *fiber.Ctx) error {
	var admin Admin

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_ADMIN(?)", id).First(&admin)
	data := fiber.Map{"Admin": admin}
	return c.Render("mgmt/admin/update_form", data)
}
*/


// 관리자 수정
// /mgmt/admin/update
/*
func UpdateAdmin(c *fiber.Ctx) error {
	id    := c.FormValue("id")
	nick  := c.FormValue("nick")
	phone := c.FormValue("phone")

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	db := database.DBConn
	db.Exec("CALL SP_UPDATE_ADMIN(?, ?, ?)", id, nick, phone)

	session.Set("flash", "관리자("+nick+")가 수정되었습니다.")
    session.Save()

	return c.Redirect("/mgmt/admin")
}
*/


// 관리자 삭제
// /mgmt/admin/delete/{id}
/*
func DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := store.SessionStore.Get(c)
    if err != nil {
        panic(err)
    }

	db := database.DBConn
	db.Exec("CALL SP_DELETE_ADMIN(?)", id)

	session.Set("flash", "관리자가 삭제되었습니다.")
    session.Save()

	return c.Redirect("/mgmt/admin")
}
*/


