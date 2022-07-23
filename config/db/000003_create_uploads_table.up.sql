CREATE TABLE IF NOT EXISTS uploads
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    filename   varchar(255)     NOT NULL UNIQUE,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);