-- 관리자 목록
DELIMITER $$
CREATE PROCEDURE SP_LIST_ADMIN()
BEGIN
	SELECT id, userid, passwd, nick, phone FROM TB_ADMIN;
END $$
DELIMITER ;


-- 관리자인가?
DELIMITER $$
CREATE PROCEDURE IS_ADMIN(
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



