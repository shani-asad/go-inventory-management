-- DROP TABLE IF EXISTS members;

CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(255),
    name VARCHAR(255)
);
