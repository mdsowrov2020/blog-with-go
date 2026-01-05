-- File: migrations/000002_add_status_to_posts.sql
-- +migrate Up
ALTER TABLE posts ADD COLUMN status VARCHAR(255) DEFAULT 'draft';

-- +migrate Down
ALTER TABLE posts DROP COLUMN status;