CREATE TABLE place (
    id VARCHAR(255) PRIMARY KEY,
    name TEXT NOT NULL,
    rating FLOAT DEFAULT 0 NOT NULL,
    weighted_rating FLOAT DEFAULT 0 NOT NULL,
    user_rating_count INT DEFAULT 0 NOT NULL,
    summary TEXT,
    last_scraped TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    image_url TEXT,
    recompute_stats BOOLEAN DEFAULT FALSE NOT NULL,
    primary_type VARCHAR(64),
    business_summary TEXT,
    price_range INT,
    earliest_review_date TIMESTAMP,
    lat FLOAT,
    lng FLOAT
);