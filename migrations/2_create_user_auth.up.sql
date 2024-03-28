create table if not exists user_auth
(
    user_uuid uuid not null primary key,
    pass_hash varchar(32) not null
);

comment on table user_auth is 'Users auth info';

comment on column user_auth.user_uuid is 'User uuid';

comment on column user_auth.pass_hash is 'User password hash';
