DROP TABLE IF EXISTS `TB_ADMIN`;
CREATE TABLE `TB_ADMIN` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `userid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `passwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'SHA256 으로 저장',
  `nick` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `USER_ID` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `TB_ERR_LOG`;
CREATE TABLE `TB_ERR_LOG` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `err_time` datetime DEFAULT (now()),
  `err_rec_gubun` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '장애/복구 구분',
  `service` varchar(20) NOT NULL COMMENT 'TCP, PING, HTTP',
  `name` varchar(255) NOT NULL,
  `ip_addr` varchar(20) DEFAULT NULL,
  `port` varchar(10) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `TB_HTTP_SERVER`;
CREATE TABLE `TB_HTTP_SERVER` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `enabled` int NOT NULL DEFAULT '1' COMMENT '활성/비활성',
  `name` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `chk_str` varchar(1000) NOT NULL DEFAULT '',
  `timeout` int NOT NULL DEFAULT '300' COMMENT '단위: ms',
  `err_cnt` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `url` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `TB_PING_SERVER`;
CREATE TABLE `TB_PING_SERVER` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `enabled` int NOT NULL DEFAULT '1' COMMENT '활성/비활성',
  `name` varchar(255) NOT NULL DEFAULT '',
  `ip_addr` varchar(20) NOT NULL DEFAULT '',
  `timeout` int NOT NULL DEFAULT '300' COMMENT '단위: ms',
  `err_cnt` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `ip_addr` (`ip_addr`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `TB_SETTING`;
CREATE TABLE `TB_SETTING` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `value` varchar(255) NOT NULL DEFAULT '',
  `memo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;

DROP TABLE IF EXISTS `TB_TCP_SERVER`;
CREATE TABLE `TB_TCP_SERVER` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `enabled` int NOT NULL DEFAULT '1',
  `name` varchar(255) NOT NULL DEFAULT '',
  `ip_addr` varchar(20) NOT NULL DEFAULT '',
  `port` int NOT NULL DEFAULT '80',
  `timeout` int NOT NULL DEFAULT '300' COMMENT '단위: ms',
  `err_cnt` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_DELETE_ADMIN`(i_id INT)
BEGIN
	DELETE FROM TB_ADMIN WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_DELETE_HTTPSERVER`(i_id INT)
BEGIN
	DELETE FROM TB_HTTP_SERVER WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_DELETE_PINGSERVER`(i_id INT)
BEGIN
	DELETE FROM TB_PING_SERVER WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_DELETE_TCPSERVER`(i_id INT)
BEGIN
	DELETE FROM TB_TCP_SERVER WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_ADMIN`(i_id INT)
BEGIN
	SELECT id, userid, passwd, nick, phone FROM TB_ADMIN WHERE id = i_id LIMIT 1;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_ERR_CNT_FOR_ALARM`()
BEGIN
	SELECT value FROM TB_SETTING WHERE code ='ERR_CNT_FOR_ALARM' LIMIT 1;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_HTTPSERVER`(i_id INT)
BEGIN
	SELECT id, name, url, chk_str, timeout FROM TB_HTTP_SERVER WHERE id = i_id LIMIT 1;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_PINGSERVER`(i_id INT)
BEGIN
	SELECT id, name, ip_addr, timeout FROM TB_PING_SERVER WHERE id = i_id LIMIT 1;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_TCPSERVER`(i_id INT)
BEGIN
	SELECT id, name, ip_addr, port, timeout FROM TB_TCP_SERVER WHERE id = i_id LIMIT 1;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_GET_TOTAL_CNT_ERR_LOG`()
BEGIN
	SELECT COUNT(*) as total_cnt FROM TB_ERR_LOG;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_INSERT_ADMIN`(
   i_userid VARCHAR(255),
   i_passwd VARCHAR(255),
   i_nick VARCHAR(255),
   i_phone VARCHAR(20))
BEGIN
  INSERT INTO TB_ADMIN(userid, passwd, nick, phone) VALUES(i_userid, SHA2(i_passwd, 256), i_nick, i_phone);
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_INSERT_ERR_LOG`(
	i_service VARCHAR(255),
	i_err_rec_gubun VARCHAR(10),
	i_name VARCHAR(255),
	i_ip_addr VARCHAR(20),
	i_port VARCHAR(10),
	i_url VARCHAR(255))
BEGIN
	INSERT INTO TB_ERR_LOG(service, err_rec_gubun, name, ip_addr, port, url) 
	VALUES(i_service, i_err_rec_gubun, i_name, i_ip_addr, i_port, i_url);
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_INSERT_HTTPSERVER`(
	i_name    VARCHAR(255),
	i_url     VARCHAR(1000),
	i_chk_str VARCHAR(255),
	i_timeout INT)
BEGIN
	INSERT INTO TB_HTTP_SERVER(name, url, chk_str, timeout) VALUES(i_name, i_url, i_chk_str, i_timeout);
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_INSERT_PINGSERVER`(
	i_name    VARCHAR(255),
	i_ip_addr VARCHAR(255),
	i_timeout INT)
BEGIN
	INSERT INTO TB_PING_SERVER(name, ip_addr, timeout) VALUES(i_name, i_ip_addr, i_timeout);
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_INSERT_TCPSERVER`(
	i_name    VARCHAR(255),
	i_ip_addr VARCHAR(255),
	i_port    INT,
	i_timeout INT)
BEGIN
	INSERT INTO TB_TCP_SERVER(name, ip_addr, port, timeout) VALUES(i_name, i_ip_addr, i_port, i_timeout);
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_IS_ADMIN`(
  i_userid VARCHAR(255),
  i_passwd VARCHAR(255)
)
BEGIN
  DECLARE CNT INT;

  SELECT COUNT(*) INTO CNT FROM TB_ADMIN 
  WHERE userid = i_userid AND passwd = SHA2(i_passwd, 256);

  SELECT IF(CNT > 0, 'Y', 'N') AS is_admin;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_ADMIN`()
BEGIN
	SELECT id, userid, passwd, nick, phone FROM TB_ADMIN;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_ERR_LOG`(
	i_page INT,
	i_cnt  INT)
BEGIN
	DECLARE start_index INT;
	SET start_index = i_page * 10;
	
	SELECT err_time, err_rec_gubun, service, name, ip_addr, port, url 
	FROM TB_ERR_LOG ORDER BY err_time DESC 
	LIMIT start_index, i_cnt;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_HTTPSERVER`()
BEGIN
	SELECT id, enabled, name, url, chk_str, timeout, err_cnt FROM TB_HTTP_SERVER ORDER BY name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_PINGSERVER`()
BEGIN
	SELECT id, enabled, name, ip_addr, timeout, err_cnt FROM TB_PING_SERVER ORDER BY name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_SETTING`()
BEGIN
	SELECT id, code, name, value, memo FROM TB_SETTING;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_LIST_TCPSERVER`()
BEGIN
	SELECT id, enabled, name, ip_addr, port, timeout, err_cnt FROM TB_TCP_SERVER ORDER BY name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_MONITOR_HTTPSERVER`()
BEGIN
	SELECT id, enabled, name, url, chk_str, timeout, err_cnt 
	FROM   TB_HTTP_SERVER 
	WHERE  enabled = 1 
	ORDER BY err_cnt DESC, name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_MONITOR_PINGSERVER`()
BEGIN
	SELECT id, enabled, name, ip_addr, timeout, err_cnt 
	FROM   TB_PING_SERVER 
	WHERE  enabled = 1 
	ORDER BY err_cnt DESC, name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_MONITOR_TCPSERVER`()
BEGIN
	SELECT id, enabled, name, ip_addr, port, timeout, err_cnt 
	FROM   TB_TCP_SERVER 
	WHERE  enabled = 1 
	ORDER BY err_cnt DESC, name;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_TOGGLE_ENABLED_HTTPSERVER`( i_id INT )
BEGIN
	DECLARE i_enabled INT;

	SELECT enabled INTO i_enabled FROM TB_HTTP_SERVER WHERE id = i_id;
	UPDATE TB_HTTP_SERVER SET enabled = IF(i_enabled = 0, 1, 0) WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_TOGGLE_ENABLED_PINGSERVER`( i_id INT )
BEGIN
	DECLARE i_enabled INT;

	SELECT enabled INTO i_enabled FROM TB_PING_SERVER WHERE id = i_id;
	UPDATE TB_PING_SERVER SET enabled = IF(i_enabled = 0, 1, 0) WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_TOGGLE_ENABLED_TCPSERVER`( i_id INT )
BEGIN
	DECLARE i_enabled INT;

	SELECT enabled INTO i_enabled FROM TB_TCP_SERVER WHERE id = i_id;
	UPDATE TB_TCP_SERVER SET enabled = IF(i_enabled = 0, 1, 0) WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_ADMIN`(
	i_id INT,
	i_nick VARCHAR(255),
	i_phone VARCHAR(20))
BEGIN
	UPDATE TB_ADMIN SET nick = i_nick, phone = i_phone WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_ADMIN_PASSWD`(
	i_id INT,
	i_passwd VARCHAR(255))
BEGIN
	UPDATE TB_ADMIN SET passwd = SHA2(i_passwd, 256) WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_HTTPSERVER`(
    i_id INT,
    i_name VARCHAR(255),
    i_url VARCHAR(1000),
	i_chk_str VARCHAR(255),
	i_timeout INT )
BEGIN
	UPDATE TB_HTTP_SERVER SET name = i_name, url = i_url, chk_str = i_chk_str, timeout = i_timeout WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_HTTP_SERVER_ERR_CNT`(
	i_id BIGINT,
	i_err_cnt INT)
BEGIN
	UPDATE TB_HTTP_SERVER SET err_cnt = i_err_cnt WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_PINGSERVER`(
    i_id INT,
    i_name VARCHAR(255),
    i_ip_addr VARCHAR(255),
	i_timeout INT )
BEGIN
	UPDATE TB_PING_SERVER SET name = i_name, ip_addr = i_ip_addr, timeout = i_timeout WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_PING_SERVER_ERR_CNT`(
	i_id BIGINT,
	i_err_cnt INT)
BEGIN
	UPDATE TB_PING_SERVER SET err_cnt = i_err_cnt WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_SETTING`(
	ALARM_USE_YN CHAR(1),
	ERR_CNT_FOR_ALARM VARCHAR(10),
	SLACK_USE_YN CHAR(1),
	SLACK_CHANNEL VARCHAR(255),
	SLACK_TOKEN VARCHAR(255),
	SLACK_USERNAME VARCHAR(255),
	ERR_LOG_SAVE_DAYS VARCHAR(10) )
BEGIN
	UPDATE TB_SETTING SET VALUE = ALARM_USE_YN WHERE CODE = 'ALARM_USE_YN';
	UPDATE TB_SETTING SET VALUE = ERR_CNT_FOR_ALARM WHERE CODE = 'ERR_CNT_FOR_ALARM';
	UPDATE TB_SETTING SET VALUE = SLACK_USE_YN WHERE CODE = 'SLACK_USE_YN';
	UPDATE TB_SETTING SET VALUE = SLACK_CHANNEL WHERE CODE = 'SLACK_CHANNEL';
	UPDATE TB_SETTING SET VALUE = SLACK_TOKEN WHERE CODE = 'SLACK_TOKEN';
	UPDATE TB_SETTING SET VALUE = SLACK_USERNAME WHERE CODE = 'SLACK_USERNAME';
	UPDATE TB_SETTING SET VALUE = ERR_LOG_SAVE_DAYS WHERE CODE = 'ERR_LOG_SAVE_DAYS';
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_TCPSERVER`(
    i_id INT,
    i_name VARCHAR(255),
    i_ip_addr VARCHAR(255),
	i_port INT,
	i_timeout INT )
BEGIN
	UPDATE TB_TCP_SERVER SET name = i_name, ip_addr = i_ip_addr, port = i_port, timeout = i_timeout WHERE id = i_id;
END ;;
DELIMITER ;

DELIMITER ;;
CREATE DEFINER=`ssmon`@`%` PROCEDURE `SP_UPDATE_TCP_SERVER_ERR_CNT`(
	i_id BIGINT,
	i_err_cnt INT)
BEGIN
	UPDATE TB_TCP_SERVER SET err_cnt = i_err_cnt WHERE id = i_id;
END ;;
DELIMITER ;

