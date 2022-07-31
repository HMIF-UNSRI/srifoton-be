CREATE TABLE IF NOT EXISTS users
(
    id                uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name              varchar(255)     NOT NULL,
    nim               varchar(255)     NOT NULL UNIQUE,
    email             varchar(255)     NOT NULL UNIQUE,
    password_hash     varchar(255)     NOT NULL,
    university        varchar(255)     NOT NULL,
    role              role_type        NOT NULL DEFAULT 'BASE',
    is_email_verified bool             NOT NULL DEFAULT FALSE,
    whatsapp_number   varchar(20)      NOT NULL,
    kpm_filename      varchar(255)     NOT NULL UNIQUE REFERENCES uploads (filename),
    created_at        timestamp        NOT NULL DEFAULT (NOW()),
    updated_at        timestamp        NOT NULL DEFAULT (NOW())
);