------------Disable/Enable Foreign Key Checks-----------------
PRAGMA foreign_keys = OFF;
PRAGMA foreign_keys = ON;

------------------- Add Banner Table -------------------------
CREATE TABLE Add_Banners (
    id INTEGER PRIMARY KEY,
    img TEXT,
    title TEXT,
    meta TEXT,
    routerPath TEXT,
    delayAnimation TEXT
);

INSERT INTO Add_Banners (id, img, title, meta, routerPath, delayAnimation)
VALUES (1, '/img/backgrounds/new/1.png', 'Hiking', '', '/', '100');

INSERT INTO Add_Banners (id, img, title, meta, routerPath, delayAnimation)
VALUES (2, '/img/backgrounds/new/3.png', 'Quads', '', '/', '200');

INSERT INTO Add_Banners (id, img, title, meta, routerPath, delayAnimation)
VALUES (3, '/img/backgrounds/new/2.png', 'Golfing', 'Enjoy Summer Deals', '/', '300');


-------------- Destinations Table--------------------
CREATE TABLE Destinations (
    id INTEGER PRIMARY KEY,
    class TEXT,
    title TEXT,
    name TEXT,
    location TEXT,
    travellers TEXT,
    hover TEXT,
    img TEXT,
    city TEXT,
    properties TEXT,
    region TEXT,
    Animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Destinations (class, img, name, properties, animation) VALUES
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/new/2/1.png', 'los angeles', '1714', '200'),
('col-xl-6 col-md-4 col-sm-6', '/img/destinations/new/2/2.png', 'london', '1714', '400'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/new/2/3.png', 'reykjavik', '1714', '600'),
('col-xl-6 col-md-4 col-sm-6', '/img/destinations/new/2/4.png', 'paris', '1714', '200'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/new/2/5.png', 'amsterdam', '1714', '400'),
('col-xl-3 col-md-4 col-sm-6', '/img/destinations/new/2/6.png', 'istanbul', '1714', '600');

-- Hotels Table

CREATE TABLE Hotels (
    id INTEGER PRIMARY KEY,
    category_id INTEGER,
    tag TEXT,
    img TEXT,
    title TEXT,
    location TEXT,
    ratings TEXT,
    Reviews TEXT,
    price TEXT,
    Animation TEXT,
    city TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (category_id) REFERENCES Categories(id)
);

INSERT INTO Hotels (id, category_id, tag, img, title, location, ratings, Reviews, price, Animation, city)
VALUES
(1, 3, 'Breakfast Included', '/img/hotels/new/1.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3014, 72.0, 100, 'new_york'),
(2, 3,'', '/img/hotels/new/2.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 2345, 85.0, 200, 'london'),
(3, 3,'best seller', '/img/hotels/new/3.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.7, 3014, 68.0, 300, 'new_york'),
(4, 3,'top rated', '/img/hotels/new/4.png', 'DoubleTree by Hilton Hotel New York Times Square West', 'Vaticano Prati, Rome', 4.5, 5633, 89.0, 400, 'new_york'),
(5, 3,'Breakfast Included', '/img/hotels/new/5.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.8, 3014, 99.0, 500, 'london'),
(6, 3,'-25% today', '/img/hotels/new/6.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3467, 79.0, 600, 'new_york'),
(7, 3,'best seller', '/img/hotels/new/7.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3014, 88.0, 700, 'new_york'),
(8, 3,'top rated', '/img/hotels/new/8.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.9, 7654, 68.0, 800, 'london'),
(9, 3,'Breakfast Included', '/img/hotels/new/9.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3245, 72.0, 900, 'new_york'),
(10, 3,'', '/img/hotels/new/10.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3432, 85.0, 1000, 'paris'),
(11, 3,'best seller', '/img/hotels/new/11.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.7, 3014, 68.0, 1100, 'new_york'),
(12, 3,'top rated', '/img/hotels/new/12.png', 'DoubleTree by Hilton Hotel New York Times Square West', 'Vaticano Prati, Rome', 4.5, 2343, 89.0, 1200, 'paris'),
(13, 3,'Breakfast Included', '/img/hotels/new/13.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.8, 3014, 99.0, 500, 'new_york'),
(14, 3,'-25% today', '/img/hotels/new/14.png', 'The Montcalm At Brewery London City', 'Westminster Borough, London', 4.7, 3467, 79.0, 600, 'istanbul'),
(15, 3,'best seller', '/img/hotels/new/15.png', 'Staycity Aparthotels Deptford Bridge Station', 'Ciutat Vella, Barcelona', 4.8, 3014, 88.0, 700, 'istanbul'),
(16, 3,'top rated', '/img/hotels/new/16.png', 'The Westin New York at Times Square West', 'Manhattan, New York', 4.9, 7654, 68.0, 800, 'new_york');

----Hotel Infos Table -------

CREATE TABLE Images (
    ID          INTEGER  PRIMARY KEY,
    hotel_id    INTEGER,
    activity_id INTEGER,
    car_id      INTEGER,
    golf_id     INTEGER,
    rental_id   INTEGER,
    tour_id     INTEGER,
    Name        TEXT,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        hotel_id
    )
    REFERENCES hotels (id),
    FOREIGN KEY (
        activity_id
    )
    REFERENCES Activities (id),
    FOREIGN KEY (
        car_id
    )
    REFERENCES Cars (id),
    FOREIGN KEY (
        golf_id
    )
    REFERENCES Golfs (id),
    FOREIGN KEY (
        rental_id
    )
    REFERENCES Rentals (id),
    FOREIGN KEY (
        tour_id
    )
    REFERENCES Tours (id) 
);


INSERT INTO Images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  name) VALUES 
-- Insert Hotels Slide Images
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/1.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/2.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/1.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/3.png');



CREATE TABLE Hotel_Infos (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    Room_Type TEXT, 
    overview TEXT,
    img TEXT,
    price TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (hotel_id) REFERENCES Hotels(id)
);

INSERT INTO Hotel_Infos (hotel_id, Room_Type, overview, img, price)
VALUES (1, 'Standard Twin Room', '
        You can directly book the best price if your travel dates are available,
        all discounts are already included. In the following house description
        you will find all information about our listing.
        <br />
        <br />
        2-room terraced house on 2 levels. Comfortable and cosy furnishings: 1
        room with 1 french bed and radio. Shower, sep. WC. Upper floor: (steep
        stair) living/dining room with 1 sofabed (110 cm, length 180 cm), TV.
        Exit to the balcony. Small kitchen (2 hot plates, oven,
      ', '/img/hotels/new/1.png', '$200 per night');

INSERT INTO Hotel_Infos (hotel_id, Room_Type, overview, img, price)
VALUES (1, 'Deluxe King Room', '
        You can directly book the best price if your travel dates are available,
        all discounts are already included. In the following house description
        you will find all information about our listing.
        <br />
        <br />
        2-room terraced house on 2 levels. Comfortable and cosy furnishings: 1
        room with 1 french bed and radio. Shower, sep. WC. Upper floor: (steep
        stair) living/dining room with 1 sofabed (110 cm, length 180 cm), TV.
        Exit to the balcony. Small kitchen (2 hot plates, oven,
      ', '/img/hotels/new/2.png', '$200 per night');




---Room Types Table

CREATE TABLE Room_Types (
    id INTEGER PRIMARY KEY,
    RoomType TEXT
   
);
INSERT INTO Room_Types (RoomType) VALUES ('Standard Twin Room');
INSERT INTO Room_Types (RoomType) VALUES ('Deluxe King Room');

CREATE TABLE Facilities (
    id INTEGER PRIMARY KEY,
    category_id INTEGER,
    product_id INTEGER,
    class_name TEXT,
    facility_title TEXT,
    facility_meta TEXT,
    exist INTEGER,
    isHilight INTEGER,
    FOREIGN KEY (category_id) REFERENCES Categories(id)
);
-- Example INSERTs for Hotel_Facilities
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, exist, isHilight) VALUES (3, 1, 'icon-no-smoke', "", 'Non-smoking rooms', "", 1, 0);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, exist, isHilight) VALUES (3, 1, 'icon-wifi', "", 'Free WiFi', "", 1, 0);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, exist, isHilight) VALUES (3, 1, 'icon-parking', "", 'Parking', "", 1, 0);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, exist, isHilight) VALUES (3, 1, 'icon-kitchen', "", 'Kitchen', "", 1, 0);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, exist, isHilight) VALUES (3, 1, 'icon-living-room', "", 'Living Area', "", 1, 0);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, facility_meta, value, isHilight) VALUES (3, 1, 'icon-shield', "", 'Safety & security', "", 1, 0);

INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (3, 1, 'icon-city', "", 'In London City Centre', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (3, 1, 'icon-airplane', "", 'Airport transfer', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (3, 1, 'icon-bell-ring', "", 'Front desk [24-hour]', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (3, 1, 'icon-tv', "", 'Premium TV channels', "", 1, 1);

INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (6, 1, 'icon-clock', "", 'In London City Centre', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (6, 1, 'icon-customer', "", 'Airport transfer', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (6, 1, 'icon-route', "", 'Front desk [24-hour]', "", 1, 1);
INSERT INTO Facilities (category_id, product_id, class_name, facility_title, value, facility_meta, isHilight) VALUES (6, 1, 'icon-access-denied', "", 'Premium TV channels', "", 1, 1);



CREATE TABLE Hotel_Benefits (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    class_name TEXT,
    Benefit_name TEXT,
    exist INTEGER,
    FOREIGN KEY (hotel_id) REFERENCES Hotels(id)
);

-- Example INSERTs for Hotel_Benefits
INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (1, 'icon-check', 'Pay at the hotel', 1);
INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (1, 'icon-check', 'Pay nothing until March 30, 2022', 1);
INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (1, 'icon-check', 'Free cancellation before April 1, 2022', 1);

INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (2, 'icon-check', 'Pay at the hotel', 1);
INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (2, 'icon-check', 'Pay nothing until March 30, 2022', 1);
INSERT INTO Hotel_Benefits (hotel_id, class_name, Benefit_name, exist) VALUES (2, 'icon-check', 'Free cancellation before April 1, 2022', 1);





------------------Slide Images Table ------------------------------
CREATE TABLE slide_images (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    activity_id INTEGER,
    car_id INTEGER,
    golf_id INTEGER,
    rental_id INTEGER,
    tour_id INTEGER,
    img TEXT, 
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id),
    FOREIGN KEY (activity_id) REFERENCES Activities(id),
    FOREIGN KEY (car_id) REFERENCES Cars(id),
    FOREIGN KEY (golf_id) REFERENCES Golfs(id),
    FOREIGN KEY (rental_id) REFERENCES Rentals(id),
    FOREIGN KEY (tour_id) REFERENCES Tours(id)
);


-- Insert Hotels Slide Images

INSERT INTO slide_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/1.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/2.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/3.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/4.png'),
(2, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/5.png'),
(2, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/6.png'),
(2, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/7.png'),
(2, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/8.png'),
(3, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/9.png'),
(3, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/10.png'),
(3, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/11.png'),
(3, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/12.png'),
(4, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/13.png'),
(4, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/14.png'),
(4, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/15.png'),
(4, NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/16.png');
----Insert Acitivities Slide Images
INSERT INTO slide_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/1.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/2.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/3.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/4.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/5.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/6.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/7.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/8.png'),
(NULL,  3, NULL, NULL, NULL, NULL,'/img/activities/new/9.png'),
(NULL,  3, NULL, NULL, NULL, NULL,'/img/activities/new/10.png'),
(NULL,  3, NULL, NULL, NULL, NULL,'/img/activities/new/11.png'),
(NULL,  3, NULL, NULL, NULL, NULL,'/img/activities/new/12.png'),
(NULL,  4, NULL, NULL, NULL, NULL,'/img/activities/new/13.png'),
(NULL,  4, NULL, NULL, NULL, NULL,'/img/activities/new/14.png'),
(NULL,  4, NULL, NULL, NULL, NULL,'/img/activities/new/15.png'),
(NULL,  4, NULL, NULL, NULL, NULL,'/img/activities/new/16.png'),
----Insert Cars Slide Images
(NULL,  NULL, 1, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 3, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 4, NULL, NULL, NULL,'/img/cars/new/4.png'),
(NULL,  NULL, 5, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 6, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 6, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 6, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 7, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 8, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 8, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 8, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 9, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 10, NULL, NULL, NULL,'/img/cars/new/4.png'),
(NULL,  NULL, 11, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 12, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 13, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 14, NULL, NULL, NULL,'/img/cars/new/4.png'),
(NULL,  NULL, 15, NULL, NULL, NULL,'/img/cars/new/1.png'),
----Insert Golf Slide Images
INSERT INTO slide_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/4.png');
----Insert Rentals Slide Images
INSERT INTO slide_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/new/1.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/new/2.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/new/3.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/new/4.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/new/5.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/new/6.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/new/7.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/new/8.png'),
(NULL,  NULL, NULL, NULL, 3, NULL,'/img/rentals/new/9.png'),
(NULL,  NULL, NULL, NULL, 3, NULL,'/img/rentals/new/10.png'),
(NULL,  NULL, NULL, NULL, 3, NULL,'/img/rentals/new/11.png'),
(NULL,  NULL, NULL, NULL, 3, NULL,'/img/rentals/new/12.png'),
(NULL,  NULL, NULL, NULL, 4, NULL,'/img/rentals/new/13.png'),
(NULL,  NULL, NULL, NULL, 4, NULL,'/img/rentals/new/14.png'),
(NULL,  NULL, NULL, NULL, 4, NULL,'/img/rentals/new/15.png'),
(NULL,  NULL, NULL, NULL, 4, NULL,'/img/rentals/new/16.png')

----Insert Tours Slide Images
(NULL,  NULL, NULL, NULL, NULL, 1,'/img/tours/new/1.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/2.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/1.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/3.png'),
(NULL,  NULL, NULL, NULL, NULL, 3,'/img/tours/new/3.png'),
(NULL,  NULL, NULL, NULL, NULL, 4,'/img/tours/new/4.png'),
(NULL,  NULL, NULL, NULL, NULL, 5,'/img/tours/new/5.png'),
(NULL,  NULL, NULL, NULL, NULL, 6,'/img/tours/new/6.png'),
(NULL,  NULL, NULL, NULL, NULL, 6,'/img/tours/new/7.png'),
(NULL,  NULL, NULL, NULL, NULL, 6,'/img/tours/new/8.png'),
(NULL,  NULL, NULL, NULL, NULL, 7,'/img/tours/new/7.png'),
(NULL,  NULL, NULL, NULL, NULL, 8,'/img/tours/new/8.png'),
(NULL,  NULL, NULL, NULL, NULL, 9,'/img/tours/new/9.png');

----------------Slide Galory--------------------------------
CREATE TABLE gallery_images (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    activity_id INTEGER,
    car_id INTEGER,
    golf_id INTEGER,
    rental_id INTEGER,
    tour_id INTEGER,
    img TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id),
    FOREIGN KEY (activity_id) REFERENCES Activities(id),
    FOREIGN KEY (car_id) REFERENCES Cars(id),
    FOREIGN KEY (golf_id) REFERENCES golfs(id),
    FOREIGN KEY (rental_id) REFERENCES Rentals(id),
    FOREIGN KEY (tour_id) REFERENCES Tours(id)
);

-- Insert Hotels Slide Images
INSERT INTO gallery_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/1.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/2.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/3.png'),
(1,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/4.png'),

(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/5.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/6.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/7.png'),
(2,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/8.png'),

(3,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/9.png'),
(3,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/10.png'),
(3,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/11.png'),
(3,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/12.png'),

(4,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/13.png'),
(4,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/14.png'),
(4,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/15.png'),
(4,  NULL, NULL, NULL, NULL, NULL,'/img/hotels/new/16.png');


INSERT INTO gallery_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
----Insert Acitivities Slide Images
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/10.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/11.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/12.png'),
(NULL,  1, NULL, NULL, NULL, NULL,'/img/activities/new/13.png'),

(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/10.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/11.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/12.png'),
(NULL,  2, NULL, NULL, NULL, NULL,'/img/activities/new/13.png');

----Insert Cars Slide Images
INSERT INTO gallery_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(NULL,  NULL, 1, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 1, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 1, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 1, NULL, NULL, NULL,'/img/cars/new/4.png'),

(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 2, NULL, NULL, NULL,'/img/cars/new/4.png'),

(NULL,  NULL, 3, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 3, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 3, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 3, NULL, NULL, NULL,'/img/cars/new/4.png'),

(NULL,  NULL, 4, NULL, NULL, NULL,'/img/cars/new/1.png'),
(NULL,  NULL, 4, NULL, NULL, NULL,'/img/cars/new/2.png'),
(NULL,  NULL, 4, NULL, NULL, NULL,'/img/cars/new/3.png'),
(NULL,  NULL, 4, NULL, NULL, NULL,'/img/cars/new/4.png');

----Insert Golf Slide Images
INSERT INTO gallery_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  gallery_img) VALUES 
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 1, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/5.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/6.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/7.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/8.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/9.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/10.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/11.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/12.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/13.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/14.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/15.png'),
(NULL,  NULL, NULL, 4, NULL, NULL,'/img/golfs/new/16.png')

(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/5.png'),
(NULL,  NULL, NULL, 2, NULL, NULL,'/img/golfs/new/6.png'),

(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/1.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/2.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/3.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/4.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/5.png'),
(NULL,  NULL, NULL, 3, NULL, NULL,'/img/golfs/new/6.png'),

----Insert Rentals Slide Images
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/single/new/1.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/single/new/2.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/single/new/3.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/single/new/4.png'),
(NULL,  NULL, NULL, NULL, 1, NULL,'/img/rentals/single/new/5.png'),

(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/single/new/1.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/single/new/2.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/single/new/3.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/single/new/4.png'),
(NULL,  NULL, NULL, NULL, 2, NULL,'/img/rentals/single/new/5.png');

----Insert Tours Slide Images
INSERT INTO gallery_images (hotel_id, activity_id, car_id, golf_id, rental_id, tour_id,  img) VALUES 
(NULL,  NULL, NULL, NULL, NULL, 1,'/img/tours/new/1.png'),
(NULL,  NULL, NULL, NULL, NULL, 1,'/img/tours/new/2.png'),
(NULL,  NULL, NULL, NULL, NULL, 1,'/img/tours/new/3.png'),
(NULL,  NULL, NULL, NULL, NULL, 1,'/img/tours/new/4.png'),

(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/5.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/6.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/7.png'),
(NULL,  NULL, NULL, NULL, NULL, 2,'/img/tours/new/8.png'),

(NULL,  NULL, NULL, NULL, NULL, 3,'/img/tours/new/9.png'),
(NULL,  NULL, NULL, NULL, NULL, 3,'/img/tours/new/10.png'),
(NULL,  NULL, NULL, NULL, NULL, 3,'/img/tours/new/11.png'),
(NULL,  NULL, NULL, NULL, NULL, 3,'/img/tours/new/12.png'),

(NULL,  NULL, NULL, NULL, NULL, 4,'/img/tours/new/13.png'),
(NULL,  NULL, NULL, NULL, NULL, 4,'/img/tours/new/14.png'),
(NULL,  NULL, NULL, NULL, NULL, 4,'/img/tours/new/15.png'),
(NULL,  NULL, NULL, NULL, NULL, 4,'/img/tours/new/16.png');

---------------- Tours Table -----------------------------
CREATE TABLE Tours (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    price TEXT, 
    tourType TEXT,
    animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Tours (ID, tag, title, price, location, duration, reviews, price, tourType, animation)
VALUES
(1, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 72.0, 'Westminster Borough, London', '16', '3014', '72', 'Full-day Tours', '100'),
(2, '', 'Westminster Walking Tour & Westminster Abbey Entry', 65.0, 'Ciutat Vella, Barcelona', '14', '2045', '65', 'Attractions & Museums', '200'),
(3, 'best seller', 'High-Speed Thames River RIB Golf in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '300'),
(4, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 99.0, 'Vaticano Prati, Rome', '20', '1458', '99', 'Bus Tours', '400'),
(5, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 72.0, 'Westminster Borough, London', '16', '3014', '72', 'Full-day Tours', '100'),
(6, '', 'Westminster Walking Tour & Westminster Abbey Entry', 65.0, 'Ciutat Vella, Barcelona', '14', '2045', '65', 'Attractions & Museums', '200'),
(7, 'best seller', 'High-Speed Thames River RIB Golf in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '300'),
(8, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 99.0, 'Vaticano Prati, Rome', '20', '1458', '99', 'Bus Tours', '400'),
(9, 'best seller', 'High-Speed Thames River RIB Golf in London', 87.0, 'Manhattan, New York', '18', '2163', '87', 'Private and Luxury', '500');


--------------Activities--------------------------------
---------------- Tours Table -----------------------------
CREATE TABLE Activities (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Activities (ID, tag, title, price, location, duration, reviews, ratings, animation)
VALUES
(1, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '100'),
(2, '', 'Westminster Walking Tour & Westminster Abbey Entry', 99.0, 'Ciutat Vella, Barcelona', '14', '2045', '4.82', '200'),
(3, 'best seller', 'High-Speed Thames River RIB Golf in London', 88.0, 'Manhattan, New York', '18', '2163', '4.82', '300'),
(4, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 55.0, 'Vaticano Prati, Rome', '20', '1458', '4.82', '400'),
(5, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '100'),
(6, 'LIKELY TO SELL OUT*', 'Stonehenge, Windsor Castle and Bath with Pub Lunch in Lacock', 67.0, 'Westminster Borough, London', '16', '94', '4.82', '200'),
(7, '', 'Westminster Walking Tour & Westminster Abbey Entry', 99.0, 'Ciutat Vella, Barcelona', '14', '2045', '4.82', '300'),
(8, 'best seller', 'High-Speed Thames River RIB Golf in London', 88.0, 'Manhattan, New York', '18', '2163', '4.82', '400'),
(9, 'top rated', 'Edinburgh Darkside Walking Tour: Mysteries, Murder and Legends', 55.0, 'Vaticano Prati, Rome', '20', '1458', '4.82', '500');

----------------Rentals Table--------------------------------------------------------------------------------------------

CREATE TABLE Rentals (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    duration TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    guest TEXT,
    bedroom TEXT,
    bed Text,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Rentals (ID, tag, title, price, location, ratings, numberOfReviews, guest, bedroom, bed, delayAnimation)
VALUES
(1, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.8', '3014', '2', '1', '1', '100'),
(2, '', 'Premium One Bedroom Luxury Living in the Heart of Mayfair', 72.0, 'Ciutat Vella, Barcelona', '4.8', '3014', '4', '2', '1', '200'),
(3, 'best seller', 'Style, Charm & Comfort in Camberwell', 72.0, 'Manhattan, New York', '4.5', '3014', '2', '1', '1', '300'),
(4, 'top rated', 'Marylebone - Oxford Street 1 bed apt with WiFi', 72.0, 'Vaticano Prati, Rome', '4.8', '3014', '3', '2', '1', '400'),
(5, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.9', '3014', '2', '1', '1', '100'),
(6, '', 'Premium One Bedroom Luxury Living in the Heart of Mayfair', 72.0, 'Ciutat Vella, Barcelona', '4.8', '3014', '4', '2', '1', '200'),
(7, 'best seller', 'Style, Charm & Comfort in Camberwell', 72.0, 'Manhattan, New York', '4.65', '3014', '2', '1', '1', '300'),
(8, 'top rated', 'Marylebone - Oxford Street 1 bed apt with WiFi', 72.0, 'Vaticano Prati, Rome', '4.8', '3014', '3', '2', '1', '400'),
(9, '', 'Luxury New Apartment With Private Garden', 72.0, 'Westminster Borough, London', '4.7', '3014', '2', '1', '1', '100');

-------------------------Cars Table----------------------------------------------
CREATE TABLE Cars (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    seat TEXT,
    type Text,
    luggage TEXT,
    transmission TEXT,
    speed TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO Cars (ID, tag, title, price, location, type, ratings, reviews, seat, luggage, transmission, speed, animation)
VALUES
(1, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(2, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(3, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(4, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(5, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(6, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(7, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(8, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(9, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(10, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(11, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100'),
(12, '', 'Jaguar F-Pace', 99.0, 'Heathrow Airport', 'SUV', '4.6', '2345', '3', '1', 'Manual', 'Unlimited', '200'),
(13, 'best seller', 'Volvo XC90', 87.0, 'Heathrow Airport', 'SUV', '4.9', '4321', '5', '1', 'Automatic', 'Unlimited', '300'),
(14, '', 'BMW 5 Series', 78.0, 'Heathrow Airport', 'LUXURY', '5.0', '2432', '3', '1', 'Automatic', 'Unlimited', '400'),
(15, '', 'Mercedes-Benz E-Class', 72.0, 'Heathrow Airport', 'LUXURY', '4.8', '3014', '4', '1', 'Automatic', 'Unlimited', '100');


------------------Golf Table---------------------------------
CREATE TABLE Golfs (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    price REAL,
    location TEXT,
    reviews TEXT,
    ratings TEXT,
    animation TEXT,
    holes TEXT,
    duration TEXT,
    name TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Golfs (ID, tag, title, price, location, reviews, ratings, animation, holes, duration, name)
VALUES
(1, 'golf only', '7 Night Golf to the Western Mediterranean', 67.0, 'Southampton', '94', '4.82', '100', '9', '18/05/23', 'Norwegian Getaway'),
(2, '', '12 Night Golf to the British Isles & Western Europe', 99.0, 'Southampton', '74', '4.82', '200', '8', '18/02/23', 'Norwegian Getaway'),
(3, 'best seller', '11 Night Golf to the Eastern Mediterranean', 88.0, 'Southampton', '56', '4.82', '300', '7', '18/03/23', 'Norwegian Getaway'),
(4, 'top rated', '13 Night Golf to the Baltic Sea Western Europe', 55.0, 'Southampton', '54', '4.82', '400', '6', '18/04/23', 'Norwegian Getaway'),
(5, 'golf only', '7 Night Golf to the Western Mediterranean', 67.0, 'Southampton', '94', '4.82', '100', '9', '18/05/23', 'Norwegian Getaway'),
(6, '', '12 Night Golf to the British Isles & Western Europe', 99.0, 'Southampton', '74', '4.82', '200', '8', '18/02/23', 'Norwegian Getaway'),
(7, 'best seller', '11 Night Golf to the Eastern Mediterranean', 88.0, 'Southampton', '56', '4.82', '300', '7', '18/03/23', 'Norwegian Getaway'),
(8, 'top rated', '13 Night Golf to the Baltic Sea Western Europe', 55.0, 'Southampton', '54', '4.82', '400', '6', '18/04/23', 'Norwegian Getaway'),
(9, 'golf only', '7 Night Golf to the Western Mediterranean', 67.0, 'Southampton', '94', '4.82', '500', '9', '18/05/23', 'Norwegian Getaway');

-------------------------Flights Table----------------------------------------------

CREATE TABLE Flights (
    ID INTEGER PRIMARY KEY,
    price REAL,
    deals TEXT,
    animation TEXT,
    SelectId TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);  

INSERT INTO Flights (ID, price, deals, animation, SelectId)
VALUES
(1, 934.0, '16', '100', 'collapse_1'),
(2, 690.0, '12', '200', 'collapse_2'),
(3, 999.0, '17', '300', 'collapse_3'),
(4, 859.0, '15', '400', 'collapse_4'),
(5, 934.0, '16', '500', 'collapse_5'),
(6, 690.0, '12', '600', 'collapse_6');

   

CREATE TABLE Flight_List (
    ID INTEGER PRIMARY KEY,
    flight_id INTEGER,
    avatar TEXT,
    arrivalAirport TEXT,
    departureAirport TEXT,
    departureTime TEXT,
    arrivalTime TEXT,
    duration TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (flight_id) REFERENCES flights(id)
);

INSERT INTO Flight_List (ID, flight_id, avatar, arrivalAirport, departureAirport, departureTime, arrivalTime, duration)
VALUES
(1, 1, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '3h 05m- Nonstop'),
(2, 1, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '5h 05m- Nonstop'),
(3, 1, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '4h 05m- Nonstop'),
(4, 2, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '6h 05m- Nonstop'),
(5, 2, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '4h 05m- Nonstop'),
(6, 3, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '7h 05m- Nonstop'),
(7, 3, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '3h 05m- Nonstop'),
(8, 4, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '5h 05m- Nonstop'),
(9, 4, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '4h 05m- Nonstop'),
(10, 5, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '6h 05m- Nonstop'),
(11, 6, '/img/flightIcons/new/1.png', 'SAW', 'STN', '14:00', '22:00', '3h 05m- Nonstop'),
(12, 6, '/img/flightIcons/new/2.png', 'SAW', 'STN', '14:00', '22:00', '5h 05m- Nonstop');

-------------------------Merchant Table -----------------------------------


CREATE TABLE Merchants (
    id INTEGER PRIMARY KEY,
    businessName TEXT,
    userName TEXT,
    firstName TEXT,
    lastName TEXT,
    email TEXT,
    phoneNumber TEXT,
    birthday DATE,
    about TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Merchant_User table inserts
INSERT INTO Merchants (businessName, userName, firstName, lastName, email, phoneNumber, birthday, about)
VALUES 
    ('ABC Store', 'abc_user', 'John', 'Doe', 'john@example.com', '123-456-7890', '1990-05-15', 'About ABC Store...'),
    ('XYZ Boutique', 'xyz_user', 'Alice', 'Smith', 'alice@example.com', '987-654-3210', '1985-08-20', 'About XYZ Boutique...'),
    ('Fashion House', 'fashion_user', 'Emily', 'Johnson', 'emily@example.com', '111-222-3333', '1988-12-10', 'About Fashion House...'),
    ('Tech Haven', 'tech_user', 'Michael', 'Brown', 'michael@example.com', '444-555-6666', '1995-04-25', 'About Tech Haven...'),
    ('Green Grocers', 'green_user', 'Sarah', 'Lee', 'sarah@example.com', '777-888-9999', '1980-11-30', 'About Green Grocers...');


CREATE TABLE LocationInfos (
    id INTEGER PRIMARY KEY,
    Merchant_id INTEGER, 
    addressLine1 TEXT,
    addressLine2 TEXT,
    city TEXT,
    state TEXT,
    country TEXT,
    zipCode TEXT,
    FOREIGN KEY (Merchant_id) REFERENCES Merchants(id)
);

-- Location_Information table inserts
INSERT INTO LocationInfos(Merchant_id, addressLine1, addressLine2, city, state, country, zipCode)
VALUES 
    (1, '123 Main St', 'Suite 101', 'Anytown', 'CA', 'USA', '12345'),
    (2, '456 Elm St', NULL, 'Sometown', 'NY', 'USA', '54321'),
    (3, '789 Oak St', 'Unit 5', 'Otherplace', 'TX', 'USA', '67890'),
    (4, '101 Pine St', 'Apt 3B', 'TechCity', 'WA', 'USA', '98765'),
    (5, '222 Maple St', NULL, 'Greenville', 'FL', 'USA', '13579');

CREATE TABLE ChangePasses (
    id INTEGER PRIMARY KEY,
    Merchant_id INTEGER, 
    currentPassword TEXT,
    newPassword TEXT,
    newPasswordAgain TEXT,
    FOREIGN KEY (Merchant_id) REFERENCES Merchants(id)
);

-- Change_Password table inserts
INSERT INTO ChangePasses (Merchant_id, currentPassword, newPassword, newPasswordAgain)
VALUES 
    (1, 'old_password', 'new_password', 'new_password'),
    (2, 'pass123', 'securePass456', 'securePass456'),
    (3, 'abc_123', 'newPass789', 'newPass789'),
    (4, 'initialPass', 'updatedPass', 'updatedPass'),
    (5, 'passWord987', 'secure123', 'secure123');


-- users Table
CREATE TABLE Users (
    ID INTEGER PRIMARY KEY,
    UserName TEXT UNIQUE,
    Password TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Customers Table
CREATE TABLE Customers (
    ID INTEGER PRIMARY KEY,
    First_Name TEXT,
    Last_Name TEXT,
    Email TEXT,
    Phone TEXT,
    Address TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Products Table
CREATE TABLE Products (
    ID INTEGER PRIMARY KEY,
    Product_Name TEXT,
    Description TEXT,
    Price REAL,
    Category TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Orders Table
CREATE TABLE Orders (
    ID INTEGER PRIMARY KEY,
    Customer_ID INTEGER,
    Order_Date TEXT,
    Total_Amount REAL,
    Status TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID)
);

-- OrderDetails Table
CREATE TABLE Order_Details (
    ID INTEGER PRIMARY KEY,
    Order_ID INTEGER,
    Product_ID INTEGER,
    Quantity INTEGER,
    Subtotal REAL,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Order_ID) REFERENCES Orders(ID),
    FOREIGN KEY (Product_ID) REFERENCES Products(ID)
);



-- Hotels Table
CREATE TABLE Travel_Bookings (
    ID INTEGER PRIMARY KEY,
    Customer_ID INTEGER,
    Package_ID INTEGER,
    Travel_Date TEXT,
    Status TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID),
    FOREIGN KEY (Package_ID) REFERENCES Tour_Packages(ID)
);

-- Payments Table
CREATE TABLE Payments (
    ID INTEGER PRIMARY KEY,
    Order_ID INTEGER,
    Amount REAL,
    Payment_Date TEXT,
    Payment_Method TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Order_ID) REFERENCES Orders(ID)
);

-- Reviews and Ratings Table
CREATE TABLE Review_Ratings (
    ID INTEGER PRIMARY KEY,
    Product_ID INTEGER,
    Customer_ID INTEGER,
    Rating INTEGER,
    Review_Text TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Product_ID) REFERENCES Products(ID),
    FOREIGN KEY (Customer_ID) REFERENCES Customers(ID)
);

---------------- The new Schema----------------------------------

CREATE TABLE Categories (
    ID                 INTEGER PRIMARY KEY,
    Name               TEXT    NOT NULL,
    Description        TEXT,
    Parent_Category_ID INTEGER,
    FOREIGN KEY (
        Parent_Category_ID
    )
    REFERENCES Categories (ID) 
);

INSERT INTO Categories (ID, Name, Description) VALUES (1, 'Flights', 'Flights Category');
INSERT INTO Categories (ID, Name, Description) VALUES (2, 'Tours', 'Tours Category');
INSERT INTO Categories (ID, Name, Description) VALUES (3, 'Hotels', 'Hotels Category');
INSERT INTO Categories (ID, Name, Description) VALUES (4, 'Rentals', 'Rentals Category');
INSERT INTO Categories (ID, Name, Description) VALUES (5, 'Cars', 'Cars Category');
INSERT INTO Categories (ID, Name, Description) VALUES (6, 'Golfs', 'Golfs Category');
INSERT INTO Categories (ID, Name, Description) VALUES (7, 'Flights', 'Flights Category');
INSERT INTO Categories (ID, Name, Description) VALUES (8, 'Activities', 'Activities Category');



------------------Testimonials Table -------------------------------------------
CREATE TABLE Testimonials (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER,
    meta TEXT,
    avatar TEXT,
    name TEXT,
    designation TEXT,
    text TEXT,
    delayAnimation INTEGER,
     FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

INSERT INTO Testimonials (id, hotel_id, meta, avatar, name, designation, text, delayAnimation)
VALUES 
(1, 4, Hotel Equatorial Melaka, /img/avatars/new/testimonials/new/1.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 100),
(2, 5,  Hotel Equatorial Melaka, /img/avatars/new/testimonials/new/2.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 200),
(3, 4,  Hotel Equatorial Melaka, /img/avatars/new/testimonials/new/3.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 300),
(4, 5, Hotel Equatorial Melaka, /img/avatars/new/testimonials/new/1.png, Annette Black, UX / UI Designer, The place is in a great location in Gumbet. The area is safe and beautiful. The apartment was comfortable and the host was kind and responsive to our requests. Really a nice place., 400);


---------------------Blogs table --------------------------
CREATE TABLE Blogs (
    id INTEGER PRIMARY KEY,
    img TEXT,
    title TEXT,
    date TEXT,
    delayAnimation INTEGER,
    details TEXT,
    tag TEXT,
    tags TEXT
);


INSERT INTO Blogs (id, img, title, date, delayAnimation, details, tag, tags)
VALUES 
(1, /img/blog/1.png, 10 European ski destinations you should visit this winter, Jan 06, 2023, 100, For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset., art, adventure_travel, food_drink),
(2, /img/blog/2.png, Booking travel during Corona: good advice in an uncertain time, April 06, 2022, 200, For decades, travelers have been escaping to the Catskills — a mountainous region in upstate New York — whenever they’ve needed a reset., beaches, beaches, family_holidays);
