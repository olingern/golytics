-- migrate:up
create table logs (
  id INTEGER PRIMARY KEY,
  ip varchar(255),
  ua varchar(255),
  lang varchar(255),
  date varchar(255),
  processed INTEGER
);

-- migrate:down
drop table logs;