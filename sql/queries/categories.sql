-- name: CreateCategory :one
INSERT INTO categories (id, name, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetActiveCategories :many
SELECT * FROM categories
WHERE is_active = true;

-- name: UpdateCategoryName :one
UPDATE categories
SET name = $1, updated_at = NOW()
WHERE id = $2

-- name: SetCategoryActive :one
UPDATE categories
SET is_active = $1, updated_at = NOW()
WHERE id = $2
RETURNING *;