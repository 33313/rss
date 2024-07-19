-- +goose Up
create table feed_follows (
    id uuid primary key not null,
    feed_id uuid not null references feeds(id)
        on update cascade
        on delete cascade,
    user_id uuid not null references users(id)
        on update cascade
        on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null,
    unique(feed_id, user_id)
);

-- +goose Down
drop table feed_follows;
