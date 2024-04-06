-- creating company table in mysql database
CREATE TABLE company (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    website VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- some insert statements to populate company table
INSERT INTO company (name, website) VALUES ('Google', 'https://www.google.com');
INSERT INTO company (name, website) VALUES ('Facebook', 'https://www.facebook.com');