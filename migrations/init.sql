CREATE TABLE links(
id                SERIAL primary key,
link              VARCHAR UNIQUE NOT NULL,
short_link        VARCHAR NOT NULL
);
