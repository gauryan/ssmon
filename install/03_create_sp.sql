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


