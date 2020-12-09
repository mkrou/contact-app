CREATE TABLE IF NOT EXISTS contacts
(
    id      serial PRIMARY KEY,
    name    varchar(30) NOT NULL,
    surname varchar(30) NOT NULL,
    phone   varchar(11) NOT NULL UNIQUE
);