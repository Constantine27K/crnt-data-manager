-- +goose Up
-- +goose StatementBegin
create table if not exists sprint
(
    id                  bigserial primary key,
    name                text      default '' not null,
    project             integer   default 0  not null,
    start               timestamp            not null,
    finish              timestamp            not null,
    status              integer              not null,
    stories_backlog     integer   default 0,
    stories_in_progress integer   default 0,
    stories_done        integer   default 0,
    issues              integer[] default array []::integer[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists sprint;
-- +goose StatementEnd
