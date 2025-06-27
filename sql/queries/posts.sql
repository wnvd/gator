-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT 
    posts.id AS post_id,
    posts.title AS post_title,
    posts.url AS post_url,
    posts.description AS post_desc,
    posts.feed_id AS posts_feed_id
FROM 
    posts
INNER JOIN
    feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE
    feed_follows.user_id = $1
ORDER BY
    posts.updated_at DESC
LIMIT $2;
