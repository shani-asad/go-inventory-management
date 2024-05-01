CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    match_cat_id INT NOT NULL,
    user_cat_id INT NOT NULL,
    message VARCHAR(120) NOT NULL CHECK (LENGTH(message) >= 5 AND LENGTH(message) <= 120),
    FOREIGN KEY (match_cat_id) REFERENCES cats(id),
    FOREIGN KEY (user_cat_id) REFERENCES cats(id)
);