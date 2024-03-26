create table if not exists user_auth
(
    id        bigint      not null,
    login     uuid        not null primary key,
    pass_hash varchar(32) not null
);

comment on table user_auth is 'Users auth info';

comment on column user_auth.login is 'User login';

comment on column user_auth.pass_hash is 'User password hash';
