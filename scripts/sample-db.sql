--
-- File generated with SQLiteStudio v3.4.4 on Fri Feb 23 14:11:45 2024
--
-- Text encoding used: System
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Activities
CREATE TABLE IF NOT EXISTS Activities (
    ID                     INTEGER  PRIMARY KEY,
    tag                    TEXT,
    title                  TEXT,
    number_of_reviews      TEXT,
    reviews_comment        TEXT,
    location               TEXT,
    latitude               TEXT,
    longitude              TEXT,
    minimum_duration       TEXT,
    group_size             TEXT,
    overview               TEXT,
    cancellation_policy    TEXT,
    whats_included         TEXT,
    highlights             TEXT,
    additional_information TEXT,
    important_information  TEXT,
    price                  TEXT,
    activity_type          TEXT,
    animation              TEXT,
    Created_At             DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At             DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Activities (
                           ID,
                           tag,
                           title,
                           number_of_reviews,
                           reviews_comment,
                           location,
                           latitude,
                           longitude,
                           minimum_duration,
                           group_size,
                           overview,
                           cancellation_policy,
                           whats_included,
                           highlights,
                           additional_information,
                           important_information,
                           price,
                           activity_type,
                           animation,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           2,
                           'LIKELY TO SELL OUT*',
                           'Marrakech to Fez via Merzouga Desert 3-Days Morocco Sahara Tour',
                           '1458',
                           'Great experience!',
                           '?????, Morocco Marrakech',
                           '41.9055',
                           '12.4655',
                           '11',
                           '52',
                           'Travel from Marrakech to Fes with a stop in the desert along the way with this small-group transfer-tour hybrid. You''ll visit the Kasbah at Ait Ben Haddou, spend time in the Gorge of Todra, and take an overnight camelback safari to a desert camp in the Sahara, all in three days',
                           'For a full refund, cancel at least 24 hours in advance of the start date of the experience.',
                           '',
                           '',
                           'Explore the beautiful mountains.',
                           'Explore the beautiful mountains.',
                           '75.99',
                           'this is a activity type: need to be dynamic',
                           '100',
                           '2024-02-20 09:37:20.1337363+00:00',
                           '2024-02-20 09:37:20.1337363+00:00'
                       );

INSERT INTO Activities (
                           ID,
                           tag,
                           title,
                           number_of_reviews,
                           reviews_comment,
                           location,
                           latitude,
                           longitude,
                           minimum_duration,
                           group_size,
                           overview,
                           cancellation_policy,
                           whats_included,
                           highlights,
                           additional_information,
                           important_information,
                           price,
                           activity_type,
                           animation,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           3,
                           'TOP RATED',
                           'Guided Pottery and Zellige Workshops in Fes Morocco',
                           '10',
                           'Exellent',
                           'City Center FEZ',
                           '45',
                           '22',
                           '8 to 10',
                           '10',
                           'People from various backgrounds and with different levels of expertise come together to share their passion for these crafts. The workshops often encourage collaboration, knowledge exchange, and creative exploration, creating a vibrant and supportive atmosphere for artistic growth.

At the end of both workshops, participants will not only take their newfound knowledge and skills but also the tangible results of their hard work. In the pottery workshop, you will proudly carry home your handmade pottery pieces, showcasing your personal touch and creativity. Whether it�s a functional ceramic item or a decorative sculpture, it will serve as a reminder of the fulfilling journey you embarked upon during the workshop. Similarly, in the Zellige workshop, participants will leave with their own mosaic creation, capturing the essence of this ancient art form. This tangible artifact will serve as a testament to your dedication and craftsmanship.',
                           'You cannot Cancel',
                           'Private transportation, Tour guide at the Workshop, Instructors teaching you how to make traditional crafts of pottery and Zellige, Pick up for free to our workshop, Drinks',
                           'Wheelchair accessible, Infants and small children can ride in a pram or stroller, Public transportation options are available nearby, Infants are required to sit on an adult�s lap, Suitable for all physical fitness levels',
                           'Private transportation, Tour guide at the Workshop, Instructors teaching you how to make traditional crafts of pottery and Zellige, Pick up for free to our workshop, Drinks',
                           'Private transportation, Tour guide at the Workshop, Instructors teaching you how to make traditional crafts of pottery and Zellige, Pick up for free to our workshop, Drinks',
                           '66',
                           'this is a activity type: need to be dynamic',
                           '100',
                           '2024-02-20 15:55:25.4207629+00:00',
                           '2024-02-20 15:55:25.4207629+00:00'
                       );

INSERT INTO Activities (
                           ID,
                           tag,
                           title,
                           number_of_reviews,
                           reviews_comment,
                           location,
                           latitude,
                           longitude,
                           minimum_duration,
                           group_size,
                           overview,
                           cancellation_policy,
                           whats_included,
                           highlights,
                           additional_information,
                           important_information,
                           price,
                           activity_type,
                           animation,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           5,
                           'LIKELY TO SELL OUT*',
                           'Marrakech to Fez via Merzouga Desert 3-Days Morocco Sahara Tour',
                           '1458',
                           'Great experience!',
                           '?????, Morocco Marrakech',
                           '41.9055',
                           '12.4655',
                           '11',
                           '52',
                           'Travel from Marrakech to Fes with a stop in the desert along the way with this small-group transfer-tour hybrid. You''ll visit the Kasbah at Ait Ben Haddou, spend time in the Gorge of Todra, and take an overnight camelback safari to a desert camp in the Sahara, all in three days',
                           'For a full refund, cancel at least 24 hours in advance of the start date of the experience.',
                           '',
                           '',
                           'Explore the beautiful mountains.',
                           'Explore the beautiful mountains.',
                           '75.99',
                           'this is a activity type: need to be dynamic',
                           '100',
                           '2024-02-19 22:09:10.9252796+00:00',
                           '2024-02-19 22:09:10.9252796+00:00'
                       );


-- Table: Admins
CREATE TABLE IF NOT EXISTS Admins (
    id    INTEGER PRIMARY KEY,
    name  TEXT    NOT NULL,
    email TEXT    NOT NULL
                  UNIQUE
);


-- Table: Categories
CREATE TABLE IF NOT EXISTS Categories (
    ID          INTEGER  PRIMARY KEY,
    Name        TEXT     NOT NULL,
    Description TEXT,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           1,
                           'Flights',
                           'Flights Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           2,
                           'Tours',
                           'Tours Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           3,
                           'Hotels',
                           'Hotels Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           4,
                           'Rentals',
                           'Rentals Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           5,
                           'Cars',
                           'Cars Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           6,
                           'Golfs',
                           'Golfs Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           7,
                           'Flights',
                           'Flights Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );

INSERT INTO Categories (
                           ID,
                           Name,
                           Description,
                           Created_At,
                           Updated_At
                       )
                       VALUES (
                           8,
                           'Activities',
                           'Activities Category',
                           '2024-02-01 12:32:14',
                           '2024-02-01 12:32:14'
                       );


-- Table: Configurations
CREATE TABLE IF NOT EXISTS Configurations (
    ID    INTEGER PRIMARY KEY AUTOINCREMENT,
    key   TEXT    UNIQUE
                  NOT NULL,
    value TEXT    NOT NULL
);

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               1,
                               'JWT-API-KEY',
                               'EL5ELKUpq782HdKzXTPP7uwPKERw6ByRStdWMOr1CHY='
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               2,
                               'REDIS_URI',
                               'redis:6379'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               3,
                               'DB_NAME',
                               'niya-voyage.db'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               4,
                               'HOST',
                               'localhost'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               5,
                               'PORT',
                               ':8080'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               6,
                               'GOOGLE-STORAGE',
                               'https://storage.googleapis.com/'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               7,
                               'GOOGLE-BUCKET-NAME',
                               'niya-voyage-app-images'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               8,
                               'GOOGLE-PROJECT-ID',
                               'niya-voyage'
                           );

INSERT INTO Configurations (
                               ID,
                               key,
                               value
                           )
                           VALUES (
                               9,
                               'TOURS-FOLDER',
                               'tours/'
                           );


-- Table: Faqs
CREATE TABLE IF NOT EXISTS Faqs (
    id             INTEGER  PRIMARY KEY,
    category_id    INTEGER,
    referrer_id    INTEGER,
    title          TEXT,
    content        TEXT,
    collapseTarget TEXT,
    Created_At     DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At     DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        category_id
    )
    REFERENCES Categories (id) 
);


