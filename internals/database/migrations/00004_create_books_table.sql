-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    category_id int NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    cover_id int REFERENCES covers(id) ON DELETE SET NULL,
    price numeric(10,2) NOT NULL,
    quantity int NOT NULL,
    discount numeric(5,2) NOT NULL DEFAULT 0 CHECK (discount <= 100 AND discount >= 0),
    added_at timestamp NOT NULL DEFAULT NOW(),
    purchase_count int DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
