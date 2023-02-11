-- +goose Up
-- +goose StatementBegin
create table if not exists sprint
(
    id          bigserial primary key,
    name        text      default '' not null,
    project     integer   default 0  not null,
    started_at  timestamp            not null,
    finished_at timestamp            not null,
    status      integer              not null,
    in_backlog  integer   default 0,
    in_progress integer   default 0,
    in_done     integer   default 0,
    issues      integer[] default array []::integer[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists sprint;
-- +goose StatementEnd
