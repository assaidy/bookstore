-- +goose Up
-- +goose StatementBegin
CREATE TABLE covers (
    id serial PRIMARY KEY,
    encoding varchar(255) NOT NULL, -- Stores encoding type (e.g., "image/jpeg", "image/png")
    content text NOT NULL           -- Stores the actual binary image data
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE covers;
-- +goose StatementEnd
