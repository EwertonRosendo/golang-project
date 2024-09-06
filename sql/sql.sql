CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS books;

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