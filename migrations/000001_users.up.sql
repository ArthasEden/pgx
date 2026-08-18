CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INTEGER NOT NULL,
    phone_number VARCHAR(20),
    is_active BOOLEAN,
    created_at TIMESTAMPTZ
);