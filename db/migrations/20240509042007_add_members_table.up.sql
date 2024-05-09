-- DROP TABLE IF EXISTS members;

CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    phoneNumber VARCHAR(255),
    name VARCHAR(255)
);
