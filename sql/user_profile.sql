create table `user_profile` (
    id int primary key auto_increment,
    user_id int not null unique,
    firstname text not null,
    lastname text not null,
    phone_number varchar(12),
    created_at datetime not null,
    created_by varchar(255) not null,
    updated_at datetime not null,
    updated_by varchar(255) not null
)