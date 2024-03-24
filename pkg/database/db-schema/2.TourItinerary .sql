-- Create table
CREATE TABLE TourItineraries (
    id INTEGER PRIMARY KEY,
    targetCollapse TEXT,
    itemNo TEXT,
    title TEXT,
    img TEXT,
    content TEXT,
    classShowHide TEXT,
    tour_id INTEGER,
    Created_At DATETIME DEFAULT CURRENT_TIMESTAMP,
    Updated_At DATETIME DEFAULT CURRENT_TIMESTAMP,    
    FOREIGN KEY (tour_id) REFERENCES tours(id)
);

-- Insert data
INSERT INTO TourItineraries (id, targetCollapse, itemNo, title, img, content, classShowHide, Tour_id)
VALUES
(1, 'item_1', '1', 'Windsor Castle', '/img/tours/list.png', 'Our first stop is Windsor Castle, the ancestral home of the British Royal family for more than 900 years and the largest, continuously occupied castle in Europe.', '', 1),
(2, 'item_2', '2', 'St. George s Chapel', '/img/tours/list.png', 'Our first stop is Windsor Castle, the ancestral home of the British Royal family for more than 900 years and the largest, continuously occupied castle in Europe.', 'show', 1),
(3, 'item_3', '3', 'The Roman Baths', '/img/tours/list.png', 'Our first stop is Windsor Castle, the ancestral home of the British Royal family for more than 900 years and the largest, continuously occupied castle in Europe.', '', 1),
(4, 'item_4', '4', 'Stonehenge', '/img/tours/list.png', 'Our first stop is Windsor Castle, the ancestral home of the British Royal family for more than 900 years and the largest, continuously occupied castle in Europe.', '', 1);
