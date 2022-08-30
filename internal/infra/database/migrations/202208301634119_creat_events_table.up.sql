CREATE TABLE IF NOT EXISTS events(
id                 serial
constraint events_pk
primary key,
"Title"            varchar,
"ShortDescription" varchar,
"Description"      varchar,
"Longitude"        double precision,
"Latitude"         double precision,
"Images"           varchar,
"Preview"          varchar,
"Date"             date,
"isEnded"          boolean,
"deletedDate"      timestamp
    );
create unique index events_id_uindex
    on events (id);