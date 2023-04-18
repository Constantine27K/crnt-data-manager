-- +goose Up
-- +goose StatementBegin
create table if not exists issue
(
    id             bigserial primary key,
    composite_name text             default '' not null,
    name           text             default '',
    issue_type     integer          default 0  not null,
    parent_id      integer          default 0,
    description    text             default '',
    comments       JSONB            default '{}',
    author         text             default '' not null,
    assigned       text             default '',
    qa             text             default '',
    reviewer       text             default '',
    template       integer          default 0,
    created_at     timestamp        default now(),
    updated_at     timestamp        default now(),
    deadline       timestamp        default now(),
    status         integer          default 0  not null,
    priority       integer          default 0  not null,
    sprint_id      integer          default 0,
    project_id     integer          default 0,
    components     integer[]        default array []::integer[],
    story_points   integer          default 0,
    payment        double precision default 0,
    time_spent     integer          default 0,
    children       integer[]        default array []::integer[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists issue;
-- +goose StatementEnd
