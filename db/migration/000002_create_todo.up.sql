CREATE TABLE todos (
    id bigserial NOT NULL,
    user_id int8 NULL,
    title varchar(30) NULL,
    "content" varchar(500) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT todos_pkey PRIMARY KEY (id),
    CONSTRAINT fk_users_todos FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX idx_todos_deleted_at ON todos USING btree (deleted_at);
CREATE INDEX idx_todos_title ON todos USING btree (title);