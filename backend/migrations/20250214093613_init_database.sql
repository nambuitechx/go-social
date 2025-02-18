-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(36) UNIQUE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
CREATE TABLE posts(
    id VARCHAR(36) PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT FK_post_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE posts;
DROP TABLE users;
-- +goose StatementEnd
