CREATE TABLE users (
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR(255) NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255) NOT NULL,          -- usually stores hashed password
    is_author  BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Optional but very commonly used:
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);