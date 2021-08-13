-- migrate:up
create table users (
  id INTEGER PRIMARY KEY,
  username varchar(255) UNIQUE,
  hashed_password varchar(255)
);

-- migrate:down
drop table users;