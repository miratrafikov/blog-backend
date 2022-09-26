CREATE ROLE blog_admin;

SET ROLE blog_admin;

CREATE DATABASE blog OWNER blog_admin;

CREATE TABLE tag (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    color VARCHAR(6)
);

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    date_created TIMESTAMP,
    content TEXT
);