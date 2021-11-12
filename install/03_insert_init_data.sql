-- 메인관리자
INSERT INTO TB_ADMIN(userid, passwd, nick, phone) VALUES('admin', SHA2('admin123', 256), '메인관리자', '');


-- 설정
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('ALARM_USE_YN', '장애알림 사용여부', 'N', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('ERR_CNT_FOR_ALARM', '알림을 위한 장애횟수', '3', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('SLACK_USE_YN', 'Slack 사용여부', 'N', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('SLACK_CHANNEL', 'Slack 채널', '', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('SLACK_TOKEN', 'Slack Webhook URL', '', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('SLACK_USERNAME', 'Slack 사용자이름', '', '');
INSERT INTO TB_SETTING(code, name, value, memo) VALUES('ERR_LOG_SAVE_DAYS', '로그 저장일', '600', '');


