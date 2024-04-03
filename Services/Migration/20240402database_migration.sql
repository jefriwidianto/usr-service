CREATE DATABASE IF NOT EXISTS diksha_serv_user;

USE diksha_serv_user;

CREATE TABLE IF NOT EXISTS t_user(
    id varchar(36) UNIQUE PRIMARY KEY,
    username VARCHAR(50),
    email VARCHAR(50),
    phone_number VARCHAR(50),
    password VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME
);

INSERT INTO t_user (id, username, email, phone_number, password) VALUES ('9a78a891-2b28-4cd8-b5c0-a33b45473204', 'test', 'test@gmail.com', '+628123456789', '827ccb0eea8a706c4c34a16891f84e7b') ON DUPLICATE KEY UPDATE id = '9a78a891-2b28-4cd8-b5c0-a33b45473204'