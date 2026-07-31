-- +goose Up
CREATE TABLE admin_actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    action TEXT NOT NULL,
    target_reservation_id UUID DEFAULT NULL REFERENCES reservations(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE admin_actions;
