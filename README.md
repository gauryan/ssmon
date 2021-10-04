# SSMON (Simple Server Monitor) - 시스템 모니터링

* `아직 개발중입니다...`
* ERD : https://dbdiagram.io/d/615640cb825b5b01461b8253
* 개발언어 : [Go](https://golang.org/)
* 운영환경 : Ubuntu Linux + MySQL + [Fiber](https://gofiber.io/) + [GORM](https://gorm.io/)
* 기능
  * PING 체크
  * TCP 포트 체크
  * HTTP 체크
  * CPU, RAM, HDD 체크 (Agent 설치)
  * Slack 알림
* 설정
  * 장애알림 사용여부 (ALARM_USE_YN)
  * 알림을 위한 장애횟수 (ERR_CNT_FOR_ALARM)
  * Slack 사용여부 (SLACK_USE_YN)
  * Slack 채널     (SLACK_CHANEL)
  * Slack 토큰     (SLACK_TOKEN)
  * Slack 사용자이름 (SLACK_USERNAME)
  * Slack 이미지 URL (SLACK_IMAGE)

### 스크린샷
<img src="screenshots/tcp_monitor.png" width="450px" title="TCP Server Monitor"/>
