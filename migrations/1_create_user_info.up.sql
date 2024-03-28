create table if not exists user_info
(
    user_uuid   uuid not null default gen_random_uuid() primary key,
    nickname    varchar(15) unique not null,
    first_name  text not null,
    second_name text,
    birthday    date not null,
    gender      smallint,
    hometown    text,
    about       text
);

comment on table user_info is 'User table';

