CREATE TABLE IF NOT EXISTS members
(
    id              uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name            varchar(255)     NOT NULL,
    nim             varchar(255)     NOT NULL,
    email           varchar(255)     NOT NULL UNIQUE,
    university      varchar(255)     NOT NULL,
    whatsapp_number varchar(20)      NOT NULL,
    kpm_filename    varchar(255)     NOT NULL UNIQUE REFERENCES uploads (filename),
    created_at      timestamp        NOT NULL DEFAULT (NOW()),
    updated_at      timestamp        NOT NULL DEFAULT (NOW())
);