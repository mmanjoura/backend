CREATE TABLE Categories (
    ID                 INTEGER PRIMARY KEY,
    [Name]               TEXT    NOT NULL,
    [Description]        TEXT,
    Created_At         DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At         DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Categories (ID, Name, Description) VALUES (1, 'Flights', 'Flights Category');
INSERT INTO Categories (ID, Name, Description) VALUES (2, 'Tours', 'Tours Category');
INSERT INTO Categories (ID, Name, Description) VALUES (3, 'Hotels', 'Hotels Category');
INSERT INTO Categories (ID, Name, Description) VALUES (4, 'Rentals', 'Rentals Category');
INSERT INTO Categories (ID, Name, Description) VALUES (5, 'Cars', 'Cars Category');
INSERT INTO Categories (ID, Name, Description) VALUES (6, 'Golfs', 'Golfs Category');
INSERT INTO Categories (ID, Name, Description) VALUES (7, 'Flights', 'Flights Category');
INSERT INTO Categories (ID, Name, Description) VALUES (8, 'Activities', 'Activities Category');