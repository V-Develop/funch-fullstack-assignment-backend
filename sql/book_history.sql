create table `book_history` (
    id int primary key auto_increment,
    uuid varchar(255) not null unique,
    -- room_id int not null,
    booker_id int not null,
    email varchar(255) not null,
    firstname text not null,
    lastname text not null,
    phone_number varchar(12) not null,
    checkin_at datetime not null,
    checkout_at datetime not null,
    -- total_price double not null,
    created_at datetime not null,
    created_by varchar(255) not null,
    updated_at datetime not null,
    updated_by varchar(255) not null
)

-- commend is for new feature (can book multiple room)