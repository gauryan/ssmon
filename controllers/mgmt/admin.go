package mgmt 

// controllers/mgmt


import (
	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"
	// "fmt"
)


// SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN;
type Admin struct {
	Id     int
	Userid string
	Passwd string
	Nick   string
	Phone  string
}


// Admin 목록
// /mgmt/admin
func ListAdmin(c *fiber.Ctx) error {
	var admins []Admin

	db := database.DBConn
	db.Raw("CALL SP_LIST_ADMIN()").Scan(&admins)
	// db.Raw("SELECT id, userid, passwd, nick, phone FROM TB_ADMIN").Scan(&admins)
	// db.Raw("SELECT ID, USERID, PASSWD, NICK, PHONE FROM TB_ADMIN").Scan(&admins)
	// 컬럼은 소문자로 써야 하며, 테이블이름은 대소문자를 가린다.

	data := fiber.Map{"Admins": admins}
	return c.Render("mgmt/admin/index", data, "base")
}


// 관리자 추가 폼
// /mgmt/admin/insert_form
func InsertFormAdmin(c *fiber.Ctx) error {
	return c.Render("mgmt/admin/insert_form", fiber.Map{})
}


// 관리자 추가
// /mgmt/admin/insert
func InsertAdmin(c *fiber.Ctx) error {
	userid  := c.FormValue("userid")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")
	nick    := c.FormValue("nick")
	phone   := c.FormValue("phone")

	if passwd1 != passwd2 {
		return c.Redirect("/mgmt/admin")
	}
	db := database.DBConn
	// db.Exec("CALL insertAdmin(?, ?, ?)", userid, passwd1, nick)
	db.Exec("CALL SP_INSERT_ADMIN(?, ?, ?, ?)", userid, passwd1, nick, phone)

	return c.Redirect("/mgmt/admin")
}


// 관리자 비밀번호변경 폼
// /mgmt/admin/chg_passwd_form/:id
func ChgPasswdFormAdmin(c *fiber.Ctx) error {
	var admin Admin

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL SP_GET_ADMIN(?)", id).First(&admin)
	data := fiber.Map{"Admin": admin}
	return c.Render("mgmt/admin/chg_passwd_form", data)
}


// 관리자 비밀번호변경
// /mgmt/admin/chg_passwd
func ChgPasswdAdmin(c *fiber.Ctx) error {
	id := c.FormValue("id")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")

	if passwd1 != passwd2 {
		return c.Redirect("/mgmt/admin")
	}
	db := database.DBConn
	db.Exec("CALL SP_UPDATE_ADMIN_PASSWD(?, ?)", id, passwd1)

	return c.Redirect("/mgmt/admin")
}


// 관리자 수정 폼
// /admin/update_form/{id}
/*
func UpdateForm(c *fiber.Ctx) error {
	type Admin struct {
		Sno    int
		Userid string
		Passwd string
		Nick   string
	}
	var admin Admin

	id := c.Params("id")

	db := database.DBConn
	db.Raw("CALL getAdmin(?)", id).First(&admin)
	data := fiber.Map{"Admin": admin}
	return c.Render("admin/update_form", data)
}
*/

// 관리자 수정
// /admin/update
/*
func Update(c *fiber.Ctx) error {
	id := c.FormValue("id")
	nick := c.FormValue("nick")

	db := database.DBConn
	db.Exec("CALL updateAdmin(?, ?)", id, nick)

	return c.Redirect("/admin")
}
*/

// 관리자 삭제
// /admin/delete/{id}
/*
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn
	db.Exec("CALL deleteAdmin(?)", id)
	return c.Redirect("/admin")
}
*/

