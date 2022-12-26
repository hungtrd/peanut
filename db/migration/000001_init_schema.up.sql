CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	username text NULL,
	email text NULL,
	password text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
