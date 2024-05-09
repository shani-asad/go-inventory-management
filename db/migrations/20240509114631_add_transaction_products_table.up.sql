CREATE TABLE transaction_products (
    transaction_id SERIAL REFERENCES transactions(id),
    product_id SERIAL REFERENCES products(id),
    quantity INT
);