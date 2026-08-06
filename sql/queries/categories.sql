-- name: CreateCategory :one
INSERT INTO categories (id, name, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    NOW(),
    NOW()
)
RETURNING *;