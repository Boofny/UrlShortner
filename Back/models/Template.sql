DROP TABLE IF EXISTS urls;

CREATE TABLE urls (
  Id        SERIAL PRIMARY KEY,
  short_url TEXT NOT NULL,
  long_url  TEXT NOT NULL
);