-- Table: GalleryImages
CREATE TABLE IF NOT EXISTS GalleryImages (
    id          INTEGER  PRIMARY KEY,
    category_id INTEGER,
    referrer_id INTEGER,-- This is the product id
    img         TEXT,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        category_id
    )
    REFERENCES Categories (id) 
);


-- Table: Itineraries
CREATE TABLE IF NOT EXISTS Itineraries (
    id          INTEGER  PRIMARY KEY,
    category_id INTEGER,
    referrer_id INTEGER,
    img         TEXT,
    title       TEXT,
    content     TEXT,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        category_id
    )
    REFERENCES tours (id) 
);


-- Table: Products
CREATE TABLE IF NOT EXISTS Products (
    id          INTEGER  PRIMARY KEY AUTOINCREMENT,
    name        TEXT,
    description TEXT,
    price       REAL,
    category_id INTEGER,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        category_id
    )
    REFERENCES Categories (id) 
);

INSERT INTO Products (
                         id,
                         name,
                         description,
                         price,
                         category_id,
                         Created_At,
                         Updated_At
                     )
                     VALUES (
                         1,
                         'Product1',
                         'Description1',
                         500.0,
                         3,
                         '2024-02-07 15:06:37',
                         '2024-02-07 15:06:37'
                     );

