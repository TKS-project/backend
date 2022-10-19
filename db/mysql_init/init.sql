CREATE DATABASE IF NOT EXISTS test;
use test;

CREATE TABLE users (
    id INT(255) auto_increment NOT NULL,
    mail VARCHAR(255) UNIQUE,
    password VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE prefectures (
    id INT(3) auto_increment NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE cities (
    id INT(255) auto_increment NOT NULL,
    prefecture_id INT(3) NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE detailed_cities (
    id INT(255) auto_increment NOT NULL,
    city_id INT(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




CREATE TABLE budgets (
    id INT(255) NOT NULL,
    transportations INT(8) NOT NULL,
    accommodation INT(255) NOT NULL,
    sightseeing INT(255) NOT NULL,
    meal INT(255) NOT NULL
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE visits (
    id INT(255) NOT NULL,
    visit_day DATE,
    address VARCHAR,
    charge INT(255),
    info VARCHAR(255),
    name VARCHAR(255),
    visit_type varchar(255),
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE transportations (
    id INT(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    departure VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    charge INT(255) NOT NULL,
    order INT(8) NOT NULL,
    is_way_there Boolean,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

