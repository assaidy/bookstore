-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart (
    user_id int REFERENCES users(id) ON DELETE CASCADE,
    book_id int REFERENCES books(id) ON DELETE CASCADE,
    quantity int NOT NULL CHECK (quantity > 0), 
    price_per_unite numeric(10,2) NOT NULL,
    PRIMARY KEY(user_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd
