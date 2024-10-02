-- Создание таблицы music_library
CREATE TABLE music_library (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,                  -- Идентификатор записи
    song_name VARCHAR(63) NOT NULL,                         -- Название песни
    group_name VARCHAR(63) NOT NULL,                        -- Название группы
    release_date DATE,                                      -- Дата релиза
    song_text TEXT,                                         -- Текст песни
    link_url VARCHAR(127)                                   -- Ссылка на песню
);