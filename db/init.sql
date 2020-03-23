create database onlinedoctor;
\c onlinedoctor;
create table if not exists users (
    id bigserial primary key,
    fname varchar(50) not null,
    lname varchar(50) not null,
    email varchar(100) not null unique,
    mobile varchar(20) not null,
    avatar varchar(255),
    password varchar(255) not null,
    userType varchar(10) not null,
    active char(1) default '0',
    verified char(1) default '0',
    creatorId bigint not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp

);

create table if not exists appointments (
    id bigserial primary key,
    doctor bigint not null,
    patient bigint not null,
    speciality bigint not null,
    sub_speciality bigint not null,
    details varchar(255),
    duration real not null default 0.0,
    cost real not null default 0.0,
    start_at timestamp not null,
    end_at timestamp not null,
    status varchar(255) not null default 'pending',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp   
);