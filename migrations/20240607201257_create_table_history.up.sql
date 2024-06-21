CREATE TABLE IF NOT EXISTS history (
  id serial PRIMARY KEY,
  origin_text VARCHAR(255),
  translation VARCHAR(255),
  language_destination VARCHAR(4),
  language_origin_detected VARCHAR(4)
)
