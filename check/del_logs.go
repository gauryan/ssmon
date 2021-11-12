package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Setting struct {
	Id    int
	Code  string
	Name  string
	Value string
	Memo  string
}

func main() {
	var settings []Setting
	var err_log_save_days int

	// .env 불러오기
	env := flag.String("e", "/home/ubuntu/ssmon/.env", ".env 파일")
	// env := flag.String("e", "/home/ubuntu/project/ssmon/.env", ".env 파일")
	flag.Parse()
	err_dot := godotenv.Load(*env)
	if err_dot != nil {
		fmt.Println("Error loading .env file")
		fmt.Println("'check_tcp -e .env파일경로'")
		return
	}

	// DB 연결
	dsn := "" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE")
	DBConn, err_gorm := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err_gorm != nil {
		fmt.Println("failed to connect database")
		return
	}

	// 로그저장일 구하기
	DBConn.Raw("CALL SP_LIST_SETTING()").First(&settings)
	for _, setting := range settings {
		if setting.Code == "ERR_LOG_SAVE_DAYS" {
			err_log_save_days, _ = strconv.Atoi(setting.Value)
		}
	}

	// 로그 삭제
	DBConn.Exec("CALL SP_DELETE_ERR_LOG(?)", err_log_save_days)
}
