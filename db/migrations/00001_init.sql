-- +goose Up
-- +goose StatementBegin
CREATE TABLE place (
	id VARCHAR(255) PRIMARY KEY,
	name TEXT NOT NULL,
	rating FLOAT DEFAULT 0 NOT NULL,
	weighted_rating FLOAT DEFAULT 0 NOT NULL,
	user_rating_count INT DEFAULT 0 NOT NULL,
	summary VARCHAR(4096),
	last_scraped TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	image_url VARCHAR(1024),
	recompute_stats BOOLEAN DEFAULT FALSE NOT NULL,
	primary_type VARCHAR(64),
	business_summary VARCHAR(4096),
	price_range INT,
	earliest_review_date TIMESTAMPTZ,
	lat FLOAT,
	lng FLOAT
);
CREATE TABLE field_category (
	id SERIAL PRIMARY KEY,
	name VARCHAR(64) NOT NULL UNIQUE
);
CREATE TABLE field (
	id SERIAL PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	category_id INTEGER NOT NULL REFERENCES field_category (id),
	UNIQUE (name, category_id)
);
CREATE TABLE place_field (
	place_id VARCHAR(64) NOT NULL REFERENCES place (id),
	field_id INTEGER NOT NULL REFERENCES field (id),
	PRIMARY KEY (place_id, field_id)
);
CREATE TABLE place_keyword (
	place_id VARCHAR NOT NULL REFERENCES place (id) ON DELETE CASCADE,
	keyword VARCHAR NOT NULL,
	COUNT INTEGER DEFAULT 1 NOT NULL,
	PRIMARY KEY (place_id, keyword)
);
CREATE TABLE "user" (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	photo_uri VARCHAR(256),
	review_count INTEGER DEFAULT 0 NOT NULL,
	photo_count INTEGER DEFAULT 0 NOT NULL,
	rating_count INTEGER DEFAULT 0 NOT NULL,
	is_local_guide BOOLEAN DEFAULT (FALSE) NOT NULL,
	score INTEGER DEFAULT 0 NOT NULL,
	long_review_count INTEGER DEFAULT 0 NOT NULL
);
CREATE TABLE review (
	id VARCHAR(255) PRIMARY KEY,
	user_id VARCHAR(64) NOT NULL REFERENCES "user" (id),
	rating INTEGER NOT NULL,
	text VARCHAR(4096) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	weight INTEGER NOT NULL,
	place_id VARCHAR(64) NOT NULL REFERENCES place (id),
	price_range INTEGER,
	summary VARCHAR(512),
	business_summary VARCHAR(512),
	CONSTRAINT ck_review_rating CHECK (
		rating > 0
		AND rating < 6
	),
	CONSTRAINT ck_review_weight CHECK (
		weight > 0
		AND weight <= 1000
	)
);
CREATE TABLE review_image (
	review_id VARCHAR(64) NOT NULL REFERENCES review (id),
	image_url VARCHAR(1024) NOT NULL,
	PRIMARY KEY (review_id, image_url)
);
CREATE TABLE review_reply (
	review_id VARCHAR(64) NOT NULL PRIMARY KEY REFERENCES review (id),
	text VARCHAR(4096) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- 
-- +goose StatementBegin
DROP TABLE review_reply;
DROP TABLE review_image;
DROP TABLE review;
DROP TABLE "user";
DROP TABLE place_keyword;
DROP TABLE place_field;
DROP TABLE field;
DROP TABLE field_category;
DROP TABLE place;
-- +goose StatementEnd