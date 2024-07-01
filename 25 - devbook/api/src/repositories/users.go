package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create User
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (NAME, NICK, EMAIL, PASSWORD) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

// Search User by Name or Nick (query string)
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNic%

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Search User by Id
func (repository Users) SearchById(Id uint64) (models.User, error) {
	lines, err := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", Id)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Search User by Email
func (repository Users) SearchByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update User
func (repository Users) Update(Id uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, Id); err != nil {
		return err
	}

	return nil
}

// Delete User
func (repository Users) Delete(Id uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(Id); err != nil {
		return err
	}

	return nil
}

// Follow User
func (repository Users) FollowUser(userId, followerId uint64) error {
	statement, err := repository.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

// Unfollow User
func (repository Users) UnfollowUser(userId, followerId uint64) error {
	statement, err := repository.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

// Get Followers
func (repository Users) GetFollowers(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt 
		from users u inner join followers f on u.id = f.follower_id 
		where f.user_id = ?`, userId)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get Following
func (repository Users) GetFollowing(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt 
		from users u inner join followers f 
		on u.id = f.user_id where f.user_id = ?`, userId)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get Password
func (repository Users) GetPassword(userId uint64) (string, error) {
	line, err := repository.db.Query("select password from users where id = ?", userId)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// Update Password
func (repository Users) UpdatePassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
