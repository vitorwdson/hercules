CREATE TABLE sessions (
    id UUID DEFAULT gen_random_uuid () NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id),
    created_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE
);
