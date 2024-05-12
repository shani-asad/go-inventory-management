CREATE TABLE transaction_products (
    transaction_id SERIAL REFERENCES transactions(id),
    product_id SERIAL,
    quantity INT
);