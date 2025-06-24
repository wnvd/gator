-- +goose Up

CREATE TABLE feeds (
    id                      UUID PRIMARY KEY,
    name                    TEXT NOT NULL,
    url                     TEXT NOT NULL UNIQUE,
    user_id                 UUID NOT NULL,
    FOREIGN KEY(user_id)    REFERENCES users(id) ON DELETE CASCADE,
    created_at              TIMESTAMP NOT NULL,
    updated_at              TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE feeds;
