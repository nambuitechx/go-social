-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(36) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);
CREATE TABLE posts(
    id VARCHAR(36) PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    CONSTRAINT FK_post_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE posts;
DROP TABLE users;
-- +goose StatementEnd
