-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd
