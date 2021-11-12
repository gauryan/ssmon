package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-ping/ping"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PingServer struct {
	Id      int
	Name    string
	Ip_addr string
	Timeout int
	Err_cnt int
}

type Setting struct {
	Id    int
	Code  string
	Name  string
	Value string
	Memo  string
}

func main() {
	var ping_servers []PingServer
	var settings []Setting
	var err_cnt int

	var err_cnt_for_alarm int
	var alarm_use_yn string
	var slack_use_yn string

	// .env 불러오기
	env := flag.String("e", "/home/ubuntu/ssmon/.env", ".env 파일")
	// env := flag.String("e", "/home/ubuntu/project/ssmon/.env", ".env 파일")
	flag.Parse()
	err_dot := godotenv.Load(*env)
	if err_dot != nil {
		fmt.Println("Error loading .env file")
		fmt.Println("'check_ping -e .env파일경로'")
		return
	}

	// DB 연결
	dsn := "" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE")
	DBConn, err_gorm := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err_gorm != nil {
		fmt.Println("failed to connect database")
		return
	}
	// 알림설정 구하기
	DBConn.Raw("CALL SP_LIST_SETTING()").First(&settings)
	for _, setting := range settings {
		if setting.Code == "ERR_CNT_FOR_ALARM" {
			err_cnt_for_alarm, _ = strconv.Atoi(setting.Value)
		}
		if setting.Code == "ALARM_USE_YN" {
			alarm_use_yn = setting.Value
		}
		if setting.Code == "SLACK_USE_YN" {
			slack_use_yn = setting.Value
		}
	}

	DBConn.Raw("CALL SP_MONITOR_PINGSERVER()").Scan(&ping_servers)
	fmt.Println("Check PING Servers...")
	for _, ts := range ping_servers {
		fail := false

		pinger, err := ping.NewPinger(ts.Ip_addr)
		_ = err
		/*
			if err != nil {
				panic(err)
			}
		*/
		pinger.Count = 1
		pinger.Timeout = time.Duration(ts.Timeout) * time.Millisecond
		err = pinger.Run() // Blocks until finished.
		/*
			if err != nil {
				panic(err)
			}
		*/
		stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
		/*
			fmt.Println(stats)
			fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		*/

		if stats.PacketLoss == 100 {
			fail = true
		}

		// fmt.Println("Faile:", fail)

		if fail == true {
			// PING 실패
			fmt.Println(ts.Name + ": PING Fail")
			err_cnt = ts.Err_cnt + 1
			DBConn.Exec("CALL SP_UPDATE_PING_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
			if ts.Err_cnt+1 == err_cnt_for_alarm {
				// 장애 로그/알림 남긴다
				// service, err_rec_gubun, name, ip_addr, port, url
				DBConn.Exec("CALL SP_INSERT_ERR_LOG(?, ?, ?, ?, null, null)", "PING", "장애", ts.Name, ts.Ip_addr)
				// 알림 설정이 되어 있으면, Slack 메시지 보낸다.
				if alarm_use_yn == "Y" && slack_use_yn == "Y" {
					msg := ":rotating_light: [장애] [PING] " + ts.Name + " 》》》 " + ts.Ip_addr
					send_slack(&settings, msg)
				}
				fmt.Println("장애 로그/알림")
			}
		} else {
			// PING 성공
			fmt.Println(ts.Name + ": PING Success")
			if ts.Err_cnt == 0 {
				// 아무것도 하지 않는다.
			} else if ts.Err_cnt >= err_cnt_for_alarm {
				err_cnt = 0
				DBConn.Exec("CALL SP_UPDATE_PING_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
				// 복구 로그/알림 남긴다
				DBConn.Exec("CALL SP_INSERT_ERR_LOG(?, ?, ?, ?, null, null)", "PING", "복구", ts.Name, ts.Ip_addr)
				// 알림 설정이 되어 있으면, Slack 메시지 보낸다.
				if alarm_use_yn == "Y" && slack_use_yn == "Y" {
					msg := ":smile: [복구] [PING] " + ts.Name + " 》》》 " + ts.Ip_addr
					send_slack(&settings, msg)
				}
				fmt.Println("복구 로그/알림")
			} else if ts.Err_cnt != 0 && ts.Err_cnt < err_cnt_for_alarm {
				err_cnt = 0
				DBConn.Exec("CALL SP_UPDATE_PING_SERVER_ERR_CNT(?, ?)", ts.Id, err_cnt)
			}
		}
	}
}

// Send Slack Message
func send_slack(settings *[]Setting, text string) {
	var slack_channel string
	var slack_token string
	var slack_username string

	for _, setting := range *settings {
		if setting.Code == "SLACK_CHANNEL" {
			slack_channel = setting.Value
		}
		if setting.Code == "SLACK_TOKEN" {
			slack_token = setting.Value
		}
		if setting.Code == "SLACK_USERNAME" {
			slack_username = setting.Value
		}
	}

	loc, err_time := time.LoadLocation("Asia/Seoul")
	if err_time != nil {
		fmt.Println(err_time)
	}
	now := time.Now()
	now_local := now.In(loc)
	log_time := now_local.Format("2006-01-02 15:04:05")
	slack_text := log_time + " " + text
	// 실제로 SLACK 메시지 보낸다.
	attachment := slack.Attachment{
		Text: slack_text,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
		Channel:     slack_channel,
		Username:    slack_username,
	}

	err := slack.PostWebhook(slack_token, &msg)
	if err != nil {
		fmt.Println(err)
	}
}
