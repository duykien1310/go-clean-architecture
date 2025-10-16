INSERT INTO "user" (user_name, password, first_name, last_name)
SELECT
  'user_' || gs::text AS user_name,
  md5(random()::text) AS password,
  initcap(md5(random()::text)::text) AS first_name,
  initcap(md5(random()::text)::text) AS last_name
FROM generate_series(1, 1000000) AS gs;
