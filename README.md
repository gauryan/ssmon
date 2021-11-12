# SSMON (Simple Server Monitor) - 시스템 모니터링

* `아직 개발중입니다...`
* ERD : https://dbdiagram.io/d/615640cb825b5b01461b8253
* 개발언어 : [Go](https://golang.org/)
* 운영환경 : Ubuntu Linux + MySQL + [Fiber](https://gofiber.io/) + [GORM](https://gorm.io/)
* 기능
  * ✔ TCP 포트 체크
  * ✔ HTTP 체크
  * ✔ PING 체크
  * ✔ Slack 알림
  * CPU 체크
  * Memory 체크
  * HDD 체크

## 스크린샷
<img src="screenshots/tcp_monitor.png" width="450px" title="TCP Server Monitor"/>

## Install

* 먼저 서버에 Ubuntu Linux 를 설치하고, MySQL을 설치합니다. Linux는 다른 배포판을 사용할 수도 있지만, Ubuntu를 권장합니다.
* MySQL 에 DB와 계정을 생성합니다.
```sql
-- SSMON DB와 계정 생성
-- DB이름과 사용자이름은 모두 ssmon 으로 해주세요!
-- 비밀번호만 원하는 것으로 변경하세요! 아래와 같은 방법으로 해주시면 됩니다.
CREATE DATABASE ssmon DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
CREATE USER 'ssmon'@'%' IDENTIFIED BY 'ssmon123'; -- 비밀번호 변경
GRANT ALL PRIVILEGES ON ssmon.* TO 'ssmon'@'%' WITH GRANT OPTION;
CREATE USER 'ssmon'@'localhost' IDENTIFIED BY 'ssmon123'; -- 비밀번호 변경
GRANT ALL PRIVILEGES ON ssmon.* TO 'ssmon'@'localhost' WITH GRANT OPTION;
FLUSH privileges;
```
* SSMON을 다운로드 합니다. (OS계정은 Ubuntu Linux의 기본인 ubuntu로 진행합니다.)
```bash
$ cd ~
$ wget https://github.com/gauryan/ssmon/releases/download/v0.1.0/ssmon_v0.1.0.tar.gz
$ tar xvfz ssmon_v0.1.0.tar.gz
```
* 테이블 생성 및 기초데이터를 입력합니다.
```bash
$ cd ssmon/install
$ mysql -u ssmon -p ssmon < 02_create_table_sp.sql
$ mysql -u ssmon -p ssmon < 03_insert_init_data.sql
```
* `.env` 파일에서 어플리케이션의 포트와 DB설정을 합니다.
```
$ cd ~/ssmon
$ vi .env
APP_PORT=3000
DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=ssmon
DB_USERNAME=ssmon
DB_PASSWORD=ssmon123 # 원하는 비밀번호로 변경!
```
* crontab 설정을 해줍니다. 프로그램 설치위치가 다르면 조정해주세요.
```bash
$ crontab -e
* * * * * /home/ubuntu/ssmon/check/check_tcp -e /home/ubuntu/ssmon/.env
* * * * * /home/ubuntu/ssmon/check/check_http -e /home/ubuntu/ssmon/.env
* * * * * /home/ubuntu/ssmon/check/check_ping -e /home/ubuntu/ssmon/.env
1 * * * * /home/ubuntu/ssmon/check/del_logs -e /home/ubuntu/ssmon/.env
```
* 마지막으로 서버를 시작해주시면 되겠습니다.
```bash
$ cd ~/ssmon
./start
```
* 이제 웹브라우저에서 설정된 IP주소와 Port로 접속해보시면 되겠습니다.
