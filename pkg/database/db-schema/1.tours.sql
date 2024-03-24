CREATE TABLE Tours (
    ID INTEGER PRIMARY KEY,
    tag TEXT,
    title TEXT,
    location TEXT,
    latitude TEXT,
    longitude TEXT,
    minimum_duration TEXT,
    group_size TEXT,
    number_of_reviews TEXT,
    reviews_comment TEXT,
    overview TEXT,
    detailed_information TEXT,
    important_information TEXT,
    price TEXT, 
    tour_type TEXT,
    animation TEXT,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Tours (tag, title, price, location, latitude, longitude, minimum_duration, group_size, number_of_reviews, reviews_comment, overview, detailed_information, important_information, tour_type, animation)
VALUES
('LIKELY TO SELL OUT*', 'Mountain Hiking', '75.99', 'Mountain Range', '34.0522', '-118.2437', '4 hours', '10', '400', 'Great experience!', 'Unless you hire a car, visiting Stonehenge, Bath, and Windsor Castle in one day is next to impossible. Designed specifically for travelers with limited time in London, this tour allows you to check off a range of southern England‘s historical attractions in just one day by eliminating the hassle of traveling between each one independently. Travel by comfortable coach and witness your guide bring each UNESCO World Heritage Site to life with commentary. Plus, all admission tickets are included in the tour price.', 'Explore the beautiful mountains.', 'Explore the beautiful mountains.' , 'Outdoor', '100'),
('best seller', 'City Tour', '49.99', 'City Center', '40.7128', '-74.0060', '3 hours', '15', '400','Informative and fun', 'Unless you hire a car, visiting Stonehenge, Bath, and Windsor Castle in one day is next to impossible. Designed specifically for travelers with limited time in London, this tour allows you to check off a range of southern England‘s historical attractions in just one day by eliminating the hassle of traveling between each one independently. Travel by comfortable coach and witness your guide bring each UNESCO World Heritage Site to life with commentary. Plus, all admission tickets are included in the tour price.', 'Discover the city landmarks.', 'Explore the beautiful mountains.' , 'Guided', '100'),
('top rated"', 'Forest Exploration', '89.99', 'Deep Forest', '45.4215', '-75.6993', '5 hours', '8', '400','Breathtaking views', 'Unless you hire a car, visiting Stonehenge, Bath, and Windsor Castle in one day is next to impossible. Designed specifically for travelers with limited time in London, this tour allows you to check off a range of southern England‘s historical attractions in just one day by eliminating the hassle of traveling between each one independently. Travel by comfortable coach and witness your guide bring each UNESCO World Heritage Site to life with commentary. Plus, all admission tickets are included in the tour price.', 'Immerse yourself in nature.', 'Explore the beautiful mountains.' , 'Adventure', '100');

