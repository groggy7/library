CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    year INT NOT NULL
);