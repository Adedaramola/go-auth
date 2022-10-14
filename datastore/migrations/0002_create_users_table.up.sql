create table users (
    id serial primary key,
    fullname varchar(100) not null,
    email varchar(255) unique not null,
    password text not null,
    created_at timestamp not null
);