CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
    full_name VARCHAR(100),
    username VARCHAR(100),
    email VARCHAR(50) NOT NULL UNIQUE,
    photo VARCHAR(100),
    user_active BOOLEAN NOT NULL DEFAULT true,
    verified_email BOOLEAN NOT NULL DEFAULT false,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);