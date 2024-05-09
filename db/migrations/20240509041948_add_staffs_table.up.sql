-- DROP TABLE IF EXISTS staffs;

CREATE TABLE staffs (
    id SERIAL PRIMARY KEY,
    phoneNumber VARCHAR(255),
    name VARCHAR(255),
    password VARCHAR(255)
);