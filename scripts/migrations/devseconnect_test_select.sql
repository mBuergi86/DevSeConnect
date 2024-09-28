-- 1. Alle Benutzer auflisten
SELECT username, full_name, email FROM users;

-- 2. Alle Posts mit Autorennamen
SELECT p.title, p.content, u.full_name AS author
FROM posts p
JOIN users u ON p.user_id = u.user_id;

-- 3. Anzahl der Posts pro Benutzer
SELECT u.username, COUNT(p.post_id) AS post_count
FROM users u
LEFT JOIN posts p ON u.user_id = p.user_id
GROUP BY u.username
ORDER BY post_count DESC;

-- 4. Posts mit ihren Tags
SELECT p.title, STRING_AGG(t.name, ', ') AS tags
FROM posts p
JOIN post_tags pt ON p.post_id = pt.post_id
JOIN tags t ON pt.tag_id = t.tag_id
GROUP BY p.post_id, p.title;

-- 5. Die 3 beliebtesten Posts (basierend auf Likes)
SELECT p.title, COUNT(l.like_id) AS like_count
FROM posts p
LEFT JOIN likes l ON p.post_id = l.post_id
GROUP BY p.post_id, p.title
ORDER BY like_count DESC
LIMIT 3;

-- 6. Alle Kommentare zu einem bestimmten Post
SELECT c.content, u.username AS commenter
FROM comments c
JOIN users u ON c.user_id = u.user_id
WHERE c.post_id = (SELECT post_id FROM posts WHERE title = 'Kubernetes vs Docker Swarm');

-- 7. Benutzer und ihre Follower
SELECT u.username, COUNT(uc.follower_id) AS follower_count
FROM users u
LEFT JOIN user_connections uc ON u.user_id = uc.followed_id
GROUP BY u.user_id, u.username
ORDER BY follower_count DESC;

-- 8. Ungelesene Nachrichten für einen bestimmten Benutzer
SELECT m.content, u.username AS sender
FROM messages m
JOIN users u ON m.sender_id = u.user_id
WHERE m.receiver_id = (SELECT user_id FROM users WHERE username = 'johndoe')
AND m.is_read = FALSE;

-- 9. Posts mit Medieninhalten
SELECT title, media_type, media_url
FROM posts
WHERE media_type IS NOT NULL;

-- 10. Benutzer, die sowohl gepostet als auch kommentiert haben
SELECT DISTINCT u.username
FROM users u
JOIN posts p ON u.user_id = p.user_id
JOIN comments c ON u.user_id = c.user_id;

-- 11. Die am häufigsten verwendeten Tags
SELECT t.name, COUNT(pt.post_id) AS usage_count
FROM tags t
JOIN post_tags pt ON t.tag_id = pt.tag_id
GROUP BY t.tag_id, t.name
ORDER BY usage_count DESC
LIMIT 5;

-- 12. Durchschnittliche Anzahl von Kommentaren pro Post
SELECT AVG(comment_count) AS avg_comments_per_post
FROM (
    SELECT p.post_id, COUNT(c.comment_id) AS comment_count
    FROM posts p
    LEFT JOIN comments c ON p.post_id = c.post_id
    GROUP BY p.post_id
) AS post_comment_counts;
