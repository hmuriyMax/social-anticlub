create table if not exists login_info
(
    token   text    not null primary key,
    user_id integer not null
);

comment on column login_info.token is 'Login token';

comment on column login_info.user_id is 'Login user';