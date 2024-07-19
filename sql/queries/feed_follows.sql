-- name: CreateFollow :one
insert into feed_follows (id, feed_id, user_id, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetUserFollows :many
select * from feed_follows
where user_id = $1;

-- name: DeleteFollow :exec
delete from feed_follows
where id = $1;

-- name: DeleteAllUserFollows :execrows
delete from feed_follows
where user_id = $1;
