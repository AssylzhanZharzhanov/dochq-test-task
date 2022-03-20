CREATE TABLE IF NOT EXISTS  answers (
    id serial primary key not null unique,
    key varchar,
    value varchar
);

CREATE TABLE IF NOT EXISTS  events (
    id serial primary key not null unique,
    event varchar,
    data jsonb
);