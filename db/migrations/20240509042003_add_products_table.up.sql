-- DROP TABLE IF EXISTS products;

CREATE TYPE product_category AS ENUM ('Clothing', 'Accessories', 'Footwear', 'Beverages');

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    sku VARCHAR(255),
    category product_category,
    imageUrl VARCHAR(255),
    notes VARCHAR(255),
    price INT,
    stock INT,
    location VARCHAR(255),
    isAvailable BOOLEAN
);
