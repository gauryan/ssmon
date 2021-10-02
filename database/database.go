package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() {
	var err error
	// dsn := "xyz:xyz123@tcp(10.0.0.91:3306)/xyz?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "xyz:xyz123@tcp(10.0.0.91:3306)/xyz"
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
