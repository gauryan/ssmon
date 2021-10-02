package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/gauryan/ssmon/config"
)

var (
	DBConn *gorm.DB
)

func Init() {
	var err error
	// dsn := "ssmon:ssmon123@tcp(10.0.0.91:3306)/ssmon?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "ssmon:ssmon123@tcp(10.0.0.91:3306)/ssmon"
	// dsn := ""+config.Config("DB_USERNAME")+":"+config.Config("DB_PASSWORD")+"@tcp("+config.Config("DB_HOST")+":"+config.Config("DB_PORT")+")/"+config.Config("DB_DATABASE")
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
