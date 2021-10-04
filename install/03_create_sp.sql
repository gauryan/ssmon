-- 관리자 목록
DELIMITER $$
CREATE PROCEDURE SP_LIST_ADMIN()
BEGIN
	SELECT id, userid, passwd, nick, phone FROM TB_ADMIN;
END $$
DELIMITER ;


-- 관리자인가?
DELIMITER $$
CREATE PROCEDURE SP_IS_ADMIN (
	i_userid VARCHAR(255),
	i_passwd VARCHAR(255)
)
BEGIN
	DECLARE CNT INT;

	SELECT COUNT(*) INTO CNT FROM TB_ADMIN
	WHERE userid = i_userid AND passwd = SHA2(i_passwd, 256);

	SELECT IF(CNT > 0, 'Y', 'N') AS is_admin;
END $$
DELIMITER ;


-- 관리자 추가
DELIMITER $$
CREATE PROCEDURE SP_INSERT_ADMIN (
	i_userid VARCHAR(255),
	i_passwd VARCHAR(255),
	i_nick VARCHAR(255),
	i_phone VARCHAR(20))
BEGIN
	INSERT INTO TB_ADMIN(userid, passwd, nick, phone) VALUES(i_userid, SHA2(i_passwd, 256), i_nick, i_phone);
END $$
DELIMITER ;


-- 관리자 1명 가져오기
DELIMITER $$
CREATE PROCEDURE SP_GET_ADMIN (i_id INT)
BEGIN
	SELECT id, userid, passwd, nick, phone FROM TB_ADMIN WHERE id = i_id LIMIT 1;
END $$
DELIMITER ;


-- 관리자 비밀번호 변경
DELIMITER $$
CREATE PROCEDURE SP_UPDATE_ADMIN_PASSWD (
	i_id INT,
	i_passwd VARCHAR(255))
BEGIN
	UPDATE TB_ADMIN SET passwd = SHA2(i_passwd, 256) WHERE id = i_id;
END $$
DELIMITER ;


-- 관리자 수정하기
DELIMITER $$
CREATE PROCEDURE SP_UPDATE_ADMIN (
	i_id INT,
	i_nick VARCHAR(255),
	i_phone VARCHAR(20))
BEGIN
	UPDATE TB_ADMIN SET nick = i_nick, phone = i_phone WHERE id = i_id;
END $$
DELIMITER ;


-- 관리자 삭제하기
DELIMITER $$
CREATE PROCEDURE SP_DELETE_ADMIN (i_id INT)
BEGIN
	DELETE FROM TB_ADMIN WHERE id = i_id;
END $$
DELIMITER ;


-- TCP서버 모니터링
DELIMITER $$
CREATE PROCEDURE SP_MONITOR_TCPSERVER()
BEGIN
    SELECT id, name, ip_addr, port, timeout, err_cnt FROM TB_TCP_SERVER order by err_cnt desc;
END $$
DELIMITER ;


-- TCP서버 목록
DELIMITER $$
CREATE PROCEDURE SP_LIST_TCPSERVER()
BEGIN
    SELECT id, name, ip_addr, port, timeout, err_cnt FROM TB_TCP_SERVER;
END $$
DELIMITER ;


-- TCP서버 추가
DELIMITER $$
CREATE PROCEDURE SP_INSERT_TCPSERVER (
	i_name    VARCHAR(255),
	i_ip_addr VARCHAR(255),
	i_port    INT,
	i_timeout INT)
BEGIN
	INSERT INTO TB_TCP_SERVER(name, ip_addr, port, timeout) VALUES(i_name, i_ip_addr, i_port, i_timeout);
END $$
DELIMITER ;


-- TCP서버 1개 가져오기
DELIMITER $$
CREATE PROCEDURE SP_GET_TCPSERVER (i_id INT)
BEGIN
	SELECT id, name, ip_addr, port, timeout FROM TB_TCP_SERVER WHERE id = i_id LIMIT 1;
END $$
DELIMITER ;


-- TCP서버 수정하기
DELIMITER $$
CREATE PROCEDURE SP_UPDATE_TCPSERVER (
    i_id INT,
    i_name VARCHAR(255),
    i_ip_addr VARCHAR(255),
	i_port INT,
	i_timeout INT )
BEGIN
    UPDATE TB_TCP_SERVER SET name = i_name, ip_addr = i_ip_addr, port = i_port, timeout = i_timeout WHERE id = i_id;
END $$
DELIMITER ;


-- TCP서버 삭제하기
DELIMITER $$
CREATE PROCEDURE SP_DELETE_TCPSERVER (i_id INT)
BEGIN
	DELETE FROM TB_TCP_SERVER WHERE id = i_id;
END $$
DELIMITER ;


