CREATE TABLE IF NOT EXISTS books (
	id bigserial NOT NULL,
	title text NULL,
	description text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
);
