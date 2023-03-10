-- +goose Up
-- +goose StatementBegin
create table if not exists project
(
    id                bigserial primary key,
    name              text      default '' not null,
    short_name        text      default '' not null,
    is_archived       bool      default false,
    responsible_teams integer[] default array []::integer[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists project;
-- +goose StatementEnd
