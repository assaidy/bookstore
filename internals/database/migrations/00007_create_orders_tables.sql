-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id),
    applied_at timestamp NOT NULL DEFAULT NOW(),
    total_price numeric(10,2) NOT NULL
);

CREATE TABLE order_book (
    order_id int REFERENCES orders(id) ON DELETE CASCADE,
    book_id int REFERENCES books(id) ON DELETE CASCADE,
    quantity int NOT NULL CHECK (quantity > 0), 
    price_per_unit numeric(10,2) NOT NULL,
    PRIMARY KEY(order_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS order_book;
-- +goose StatementEnd
