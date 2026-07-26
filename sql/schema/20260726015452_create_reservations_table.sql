-- +goose Up
CREATE TABLE reservations (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
	category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
	start_time TIMESTAMP NOT NULL,
	end_time TIMESTAMP NOT NULL,
	note TEXT,
	status TEXT NOT NULL DEFAULT 'confirmed' CHECK (status IN ('confirmed', 'cancelled')),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE reservations;
