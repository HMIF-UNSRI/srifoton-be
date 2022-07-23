CREATE TYPE role_type AS ENUM ('BASE', 'ADMIN');
CREATE TYPE competition_type AS ENUM ('CP', 'WEB','UI/UX','ESPORT');

CREATE TABLE IF NOT EXISTS uploads
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    filename   varchar(255)     NOT NULL UNIQUE,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);

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


CREATE TABLE IF NOT EXISTS teams
(
    id               uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    id_lead          uuid             NOT NULL UNIQUE REFERENCES users,
    name             varchar(255)     NOT NULL,
    competition      competition_type NOT NULL,
    is_confirmed     bool             NOT NULL DEFAULT FALSE,
    payment_filename varchar(255)     NOT NULL REFERENCES uploads (filename),
    id_member1       uuid REFERENCES members ON DELETE SET NULL ON UPDATE CASCADE,
    id_member2       uuid REFERENCES members ON DELETE SET NULL ON UPDATE CASCADE,
    created_at       timestamp        NOT NULL DEFAULT (NOW()),
    updated_at       timestamp        NOT NULL DEFAULT (NOW())
);

