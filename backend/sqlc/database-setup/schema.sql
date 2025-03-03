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