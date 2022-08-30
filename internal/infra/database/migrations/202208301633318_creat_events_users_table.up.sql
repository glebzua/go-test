CREATE TABLE IF NOT EXISTS events_users(
 id bigserial PRIMARY KEY,
 name          varchar(40)                                            not null,
    email         varchar(40)                                            not null,
    passhash      char(60)                                               not null,
    created_date  timestamp default (now())::timestamp without time zone not null,
    "deletedDate" timestamp,
    role          integer
    );