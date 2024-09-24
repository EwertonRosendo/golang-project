CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(150) not null,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE books(
    id int auto_increment primary key,
    title varchar(200) not null unique,
    subtitle varchar(200),
    description varchar(500),
    author varchar(200) not null ,
    publisher varchar(100),
    published_at varchar(10),
    cover varchar(200) not null unique,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;    

CREATE TABLE reviews (
    id int auto_increment primary key,
    book_id int NOT NULL,
    user_id int NOT NULL,
    status int,
    rating float,
    review varchar(400),
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    CreatedAt timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE comments (
    id int auto_increment primary key,
    review_id int NOT NULL,
    user_id int NOT NULL,
    comment varchar(400),
    FOREIGN KEY (review_id) REFERENCES reviews(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    CreatedAt timestamp default current_timestamp()
)ENGINE=INNODB;