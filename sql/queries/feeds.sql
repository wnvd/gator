-- name: CreateFeed :one
INSERT INTO feeds (id, url, name, user_id, created_at, updated_at, last_fetched_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE feeds.id = $1;

-- name: DeleteAllFeeds :exec
DROP TABLE feeds;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, 
    feeds.name AS feed_name,
    users.name As user_name
FROM 
    inserted_feed_follow 
INNER JOIN 
    feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN 
    users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name As feed_name, users.name As user_name
FROM 
    feed_follows
INNER JOIN 
    feeds ON feeds.id = feed_follows.feed_id
INNER JOIN 
    users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows 
WHERE user_id = $1 AND feed_id = $2;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at  = $1, updated_at = $2
WHERE  id = $3;

-- name: GetNextFeedToFetch :one
SELECT * 
FROM
    feeds
ORDER BY
    last_fetched_at ASC NULLS FIRST
LIMIT 1;
