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

func (dbs *DBService) CreateUser(inout *models.User) error {
	query := `
    INSERT INTO users(name, username, password, email, address, joined_at)
    VALUES($1, $2, $3, $4, $5, $6)
    RETURNING id;
    `
	if err := dbs.db.QueryRow(
		query,
		inout.Name,
		inout.Username,
		inout.Password,
		inout.Email,
		inout.Address,
		inout.JoinedAt,
	).Scan(&inout.Id); err != nil {
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

func (dbs *DBService) GetAllUsers() ([]*models.User, error) {
	query := `
    SELECT
        id,
        name,
        username,
        email,
        Address,
        joined_at
    FROM users;
    `
	rows, err := dbs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Address,
			&user.JoinedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (dbs *DBService) UpdateUser(newUser *models.User) error {
	query := `
    UPDATE users
    SET 
        name = $1,
        username = $2,
        email = $3,
        password = $4,
        address = $5
    WHERE id = $6;
    `
	if _, err := dbs.db.Exec(
		query,
		newUser.Name,
		newUser.Username,
		newUser.Email,
		newUser.Password,
		newUser.Address,
		newUser.Id,
	); err != nil {
		return err
	}

	return nil
}

func (dbs *DBService) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1;`
	if _, err := dbs.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------
// > category
// --------------------------------------------------
func (dbs *DBService) CheckCategoryConflict(name string) (bool, error) {
	query := `SELECT 1 FROM categories WHERE name = $1 LIMIT 1;`
	return dbs.checkRow(query, name)
}

func (dbs *DBService) CheckIfCategoryExists(id int) (bool, error) {
	query := `SELECT 1 FROM categories WHERE id = $1 LIMIT 1;`
	return dbs.checkRow(query, id)
}

func (dbs *DBService) CreateCategory(inout *models.Category) error {
	query := `
    INSERT INTO categories(name)
    VALUES($1)
    RETURNING id;
    `
	if err := dbs.db.QueryRow(query, inout.Name).Scan(&inout.Id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetAllCategories() ([]*models.Category, error) {
	query := `SELECT id, name FROM categories;`
	rows, err := dbs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cats := make([]*models.Category, 0)
	for rows.Next() {
		cat := models.Category{}
		if err := rows.Scan(&cat.Id, &cat.Name); err != nil {
			return nil, err
		}
		cats = append(cats, &cat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cats, nil
}

func (dbs *DBService) GetCategoryById(id int) (*models.Category, error) {
	query := `SELECT name FROM categories WHERE id = $1;`
	cat := models.Category{Id: id}
	if err := dbs.db.QueryRow(query, id).Scan(&cat.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cat, nil
}

func (dbs *DBService) UpdateCategory(cat *models.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2;`
	if _, err := dbs.db.Exec(query, cat.Name, cat.Id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1;`
	if _, err := dbs.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}
