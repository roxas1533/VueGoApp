CREATE DATABASE sample_db;
use sample_db;

CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  mail_adress varchar(255) not null,
  password varchar(1024) not null, 
  name varchar(255) not null,
  created_time datetime not null default current_timestamp,
  updated_time datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);