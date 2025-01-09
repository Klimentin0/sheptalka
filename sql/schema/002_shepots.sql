-- +goose Up
CREATE TABLE shepots (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    body TEXT NOT NULL,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE shepots;