INSERT INTO Products (
                         id,
                         name,
                         description,
                         price,
                         category_id,
                         Created_At,
                         Updated_At
                     )
                     VALUES (
                         2,
                         'Product2',
                         'Description2',
                         750.0,
                         3,
                         '2024-02-07 15:06:37',
                         '2024-02-07 15:06:37'
                     );

INSERT INTO Products (
                         id,
                         name,
                         description,
                         price,
                         category_id,
                         Created_At,
                         Updated_At
                     )
                     VALUES (
                         3,
                         'Product3',
                         'Description3',
                         1000.0,
                         3,
                         '2024-02-07 15:06:37',
                         '2024-02-07 15:06:37'
                     );

INSERT INTO Products (
                         id,
                         name,
                         description,
                         price,
                         category_id,
                         Created_At,
                         Updated_At
                     )
                     VALUES (
                         4,
                         'Product4',
                         'Description4',
                         1500.0,
                         3,
                         '2024-02-07 15:06:37',
                         '2024-02-07 15:06:37'
                     );

INSERT INTO Products (
                         id,
                         name,
                         description,
                         price,
                         category_id,
                         Created_At,
                         Updated_At
                     )
                     VALUES (
                         5,
                         'Product5',
                         'Description5',
                         2000.0,
                         3,
                         '2024-02-07 15:06:37',
                         '2024-02-07 15:06:37'
                     );


-- Table: Reviews
CREATE TABLE IF NOT EXISTS Reviews (
    id         INTEGER  PRIMARY KEY AUTOINCREMENT,
    text       TEXT,
    rating     INTEGER,
    product_ID INTEGER,
    user_ID    INTEGER,
    username   TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        product_ID
    )
    REFERENCES products (id),
    FOREIGN KEY (
        user_ID
    )
    REFERENCES users (id) 
);


-- Table: SlideImages
CREATE TABLE IF NOT EXISTS SlideImages (
    id          INTEGER  PRIMARY KEY,
    category_id INTEGER,
    referrer_id INTEGER,
    img         TEXT,
    Created_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (
        category_id
    )
    REFERENCES Categories (id) 
);


-- Table: Tours
CREATE TABLE IF NOT EXISTS Tours (
    ID                     INTEGER  PRIMARY KEY,
    tag                    TEXT,
    title                  TEXT,
    number_of_reviews      TEXT,
    reviews_comment        TEXT,
    location               TEXT,
    latitude               TEXT,
    longitude              TEXT,
    minimum_duration       TEXT,
    group_size             TEXT,
    overview               TEXT,
    cancellation_policy    TEXT,
    whats_included         TEXT,
    highlights             TEXT,
    additional_information TEXT,
    important_information  TEXT,
    price                  TEXT,
    tour_type              TEXT,
    animation              TEXT,
    Created_At             DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At             DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Tours (
                      ID,
                      tag,
                      title,
                      number_of_reviews,
                      reviews_comment,
                      location,
                      latitude,
                      longitude,
                      minimum_duration,
                      group_size,
                      overview,
                      cancellation_policy,
                      whats_included,
                      highlights,
                      additional_information,
                      important_information,
                      price,
                      tour_type,
                      animation,
                      Created_At,
                      Updated_At
                  )
                  VALUES (
                      0,
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '',
                      '2024-02-18 09:54:45.1374579+00:00',
                      '2024-02-21 16:29:52.7902547+00:00'
                  );


-- Table: Users
CREATE TABLE IF NOT EXISTS Users (
    id         INTEGER  PRIMARY KEY AUTOINCREMENT,
    firstname  TEXT,
    lastname   TEXT,
    email      TEXT,
    password   TEXT     DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Created_At DATETIME DEFAULT (CURRENT_TIMESTAMP),
    token    string  
);

INSERT INTO Users (
                      id,
                      firstname,
                      lastname,
                      email,
                      password,
                      Updated_At,
                      Created_At,
                      token
                  )
                  VALUES (
                      42,
                      'mustapha',
                      'manjoura',
                      'mustapha.manjoura@gmail.com',
                      '$2a$14$M56c11LsCvQNvqtU9f3ZCugna0JhQLr5FNxtv8AQxC3JwIC5sSJfe',
                      '2024-02-14 15:42:41.1779984+00:00',
                      'k1U6pO+9qZteWy+yE52Z56qSBqmJ1orl27r/28AfkIA=',
                      1
                  );


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
