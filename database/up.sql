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

DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    user_id SERIAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    title VARCHAR(32) NOT NULL, 
    description VARCHAR(255) NOT NULL,
    image bytea NOT NULL,
    status BOOLEAN NOT NULL,
    date_finish timestamp
);

