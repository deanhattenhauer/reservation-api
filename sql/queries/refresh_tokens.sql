-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, user_id, expires_at)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT users.* FROM users
INNER JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE expires_at > NOW() AND revoked_at IS NULL AND refresh_tokens.token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW()
WHERE token = $1;