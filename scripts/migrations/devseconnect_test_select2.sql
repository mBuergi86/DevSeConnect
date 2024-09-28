-- 1. Alle Netzwerkverbindungen anzeigen
SELECT 
    n.network_id,
    u1.username AS user1,
    u2.username AS user2,
    n.created_at
FROM 
    network n
JOIN 
    users u1 ON n.user_id1 = u1.user_id
JOIN 
    users u2 ON n.user_id2 = u2.user_id;

-- 2. Anzahl der Netzwerkverbindungen pro Benutzer
SELECT 
    u.username,
    COUNT(*) AS connection_count
FROM 
    users u
LEFT JOIN 
    (SELECT user_id1 AS user_id FROM network
     UNION ALL
     SELECT user_id2 FROM network) n ON u.user_id = n.user_id
GROUP BY 
    u.user_id, u.username
ORDER BY 
    connection_count DESC;

-- 3. Benutzer mit den meisten Netzwerkverbindungen
SELECT 
    u.username,
    COUNT(*) AS connection_count
FROM 
    users u
JOIN 
    (SELECT user_id1 AS user_id FROM network
     UNION ALL
     SELECT user_id2 FROM network) n ON u.user_id = n.user_id
GROUP BY 
    u.user_id, u.username
ORDER BY 
    connection_count DESC
LIMIT 5;

-- 4. Gemeinsame Verbindungen zwischen zwei Benutzern
WITH user1_connections AS (
    SELECT user_id2 AS connection_id FROM network WHERE user_id1 = (SELECT user_id FROM users WHERE username = 'johndoe')
    UNION
    SELECT user_id1 FROM network WHERE user_id2 = (SELECT user_id FROM users WHERE username = 'johndoe')
),
user2_connections AS (
    SELECT user_id2 AS connection_id FROM network WHERE user_id1 = (SELECT user_id FROM users WHERE username = 'janesmit')
    UNION
    SELECT user_id1 FROM network WHERE user_id2 = (SELECT user_id FROM users WHERE username = 'janesmit')
)
SELECT 
    u.username AS common_connection
FROM 
    user1_connections uc1
JOIN 
    user2_connections uc2 ON uc1.connection_id = uc2.connection_id
JOIN 
    users u ON uc1.connection_id = u.user_id;

-- 5. Neueste Netzwerkverbindungen
SELECT 
    u1.username AS user1,
    u2.username AS user2,
    n.created_at
FROM 
    network n
JOIN 
    users u1 ON n.user_id1 = u1.user_id
JOIN 
    users u2 ON n.user_id2 = u2.user_id
ORDER BY 
    n.created_at DESC
LIMIT 10;

-- 6. Benutzer ohne Netzwerkverbindungen
SELECT 
    u.username
FROM 
    users u
LEFT JOIN 
    (SELECT user_id1 AS user_id FROM network
     UNION
     SELECT user_id2 FROM network) n ON u.user_id = n.user_id
WHERE 
    n.user_id IS NULL;

-- 7. Durchschnittliche Anzahl von Netzwerkverbindungen pro Benutzer
SELECT 
    AVG(connection_count) AS avg_connections
FROM 
    (SELECT 
        u.user_id,
        COUNT(*) AS connection_count
    FROM 
        users u
    LEFT JOIN 
        (SELECT user_id1 AS user_id FROM network
         UNION ALL
         SELECT user_id2 FROM network) n ON u.user_id = n.user_id
    GROUP BY 
        u.user_id) AS user_connections;
