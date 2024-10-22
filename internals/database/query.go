package database

import (
	"database/sql"
	"errors"
	"github.com/assaidy/bookstore/internals/models"
)

// --------------------------------------------------
// > user
// --------------------------------------------------
func (dbs *DBService) CheckUsernameAndEmailConflict(username string, email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = $1 OR email = $2 LIMIT 1;`
	return dbs.checkRow(query, username, email)
}

func (dbs *DBService) CheckUsernameConflict(username string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = $1 LIMIT 1;`
	return dbs.checkRow(query, username)
}

func (dbs *DBService) CheckEmailConflict(email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = $1 LIMIT 1;`
	return dbs.checkRow(query, email)
}

func (dbs *DBService) CheckIfUserExists(id int) (bool, error) {
	query := `SELECT 1 FROM users WHERE id = $1 LIMIT 1;`
	return dbs.checkRow(query, id)
}

func (dbs *DBService) checkRow(query string, args ...any) (bool, error) {
	if err := dbs.db.QueryRow(query, args...).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (dbs *DBService) CreateUser(user *models.User) error {
	query := `
    INSERT INTO users(name, username, password, email, address, joined_at)
    VALUES($1, $2, $3, $4, $5, $6);
    `
	if _, err := dbs.db.Exec(query,
		user.Name, user.Username, user.Password, user.Email, user.Address, user.JoinedAt); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetUserById(id int) (*models.User, error) {
	query := `
    SELECT
        name,
        email,
        username,
        password,
        address,
        joined_at
    FROM users
    WHERE id = $1;
    `
	user := models.User{Id: id}
	if err := dbs.db.QueryRow(query, id).Scan(
		&user.Name, &user.Email, &user.Username, &user.Password, &user.Address, &user.JoinedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (dbs *DBService) GetUserByUsername(username string) (*models.User, error) {
	query := `
    SELECT
        id,
        name,
        email,
        password,
        Address,
        joined_at
    FROM users
    WHERE username = $1;
    `
	user := models.User{Username: username}
	if err := dbs.db.QueryRow(query, username).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.Address, &user.JoinedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
