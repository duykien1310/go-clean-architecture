INSERT INTO following (user_id, follow_user_id)
SELECT
  gs AS user_id,
  1  AS follow_user_id
FROM generate_series(100, 999999) AS gs;
