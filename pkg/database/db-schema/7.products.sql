CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    [name] TEXT,
    [description] TEXT,
    price REAL,
    category_id integer,
     Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES Categories(id)
);

INSERT INTO products ([name], [description], price, category_id) VALUES
('Product1', 'Description1', 500.00, 3),
('Product2', 'Description2', 750.00, 3),
('Product3', 'Description3', 1000.00, 3),
('Product4', 'Description4', 1500.00, 3),
('Product5', 'Description5', 2000.00, 3);
