-- +goose Up
-- +goose StatementBegin
create table if not exists project
(
    id                bigserial primary key,
    name              text      default '' not null,
    short_name        text      default '' not null,
    is_archived       bool      default false,
    responsible_teams integer[] default array []::integer[],
    description       text      default '',
    department        bigserial            not null,
    responsible       text                 not null default ''
);
-- +goose StatementEnd
-- +goose StatementBegin
create unique index unique_project on project (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists project;
-- +goose StatementEnd
-- +goose StatementBegin
drop index if exists unique_project;
-- +goose StatementEnd