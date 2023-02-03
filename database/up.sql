DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    lastname VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL, 
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    status BOOLEAN NOT NULL,
    pass VARCHAR(255) NOT NULL,
    position BOOLEAN NOT NULL
);
