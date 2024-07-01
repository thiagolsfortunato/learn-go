USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  nick VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL,
  createdAt TIMESTAMP DEFAULT current_timestamp()
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS followers(
    user_id INT NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS posts(
    id INT auto_increment PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,

    author_id INT NOT NULL,
    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes INT DEFAULT 0,
    createdAt TIMESTAMP DEFAULT current_timestamp
) ENGINE=INNODB;
