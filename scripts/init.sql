create database test;

use test;

create table File (
    ID int not null AUTO_INCREMENT primary key,
    Name varchar(25),
    Size bigint,
    CreatedTime timestamp,
    LastUpdatedTime timestamp,
    IsDir boolean
);