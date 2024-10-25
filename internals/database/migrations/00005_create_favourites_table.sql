-- +goose Up
-- +goose StatementBegin
CREATE TABLE favourites (
    user_id int REFERENCES users(id) ON DELETE CASCADE, 
    book_id int REFERENCES books(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE favourites;
-- +goose StatementEnd
