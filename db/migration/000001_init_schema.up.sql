CREATE TABLE users_table
( 
    id SERIAL NOT NULL UNIQUE, 
    name VARCHAR(255) not null,
    surname VARCHAR(255) not null,
    patronymic VARCHAR(255)
);