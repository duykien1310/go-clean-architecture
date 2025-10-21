INSERT INTO post (title, content, user_id)
SELECT 
    CONCAT('title ', p.i, ' - user ', u.id) AS title,
    CONCAT('content ', p.i, ' - user ', u.id) AS content,
    u.id AS user_id
FROM (
    SELECT id FROM "user" WHERE id BETWEEN 1 AND 1000
) AS u
CROSS JOIN generate_series(1, 200) AS p(i);


====
-- mỗi user follow ngẫu nhiên 300 người khác
INSERT INTO following (user_id, follow_user_id)
SELECT DISTINCT user_id, follow_user_id
FROM (
    SELECT 
        u.id AS user_id,
        (trunc(random() * 1000) + 1)::int AS follow_user_id
    FROM generate_series(1, 1000) AS u(id)
    CROSS JOIN generate_series(1, 300) g
    WHERE (trunc(random() * 1000) + 1)::int <> u.id
) t;
