CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    [name] TEXT,
    [email] TEXT,
    [password] TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email, password) VALUES
('User1', 'user1@example.com', 'password1'),
('User2', 'user2@example.com', 'password2'),
('Admin', 'paw1a@yandex.ru', '123'),
('Admin 2', 'admin@admin.com', 'admin');