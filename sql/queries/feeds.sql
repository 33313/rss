-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at, last_fetched_at)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetFeeds :many
select * from feeds;

-- name: GetNextFeedsToFetch :many
select * from feeds
order by last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedFetched :exec
update feeds set last_fetched_at = NOW(), updated_at = NOW()
where id = $1;
