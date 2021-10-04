-- 메인관리자
INSERT INTO TB_ADMIN(userid, passwd, nick, phone) VALUES('admin', SHA2('admin123', 256), '메인관리자', '');


-- 설정
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('ALARM_USE_YN', '장애알림 사용여부', '', 'N');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('ERR_CNT_FOR_ALARM', '알림을 위한 장애횟수', '', '3');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('SLACK_USE_YN', 'Slack 사용여부', '', 'N');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('SLACK_CHANEL', 'Slack 채널', '', '');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('SLACK_TOKEN', 'Slack 토큰', '', '');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('SLACK_USERNAME', 'Slack 사용자이름', '', '');
INSERT INTO TB_SETTING(code, name, memo, value) VALUES('SLACK_IMAGE', 'Slack 이미지 URL', '', '');


