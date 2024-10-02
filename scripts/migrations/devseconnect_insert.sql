-- Users
INSERT INTO users (username, email, password_hash, first_name, last_name, bio, profile_picture)
VALUES 
('johndoe', 'john.doe@email.com', 'hashed_password_1', 'John', 'Doe', 'DevOps engineer with 5 years of experience', 'profile1.jpg'),
('janesmit', 'jane.smith@email.com', 'hashed_password_2', 'Jane', 'Smith', 'Full-stack developer passionate about cloud technologies', 'profile2.jpg'),
('mikebrown', 'mike.brown@email.com', 'hashed_password_3', 'Mike', 'Brown', 'Software architect specializing in microservices', 'profile3.jpg');

-- Posts
-- Insert posts
INSERT INTO posts (user_id, title, content, media_type, media_url)
VALUES 
((SELECT user_id FROM users WHERE username = 'johndoe'), 'Best practices for CI/CD pipelines', 
 'In this post, I will share some best practices for setting up efficient CI/CD pipelines...', 'image', 'cicd_diagram.png'),

((SELECT user_id FROM users WHERE username = 'janesmit'), 'Kubernetes vs Docker Swarm', 
 'Let''s compare two popular container orchestration platforms...', NULL, NULL),

((SELECT user_id FROM users WHERE username = 'mikebrown'), 'Microservices architecture patterns', 
 'Here are some common patterns I''ve encountered in microservices architectures...', 'image', 'microservices_patterns.jpg');

-- Comments
-- Insert comments
INSERT INTO comments (post_id, user_id, content)
VALUES 
((SELECT post_id FROM posts WHERE title = 'Best practices for CI/CD pipelines'), 
 (SELECT user_id FROM users WHERE username = 'janesmit'), 
 'Great post! I would also add that it''s important to have proper error handling in your pipelines.'),

((SELECT post_id FROM posts WHERE title = 'Kubernetes vs Docker Swarm'), 
 (SELECT user_id FROM users WHERE username = 'mikebrown'), 
 'Interesting comparison. In my experience, Kubernetes has been more suitable for larger, more complex deployments.'),

((SELECT post_id FROM posts WHERE title = 'Microservices architecture patterns'), 
 (SELECT user_id FROM users WHERE username = 'johndoe'), 
 'The saga pattern has been particularly useful in our recent projects.');

-- Tags
INSERT INTO tags (name) VALUES ('DevOps'), ('CI/CD'), ('Kubernetes'), ('Docker'), ('Microservices');

-- Post Tags
INSERT INTO post_tags (post_id, tag_id)
VALUES 
((SELECT post_id FROM posts WHERE title = 'Best practices for CI/CD pipelines'), (SELECT tag_id FROM tags WHERE name = 'DevOps')),
((SELECT post_id FROM posts WHERE title = 'Best practices for CI/CD pipelines'), (SELECT tag_id FROM tags WHERE name = 'CI/CD')),
((SELECT post_id FROM posts WHERE title = 'Kubernetes vs Docker Swarm'), (SELECT tag_id FROM tags WHERE name = 'Kubernetes')),
((SELECT post_id FROM posts WHERE title = 'Kubernetes vs Docker Swarm'), (SELECT tag_id FROM tags WHERE name = 'Docker')),
((SELECT post_id FROM posts WHERE title = 'Microservices architecture patterns'), (SELECT tag_id FROM tags WHERE name = 'Microservices'));

-- Likes
INSERT INTO likes (post_id, user_id)
VALUES 
((SELECT post_id FROM posts WHERE title = 'Best practices for CI/CD pipelines'), (SELECT user_id FROM users WHERE username = 'janesmit')),
((SELECT post_id FROM posts WHERE title = 'Kubernetes vs Docker Swarm'), (SELECT user_id FROM users WHERE username = 'johndoe')),
((SELECT post_id FROM posts WHERE title = 'Microservices architecture patterns'), (SELECT user_id FROM users WHERE username = 'janesmit'));

-- Network
INSERT INTO network (user_id1, user_id2)
VALUES 
((SELECT user_id FROM users WHERE username = 'johndoe'), (SELECT user_id FROM users WHERE username = 'janesmit')),
((SELECT user_id FROM users WHERE username = 'johndoe'), (SELECT user_id FROM users WHERE username = 'mikebrown')),
((SELECT user_id FROM users WHERE username = 'janesmit'), (SELECT user_id FROM users WHERE username = 'mikebrown'));

-- User Connections
INSERT INTO user_connections (follower_id, followed_id)
VALUES 
((SELECT user_id FROM users WHERE username = 'johndoe'), (SELECT user_id FROM users WHERE username = 'janesmit')),
((SELECT user_id FROM users WHERE username = 'janesmit'), (SELECT user_id FROM users WHERE username = 'mikebrown')),
((SELECT user_id FROM users WHERE username = 'mikebrown'), (SELECT user_id FROM users WHERE username = 'johndoe'));

-- Messages
INSERT INTO messages (sender_id, receiver_id, content)
VALUES 
((SELECT user_id FROM users WHERE username = 'johndoe'), (SELECT user_id FROM users WHERE username = 'janesmit'), 'Hey Jane, I loved your post about Kubernetes. Can we discuss it further?'),
((SELECT user_id FROM users WHERE username = 'janesmit'), (SELECT user_id FROM users WHERE username = 'mikebrown'), 'Mike, I have a question about your microservices post. Do you have time for a quick call?'),
((SELECT user_id FROM users WHERE username = 'mikebrown'), (SELECT user_id FROM users WHERE username = 'johndoe'), 'John, I''m working on a new CI/CD setup. Can I get your opinion on it?');
