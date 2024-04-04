CREATE TABLE IF NOT EXISTS credits (
	id SERIAL PRIMARY KEY,
	time VARCHAR(255),
	place VARCHAR(255),
	musical VARCHAR(255) NOT NULL,
	character VARCHAR(255) NOT NULL,
	actor_id INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT credits_actor_id_foreign FOREIGN KEY (actor_id) REFERENCES public.actors (id) ON UPDATE CASCADE ON DELETE CASCADE
)
;
CREATE INDEX credits_actor_id_index ON credits (actor_id);
COMMENT ON COLUMN credits.musical IS '音樂劇名稱';
COMMENT ON COLUMN credits.character IS '飾演角色';
COMMENT ON COLUMN credits.place IS '地點';
COMMENT ON COLUMN credits.time IS '時間';