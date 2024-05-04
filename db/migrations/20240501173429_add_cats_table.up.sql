DROP TYPE IF EXISTS cat_breed;
CREATE TYPE cat_breed AS ENUM (
    'Persian',
    'Maine Coon',
    'Siamese',
    'Ragdoll',
    'Bengal',
    'Sphynx',
    'British Shorthair',
    'Abyssinian',
    'Scottish Fold',
    'Birman'
);

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    user_id INT,
    name VARCHAR(255) NOT NULL,
    race cat_breed NOT NULL,
    sex VARCHAR(10) NOT NULL CHECK (sex IN ('male', 'female')),
    age_in_month INT NOT NULL,
    description VARCHAR(20) NOT NULL,
    image_urls TEXT[] NOT NULL,
    has_matched BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);