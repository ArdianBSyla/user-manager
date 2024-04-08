-- create user table with company_id as foreign key in mysql database with unique email address
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    company_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES company(id)
);

-- some insert statements to populate user table
INSERT INTO user (first_name, last_name, email, company_id) VALUES ('John', 'Doe', 'john@email.com', 1);
INSERT INTO user (first_name, last_name, email, company_id) VALUES ('Ardian', 'Syla', 'ardianbsyla@gmail.com', 1);
