CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    description TEXT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO products (name, price, description, quantity) VALUES
    ('Aqua Bottle 600ml', 4000, 'Mineral water bottle 600ml', 200),
    ('Indomie Goreng', 3000, 'Indonesian instant noodles', 500),
    ('Kopi ABC Sachet', 1500, 'Instant coffee sachet', 100),
    ('SilverQueen Chocolate', 12000, 'Milk chocolate bar with cashew', 50),
    ('Teh Botol Sosro', 5000, 'Sweetened bottled tea drink', 150);

