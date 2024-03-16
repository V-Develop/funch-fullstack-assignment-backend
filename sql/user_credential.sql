create table `user_credential` (
    id int primary key auto_increment,
    uuid varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    last_login datetime,
    session_id varchar(255) null,
    refresh_token text,
    is_backlist boolean not null,
    otp varchar(255),
    otp_expire datetime,
    is_verify boolean not null,
    created_at datetime not null,
    created_by varchar(255) not null,
    updated_at datetime not null,
    updated_by varchar(255) not null
)