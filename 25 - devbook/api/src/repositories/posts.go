package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create Post
func (repository Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

// Get Post by Id
func (repository Posts) GetPostById(postId uint64) (models.Post, error) {
	lines, err := repository.db.Query(
		`select p.*, u.nick from 
		posts p inner join users u 
		on p.author_id = u.id
		where p.id = ?`, postId)

	if err != nil {
		return models.Post{}, err
	}

	defer lines.Close()

	var post models.Post

	if lines.Next() {
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// Get All Posts by User ID (own and following)
func (repository Posts) GetPosts(userId uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(
		`select distinct p.*, u.nick from posts p 
		join users u on u.id = p.author_id 
		join followers f on p.author_id = f.user_id 
		where u.id = ? or f.follower_id = ?
		order by 1 desc`, userId, userId,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()
	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update Post
func (repository Posts) UpdatePost(postId uint64, post models.Post) error {
	statement, err := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

// Delete Post
func (repository Posts) DeletePost(postId uint64) error {
	statement, err := repository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

// Get Posts by User
func (repository Posts) GetPostsByUser(userId uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from posts p
		join users u on u.id = p.author_id
		where p.author_id = ?`, userId,
	)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) LikePost(postId uint64) error {
	statement, err := repository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (repository Posts) DislikePost(postId uint64) error {
	statement, err := repository.db.Prepare("update posts set likes = case when likes > 0 then likes - 1 else 0 end where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
