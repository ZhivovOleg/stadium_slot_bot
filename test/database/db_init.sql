----------
-- SCHEMAS
----------
CREATE TABLE stadium_slot_users (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    birthdate date NOT NULL,
    is_admin boolean NOT NULL,
    chat_id int NOT NULL UNIQUE,
    UNIQUE (first_name, last_name)
);
CREATE TABLE stadium_slot_racers (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id int NOT NULL,
    racer_number int,
    raсer_team varchar,
    CONSTRAINT fk_racers_users FOREIGN KEY(user_id) REFERENCES stadium_slot_users(id) ON DELETE CASCADE
);
CREATE TABLE stadium_slot_races (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    race_date date NOT NULL,
    race_name varchar NOT NULL,
    race_info varchar NOT NULL
);
CREATE TABLE stadium_slot_race_results (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    race_id int NOT NULL,
    racer_id int NOT NULL,
    result_time bigint,
    result_position int,
    UNIQUE (race_id, racer_id),
    CONSTRAINT fk_result_race FOREIGN KEY(race_id) REFERENCES stadium_slot_races(id) ON DELETE CASCADE,
    CONSTRAINT fk_result_racer FOREIGN KEY(racer_id) REFERENCES stadium_slot_racers(id) ON DELETE CASCADE
);
CREATE TABLE stadium_slot_info (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    message_name varchar NOT NULL,
    message_text varchar NOT NULL
);
CREATE TABLE stadium_slot_coaches (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id int NOT NULL,
    contacts varchar NOT NULL,
    CONSTRAINT fk_coaches_users FOREIGN KEY(user_id) REFERENCES stadium_slot_users(id) ON DELETE CASCADE
);
CREATE TABLE stadium_slot_coaches_shedule (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    coach_id int NOT NULL,
    slot timestamp NOT NULL,
    approved boolean NOT NULL,
    UNIQUE (coach_id, slot),
    CONSTRAINT fk_shedule_coach FOREIGN KEY(coach_id) REFERENCES stadium_slot_coaches(id) ON DELETE CASCADE
);
---------
-- DATA
---------
INSERT INTO stadium_slot_users (
        id,
        first_name,
        last_name,
        birthdate,
        is_admin,
        chat_id
    ) OVERRIDING SYSTEM VALUE
VALUES (
        1,
        'Олег',
        'Живов',
        '1982-04-06',
        TRUE,
        188941082
    );
INSERT INTO stadium_slot_racers (user_id, racer_number, raсer_team)
VALUES (1, 139, 'SLOT');
INSERT INTO stadium_slot_info (message_name, message_text)
VALUES (
        'stadium_coordinates',
        '<pre>55.860649,49.245509</pre>'
    ),
    (
        'stadium_navi',
        '<a href="https://yandex.ru/maps/-/CDRer60a">яндекс-карты</a>'
    ),
    (
        'stadium_rent',
        E'Запись на прокат доступна следующими способами:\r\nТелефон: +79093081081\r\nWhatsApp: <a href = "https://wa.me/79093081081">перейти в чат</a>\r\nTelegram: @Gerapolika\r\nСайт: https://enduroshkolakzn.ru'
    );
INSERT INTO stadium_slot_coaches (user_id, contacts)
VALUES (1, 'Telegram: @AidahoOleg');
INSERT INTO stadium_slot_races (race_date, race_name, race_info)
VALUES (
        '01-05-2024',
        'Кубок стадиона SLOT',
        E'<a href="https://disk.yandex.ru/i/A6CQFMquPQtHyw">ссылка на регламент</a>\r\n<a href="https://marshalone.ru/card/70d669b0-7e6b-41b0-b213-4c2a8c5ffa52">РЕГИСТРАЦИЯ</a>\r\n<a href="https://t.me/stadium_slot/55">Схемы трасс все классов</a>'
    )