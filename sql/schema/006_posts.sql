-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT ,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feed_id uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;

