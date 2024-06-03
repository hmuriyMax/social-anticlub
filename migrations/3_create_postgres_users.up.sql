create role grafanareader with password 'grafana2001' login;
grant usage on schema public to grafanareader;
grant select on all tables in schema public to grafanareader;

create role mshchemilkin with password 'heho2001' login;
grant all on database social to mshchemilkin;

create role repluser with password 'replsocial' login replication;
