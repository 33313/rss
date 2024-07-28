-- +goose Up
CREATE TABLE feeds (
    id uuid not null primary key,
    name text not null,
    url text unique not null,
    user_id uuid not null references users(id)
        on delete cascade
        on update cascade,
    created_at timestamp not null,
    updated_at timestamp not null,
    last_fetched_at timestamp default null
);

-- +goose Down
DROP TABLE feeds;
