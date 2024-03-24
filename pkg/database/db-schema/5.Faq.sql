CREATE TABLE Faq (
    id INTEGER PRIMARY KEY,
    collapseTarget TEXT,
    title TEXT,
    content TEXT,
    category_id integer,
    referrer_id INTEGER, -- This is the product id
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES Categories(id)
);

INSERT INTO Faq (id, collapseTarget, title, content, category_id, referrer_id) VALUES
(1, 'One', 'What do I need to hire a car?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.', 2, 1),
(2, 'Two', 'How old do I have to be to rent a car?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.', 2, 2),
(3, 'Three', 'Can I book a hire car for someone else?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.', 2, 3),
(4, 'Four', 'How do I find the cheapest car hire deal?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.', 2, 1),
(5, 'Five', 'What should I look for when Im choosing a car?', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.', 2,1);
