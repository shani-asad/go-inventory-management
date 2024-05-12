-- DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    sku VARCHAR(255),
    category VARCHAR(50),
    image_url VARCHAR(255),
    notes VARCHAR(255),
    price INT,
    stock INT,
    location VARCHAR(255),
    is_available BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);