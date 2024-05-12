CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_id SERIAL REFERENCES customers(id),
    paid INT,
    change INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);