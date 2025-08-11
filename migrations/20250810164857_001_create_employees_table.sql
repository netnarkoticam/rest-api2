-- +goose Up
-- +goose StatementBegin
CREATE TABLE employees(
    id SERIAL PRIMARY KEY,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    middle_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    address TEXT NOT NULL,
    department TEXT NOT NULL,
    hire_date TIMESTAMP NOT NULL DEFAULT now(),
    fire_date TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
