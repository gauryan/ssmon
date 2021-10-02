package controllers


import (
	"github.com/gauryan/ssmon/database"
	"github.com/gofiber/fiber/v2"
)

// Admin 목록
// /admin
func ListAdmin(c *fiber.Ctx) error {
	type Admin struct {
		Sno    int
		Userid string
		Nick   string
	}
	var admins []Admin

	db := database.DBConn
	// db.Raw("SELECT sno, userid, nick FROM admins").Scan(&admins)
	db.Raw("CALL listAdmins()").Scan(&admins)

	data := fiber.Map{"Admins": admins}
	return c.Render("admin/index", data, "base")
}

// 관리자 추가 폼
// /admin/inser_form
func InsertForm(c *fiber.Ctx) error {
	return c.Render("admin/insert_form", fiber.Map{})
}

// 관리자 추가
// /admin/insert
func Insert(c *fiber.Ctx) error {
	userid := c.FormValue("userid")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")
	nick := c.FormValue("nick")

	if passwd1 != passwd2 {
		return c.Redirect("/admin")
	}
	db := database.DBConn
	db.Exec("CALL insertAdmin(?, ?, ?)", userid, passwd1, nick)

	return c.Redirect("/admin")
}

// 관리자 비밀번호변경 폼
// /admin/chg_passwd_form/:id
func ChgPasswdForm(c *fiber.Ctx) error {
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
	return c.Render("admin/chg_passwd_form", data)
}

// 관리자 비밀번호변경
// /admin/chg_passwd
func ChgPasswd(c *fiber.Ctx) error {
	id := c.FormValue("id")
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")

	if passwd1 != passwd2 {
		return c.Redirect("/admin")
	}
	db := database.DBConn
	db.Exec("CALL updateAdminPassword(?, ?)", id, passwd1)

	return c.Redirect("/admin")
}

// 관리자 수정 폼
// /admin/update_form/{id}
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

// 관리자 수정
// /admin/update
func Update(c *fiber.Ctx) error {
	id := c.FormValue("id")
	nick := c.FormValue("nick")

	db := database.DBConn
	db.Exec("CALL updateAdmin(?, ?)", id, nick)

	return c.Redirect("/admin")
}

// 관리자 삭제
// /admin/delete/{id}
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn
	db.Exec("CALL deleteAdmin(?)", id)
	return c.Redirect("/admin")
}
