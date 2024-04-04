CREATE TABLE IF NOT EXISTS actors (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	translated_name VARCHAR(255),
	nick_name VARCHAR(255),
	nationality VARCHAR(255),
	born VARCHAR(255),
    content TEXT,
    socials JSON,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
)
;
COMMENT ON COLUMN actors.name IS '本名';
COMMENT ON COLUMN actors.translated_name IS '譯名';
COMMENT ON COLUMN actors.nick_name IS '暱稱，以逗號分隔';
COMMENT ON COLUMN actors.nationality IS '國籍';
COMMENT ON COLUMN actors.born IS '出生';
COMMENT ON COLUMN actors.content IS '內文';
COMMENT ON COLUMN actors.socials IS '使用的社群軟體';