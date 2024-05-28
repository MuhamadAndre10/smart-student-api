CREATE TABLE IF NOT EXISTS users (
  id SERIAL,
    full_name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    photo VARCHAR(100),
    user_active BOOLEAN NOT NULL DEFAULT true,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);