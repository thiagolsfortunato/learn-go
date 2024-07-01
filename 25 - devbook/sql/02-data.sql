insert into users (name, nick, email, password) values
  ("User1", "user_1", "user1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- user1
  ("User2", "user_2", "user2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- user2
  ("User3", "user_3", "user3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- user3

insert into followers(user_id, follower_id) values (1, 2), (1, 3), (2, 1);
insert into followers(user_id, follower_id) values (2, 3), (3, 1);

insert into posts(title, content, author_id) values
("Post from user_1", "This post is from user_1!", 1),
("Post from user_2", "This post is from user_2!", 2),
("Post from user_3", "This post is from user_3!", 3);
