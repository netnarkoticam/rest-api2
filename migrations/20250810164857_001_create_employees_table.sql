-- +goose Up
-- +goose StatementBegin
CREATE TABLE employees(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
lastname TEXT NOT NULL,
firstname TEXT NOT NULL,
midlename TEXT NOT NULL,
phone TEXT NOT NULL,
address TEXT NOT NULL,
department TEXT NOT NULL,
hiredate DATE NOT NULL,
firedate DATE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
