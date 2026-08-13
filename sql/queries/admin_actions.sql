-- name: LogAdminCancelReservation :one
INSERT INTO admin_actions (id, admin_user_id, action, target_reservation_id, created_at)
VALUES (
    gen_random_uuid(),
    $1,
    'cancel_reservation',
    $2,
    NOW()
)
RETURNING *;