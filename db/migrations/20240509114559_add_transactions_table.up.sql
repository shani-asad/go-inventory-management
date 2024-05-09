CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_id SERIAL REFERENCES customers(id),
    paid INT,
    change INT
);