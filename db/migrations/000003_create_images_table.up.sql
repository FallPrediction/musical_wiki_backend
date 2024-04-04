CREATE TABLE images (
	id SERIAL PRIMARY KEY,
	image_name VARCHAR(255) NOT NULL,
	mime VARCHAR(255),
	actor_id INTEGER NOT NULL,
	image_type VARCHAR(50) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
)
;
COMMENT ON COLUMN images.image_name IS '圖片名稱';
COMMENT ON COLUMN images.image_type IS '圖片類型(AVATAR, GALLERY)';