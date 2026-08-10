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

-- name: GetReservationsByUser :many 
SELECT * FROM reservations
WHERE user_id = $1;

-- name: GetAllReservations :many
SELECT * FROM reservations;

-- name: CancelReservation :one
UPDATE reservations
SET status = 'cancelled', updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: CancelReservationByAdmin :one
UPDATE reservations
SET status = 'cancelled', updated_at = NOW()
WHERE id = $1
RETURNING *;