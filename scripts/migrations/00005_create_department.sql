-- +goose Up
-- +goose StatementBegin
create table if not exists department
(
    id       bigserial primary key,
    name     text not null default '',
    projects integer[]     default array []::integer[]
);
-- +goose StatementEnd
-- +goose StatementBegin
create unique index unique_department on department (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists department;
-- +goose StatementEnd
-- +goose StatementBegin
drop index if exists unique_department;
-- +goose StatementEnd