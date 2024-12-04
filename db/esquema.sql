create database task_manager;
use task_manager;

create table Users(
    id varchar(255) not null,
    first_name varchar(45) not null,
    last_name varchar(45) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    primary key(id)
)

create table Tasks(
    id varchar(255) not null,  
    name varchar(45) not null,
    description varchar(255) not null,
    due_date date not null,
    create_at date not null,
    update_at date,
    status tinyint(1),
    user_id varchar(255) not null,
    primary key(id),
    foreign key (user_id) references Users(id)
)

create table Proyects(
    
)