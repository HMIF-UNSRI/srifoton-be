CREATE TYPE role_type AS ENUM ('BASE', 'ADMIN');

CREATE TABLE IF NOT EXISTS users
(
    id                uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    email             varchar(255)     NOT NULL UNIQUE,
    password          varchar(255)     NOT NULL,
    role              role_type        NOT NULL DEFAULT 'BASE',
    is_email_verified bool             NOT NULL DEFAULT FALSE,
    created_at        timestamp        NOT NULL DEFAULT (NOW()),
    updated_at        timestamp        NOT NULL DEFAULT (NOW())
);

DROP TABLE IF EXISTS users;