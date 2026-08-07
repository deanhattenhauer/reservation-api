-- name: CreateReservation :one
INSERT INTO reservations (id, user_id, category_id, start_time, end_time, note, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW()
)
RETURNING *;