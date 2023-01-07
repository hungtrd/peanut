CREATE TABLE contents (
                       id bigserial NOT NULL,
                       "name" text NULL,
                       thumbnail text NULL,
                       description text NULL,
                       play_time text NULL,
                       resolution text NULL,
                       aspect_ratio text NULL,
                       tag text NULL,
                       category text NULL,
                       created_at timestamptz NULL,
                       updated_at timestamptz NULL,
                       deleted_at timestamptz NULL
);
CREATE INDEX idx_users_deleted_at ON users USING btree (deleted_at);