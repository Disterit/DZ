CREATE TABLE IF NOT EXISTS recipe  (
    id SERIAL PRIMARY KEY,
    name_recipe VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS ingredient  (
    id SERIAL PRIMARY KEY,
    id_recipe INT REFERENCES recipe(id) ON DELETE CASCADE NOT NULL,
    name_ingredient VARCHAR(50) NOT NULL
);