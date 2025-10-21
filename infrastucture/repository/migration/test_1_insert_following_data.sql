INSERT INTO following (user_id, follow_user_id)
SELECT
  gs AS user_id,
  1  AS follow_user_id
FROM generate_series(100, 999999) AS gs;




-- Test 1: Fan-out (Create Post)
-- 1000 users đồng thời đăng bài
-- Mỗi users có 500 followers
-- đo tổng thời gian xử lý từ khi gọi CreatePost() đến khi CachePost() cho toàn bộ followers xong.
DO $$
DECLARE
    f_id INT;
BEGIN
    FOR f_id IN 1..1000 LOOP
        INSERT INTO following (user_id, follow_user_id)
        SELECT u_id, f_id
        FROM (
            SELECT DISTINCT (1 + floor(random() * 10000))::INT AS u_id
            FROM generate_series(1, 100000)
        ) AS followers
        WHERE f_id <> u_id
        LIMIT 500;
    END LOOP;
END $$;
