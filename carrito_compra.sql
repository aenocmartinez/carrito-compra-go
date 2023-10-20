DROP DATABASE IF EXISTS carritodb;

create database carritodb;

use carritodb;


create table items(
	id bigint auto_increment,
	name varchar(60) not null,
	amount int default 0,
	primary key(id)
);
