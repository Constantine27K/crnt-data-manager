-- +goose Up
-- +goose StatementBegin
create table if not exists team
(
    id             bigserial primary key,
    name           text      default '',
    members        text[]    default array []::text[],
    projects       integer[] default array []::integer[],
    tech_owner     text      default '',
    business_owner text      default '',
    department     text      default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists team;
-- +goose StatementEnd
