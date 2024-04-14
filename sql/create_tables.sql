CREATE SCHEMA IF NOT EXISTS banners_manager;

CREATE TABLE IF NOT EXISTS banners_manager.banners
(
    id         BIGSERIAL PRIMARY KEY,
    feature_id INT   NOT NULL,
    tag_ids    INT[] NOT NULL CHECK (cardinality(tag_ids) > 0),
    content    JSON  NOT NULL,
    is_active  BOOLEAN   DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
)