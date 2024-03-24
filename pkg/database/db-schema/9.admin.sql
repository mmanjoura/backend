CREATE TABLE admins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    [name] TEXT,
    email TEXT,
    [password] TEXT,
       Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO admins ([name], email, [password]) VALUES
('Admin', 'admin@example.com', 'adminpassword'),
('Admin2', 'admin2@example.com', 'adminpassword2');
