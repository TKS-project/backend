CREATE DATABASE IF NOT EXISTS test;
use test;

-- CREATE TABLE user (
--     user_id INT(8) NOT NULL,
--     user_name VARCHAR(200) NOT NULL,
--     password varchar(270) NOT  NULL,
--     mail varchar(30) unique
--     ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ALTER TABLE user
--   ADD PRIMARY KEY (`user_id`);


CREATE TABLE users (
    id INT(255) auto_increment NOT NULL,
    mail VARCHAR(255) UNIQUE,
    password VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id),
)

CREATE TABLE travels (
    id INT(255) auto_increment NOT NULL,
    user_id INT(255) NOT NULL,
    destination_id INT(255) NOT NULL,
    departure_id INT(255) NOT NULL,
    max_person INT(8) NOT NULL,
    max_day INT(8) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id),
)

