----------------Slide Galory--------------------------------
CREATE TABLE GalleryImages (
    id INTEGER PRIMARY KEY,
    category_id integer,
    referrer_id INTEGER, -- This is the product id
    img TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (category_id) REFERENCES Categories(id)
);

INSERT INTO GalleryImages (category_id, referrer_id,  img) VALUES 
(2,   1,'/img/tours/new/1.png'),
(2,   2,'/img/tours/new/2.png'),
(2,   2,'/img/tours/new/1.png'),
(2,   2,'/img/tours/new/3.png'),
(2,   3,'/img/tours/new/3.png'),
(2,   4,'/img/tours/new/4.png'),
(2,   5,'/img/tours/new/5.png'),
(2,   6,'/img/tours/new/6.png'),
(2,   6,'/img/tours/new/7.png'),
(2,   6,'/img/tours/new/8.png'),
(2,   7,'/img/tours/new/7.png'),
(2,   8,'/img/tours/new/8.png'),
(2,   9,'/img/tours/new/9.png');