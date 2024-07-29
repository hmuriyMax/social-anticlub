create role social with password 'qwerty' login;

create role grafanareader with password 'grafana2001' login;
grant usage on schema public to grafanareader;
grant select on all tables in schema public to grafanareader;

create role mshchemilkin with password 'heho2001' login;
grant usage on schema public to social, mshchemilkin;
grant all privileges on database social to mshchemilkin, social;

create role repluser with password 'replsocial' login replication;
