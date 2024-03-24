CREATE TABLE reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    [text] TEXT,
    rating INTEGER,
    product_ID INTEGER,
    user_ID INTEGER,
    username TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_ID) REFERENCES products(id),
    FOREIGN KEY (user_ID) REFERENCES users(id)
);

INSERT INTO reviews ([text], rating, product_ID, user_ID, username) VALUES
('Great product!', 5, 1, 1, 'User1'),
('Not bad, could be better', 3, 2, 2, 'User2');
-- Add more review data as needed;

