-- +goose Up
CREATE TABLE categories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE categories;
