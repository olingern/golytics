CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
CREATE TABLE logs (
  id INTEGER PRIMARY KEY,
  ip varchar(255),
  ua varchar(255),
  lang varchar(255),
  date varchar(255),
  processed INTEGER
);
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  username varchar(255) UNIQUE,
  hashed_password varchar(255)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20210404033241'),
  ('20210812195543');
