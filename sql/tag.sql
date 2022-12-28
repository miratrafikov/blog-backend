CREATE TABLE tag (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    color VARCHAR(6)
);