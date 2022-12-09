CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY NOT NULL, 
    full_name  VARCHAR(50) NOT NULL,
    bio TEXT NOT NULL,
    phone_number VARCHAR(20),
    password TEXT NOT NULL
);

