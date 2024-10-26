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

// --------------------------------------------------
// > cover
// --------------------------------------------------
func (dbs *DBService) CreateCover(inout *models.Cover) error {
	query := `
    INSERT INTO covers (encoding, content)
    VALUES ($1, $2)
    RETURNING id;
    `
	if err := dbs.db.QueryRow(
		query,
		inout.Encoding,
		inout.Content,
	).Scan(&inout.Id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetCoverById(id int) (*models.Cover, error) {
	query := `SELECT encoding, content FROM covers WHERE id = $1;`
	cov := models.Cover{Id: id}
	if err := dbs.db.QueryRow(query, id).Scan(&cov.Encoding, &cov.Content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cov, nil
}

func (dbs *DBService) UpdateCover(cov *models.Cover) error {
	query := `UPDATE covers SET encoding = $1, content = $2 WHERE id = $3;`
	if _, err := dbs.db.Exec(query, cov.Encoding, cov.Content, cov.Id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) CheckIfCoverExists(id int) (bool, error) {
	query := `SELECT 1 FROM covers WHERE id = $1 LIMIT 1;`
	return dbs.checkRow(query, id)
}

func (dbs *DBService) DeleteCover(id int) error {
	query := `DELETE FROM covers WHERE id = $1;`
	if _, err := dbs.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------
// > book
// --------------------------------------------------
func (dbs *DBService) CreateBook(inout *models.Book) error {
	query := `
    INSERT INTO books(
        title,
        description,
        category_id,
        cover_id,
        price,
        quantity,
        discount,
        added_at
    )
    VALUES($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id;
    `
	if err := dbs.db.QueryRow(
		query,
		inout.Title,
		inout.Description,
		inout.CategoryId,
		inout.CoverId,
		inout.Price,
		inout.Quantity,
		inout.Discount,
		inout.AddedAt,
	).Scan(&inout.Id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetBookById(id int) (*models.Book, error) {
	query := `
    SELECT
        title,
        description,
        category_id,
        cover_id,
        price,
        quantity,
        discount,
        added_at,
        purchase_count
    FROM books
    WHERE id = $1;
    `
	book := models.Book{Id: id}
	if err := dbs.db.QueryRow(query, id).Scan(
		&book.Title,
		&book.Description,
		&book.CategoryId,
		&book.CoverId,
		&book.Price,
		&book.Quantity,
		&book.Discount,
		&book.AddedAt,
		&book.PurchaseCount,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (dbs *DBService) GetAllBooksByCategory(cid int) ([]*models.Book, error) {
	query := `
    SELECT
        id,
        title,
        description,
        cover_id,
        price,
        quantity,
        discount,
        added_at,
        purchase_count
    FROM books
    WHERE category_id = $1; 
    `
	rows, err := dbs.db.Query(query, cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*models.Book, 0)

	for rows.Next() {
		book := models.Book{CategoryId: cid}
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.CoverId,
			&book.Price,
			&book.Quantity,
			&book.Discount,
			&book.AddedAt,
			&book.PurchaseCount,
		); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (dbs *DBService) UpdateBook(book *models.Book) error {
	query := `
    UPDATE books
    SET 
        title = $1,
        description = $2,
        category_id = $3,
        price = $4,
        quantity = $5,
        discount = $6,
    WHERE id = $7;
    `
	if _, err := dbs.db.Exec(
		query,
		book.Title,
		book.Description,
		book.CategoryId,
		book.Price,
		book.Quantity,
		book.Discount,
		book.Id,
	); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) CheckIfBookExists(id int) (bool, error) {
	query := `SELECT 1 FROM books WHERE id = $1 LIMIT 1;`
	return dbs.checkRow(query, id)
}

func (dbs *DBService) DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id = $1;`
	if _, err := dbs.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetAllBooks() ([]*models.Book, error) {
	query := `
    SELECT
        id,
        title,
        description,
        category_id,
        cover_id,
        price,
        quantity,
        discount,
        added_at,
        purchase_count
    FROM books;
    `
	rows, err := dbs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*models.Book, 0)

	for rows.Next() {
		book := models.Book{}
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.CategoryId,
			&book.CoverId,
			&book.Price,
			&book.Quantity,
			&book.Discount,
			&book.AddedAt,
			&book.PurchaseCount,
		); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// --------------------------------------------------
// > favourites
// --------------------------------------------------
func (dbs *DBService) AddBookToFavourites(uid, bid int) error {
	query := `INSERT INTO favourites (user_id, book_id) VALUES ($1, $2);`
	if _, err := dbs.db.Exec(query, uid, bid); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetAllBooksInFavourites(uid int) ([]*models.Book, error) {
	query := `SELECT book_id FROM favourites WHERE user_id = $1;`

	rows, err := dbs.db.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*models.Book, 0)

	for rows.Next() {
		book := models.Book{}
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.CategoryId,
			&book.CoverId,
			&book.Price,
			&book.Quantity,
			&book.Discount,
			&book.AddedAt,
			&book.PurchaseCount,
		); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (dbs *DBService) DeleteBookFromFavourites(uid, bid int) error {
	query := `DELETE FROM favourites WHERE user_id = $1 AND book_id = $2;`
	if _, err := dbs.db.Exec(query, uid, bid); err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------
// > cart
// --------------------------------------------------
func (dbs *DBService) AddBookToCart(uid, bid, quantity int) error {
	query := `
    INSERT INTO cart (user_id, book_id, quantity)
    VALUES ($1, $2, $3);
    `
	if _, err := dbs.db.Exec(query, uid, bid, quantity); err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) GetBookFromCart(uid, bid int) (*models.CartBook, error) {
	query := `SELECT quantity FROM cart WHERE uid = $1 AND bid = $2;`
	cb := models.CartBook{UserId: uid, BookId: bid}
	if err := dbs.db.QueryRow(query, uid, bid).Scan(&cb.Quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cb, nil
}

func (dbs *DBService) GetBooksInCart(uid int) ([]*models.Book, error) {
	query := `SELECT book_id FROM favourites WHERE user_id = $1;`

	rows, err := dbs.db.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*models.Book, 0)

	for rows.Next() {
		book := models.Book{}
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.CategoryId,
			&book.CoverId,
			&book.Price,
			&book.Quantity,
			&book.Discount,
			&book.AddedAt,
			&book.PurchaseCount,
		); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (dbs *DBService) DeleteBookFromCart(uid, bid int) error {
	query := `DELETE FROM cart WHERE user_id = $1 AND book_id = $2;`
	if _, err := dbs.db.Exec(query, uid, bid); err != nil {
		return err
	}
	return nil
}
