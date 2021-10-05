package main

import (
	"fmt"
	"net"
	"time"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TcpServer struct {
	Id      int
	Name    string
	Ip_addr string
	Port    int
	Timeout int
	Err_cnt int
}

type Result struct {
	Value string
}

func main() {
	var tcp_servers []TcpServer
	var result Result
	var err_cnt int

	// .env 불러오기
	err_dot := godotenv.Load("../.env")
	if err_dot != nil {
        panic("Error loading .env file")
    }

	// DB 연결
	dsn := ""+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")
	DBConn, err_gorm := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err_gorm != nil {
        panic("failed to connect database")
    }
	DBConn.Raw("CALL SP_GET_ERR_CNT_FOR_ALARM()").First(&result)
	err_cnt_for_alarm, _ := strconv.Atoi(result.Value)


	DBConn.Raw("CALL SP_MONITOR_TCPSERVER()").Scan(&tcp_servers)
	fmt.Println("Check TCP Servers...")
	for _, ts := range tcp_servers {
		conn, err_tcp := net.DialTimeout("tcp", ts.Ip_addr+":"+strconv.Itoa(ts.Port), time.Duration(ts.Timeout) * time.Millisecond)
		if nil != err_tcp {
			// TCP 연결 실패
			fmt.Println(ts.Name+": TCP Connection Fail")
			err_cnt = ts.Err_cnt + 1
			DBConn.Exec("CALL SP_UPDATE_TCP_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
			if ts.Err_cnt+1 == err_cnt_for_alarm {
				// 장애 로그/알림 남긴다
				// service, err_rec_gubun, name, ip_addr, port, url
				DBConn.Exec("CALL SP_INSERT_ERR_LOG(?, ?, ?, ?, ?, ?)", "TCP", "장애", ts.Name, ts.Ip_addr, ts.Port, "")
				fmt.Println("장애 로그/알림")
			}
		} else {
			// TCP 연결 성공
			conn.Close()
			fmt.Println(ts.Name+": TCP Connection Success")
			if ts.Err_cnt == 0 {
				// 아무것도 하지 않는다.
			} else if ts.Err_cnt >= err_cnt_for_alarm {
				err_cnt = 0
				DBConn.Exec("CALL SP_UPDATE_TCP_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
				// 복구 로그/알림 남긴다
				DBConn.Exec("CALL SP_INSERT_ERR_LOG(?, ?, ?, ?, ?, ?)", "TCP", "복구", ts.Name, ts.Ip_addr, ts.Port, "")
				fmt.Println("복구 로그/알림")
			} else if ts.Err_cnt != 0 && ts.Err_cnt < err_cnt_for_alarm {
				err_cnt = 0
				DBConn.Exec("CALL SP_UPDATE_TCP_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
			}
		}
	}
}
