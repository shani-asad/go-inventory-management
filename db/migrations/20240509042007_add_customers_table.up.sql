-- DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(255),
    name VARCHAR(255)
);
