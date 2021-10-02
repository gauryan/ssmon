-- SSMON DB와 계정 생성
-- 원하는 이름과 비밀번호로 변경하세요!
CREATE DATABASE ssmon DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
CREATE USER 'ssmon'@'%' IDENTIFIED BY 'ssmon123';
GRANT ALL PRIVILEGES ON ssmon.* TO 'ssmon'@'%' WITH GRANT OPTION;
CREATE USER 'ssmon'@'localhost' IDENTIFIED BY 'ssmon123';
GRANT ALL PRIVILEGES ON ssmon.* TO 'ssmon'@'localhost' WITH GRANT OPTION;
FLUSH privileges;

