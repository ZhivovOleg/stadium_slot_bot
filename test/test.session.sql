--SELECT *
--FROM stadium_slot_users;
--SELECT *
--FROM stadium_slot_racers;
--SELECT *
--FROM stadium_slot_info;
--DELETE FROM stadium_slot_racers;
--DELETE FROM stadium_slot_users;
--INSERT INTO stadium_slot_users (id, first_name, last_name, birthdate, is_admin, chat_id) overriding system value VALUES (1, 'Олег', 'Живов', '1982-04-06', TRUE, 188941082);
--SELECT settings->'d' FROM settings WHERE servicename = 'testService';
--SELECT settings FROM settings WHERE serviceName = 'testService';
--SELECT settings->'c'->'ca' FROM settings WHERE servicename = 'testService';
--SELECT * FROM settings;
--DELETE from settings WHERE servicename = '';
--UPDATE settings SET settings = '{"a":"a1"}' WHERE servicename = 'newService';
--UPDATE settings SET settings = jsonb_set(settings, '{b, ba}', '"zzzzzzz"', true) WHERE servicename = 'newService';
--UPDATE settings SET settings = jsonb_set(settings, '{c,ca}', '10', FALSE) WHERE servicename = 'testService';
--UPDATE settings SET settings = settings::jsonb #- '{c,ca}' WHERE servicename = 'testService';
--SELECT *
--FROM stadium_slot_races
--WHERE race_date::date < NOW()::date
--ORDER BY race_date
--LIMIT 1;
--SELECT *
--FROM stadium_slot_info
--WHERE message_name = 'stadium_coordinates'
--    AND message_name = 'stadium_navi';
--UPDATE stadium_slot_info
--SET message_text = E'Запись на прокат доступна следующими способами:\r\nТелефон: +79093081081\r\nWhatsApp: <a href="https://wa.me/79093081081">перейти в чат</a>\r\nTelegram: @Gerapolika\r\nСайт: https://enduroshkolakzn.ru'
--WHERE message_name = 'stadium_rent';
SELECT race_date,
    race_name,
    race_info
FROM stadium_slot_races
WHERE race_date::date < NOW()::date