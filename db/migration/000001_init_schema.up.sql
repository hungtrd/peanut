CREATE TABLE users (
    id bigserial NOT NULL,
    username text NULL,
    email text NULL,
    "password" text NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON users USING btree (deleted_at);