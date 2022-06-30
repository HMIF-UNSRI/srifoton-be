
CREATE TYPE role_type AS ENUM ('BASE', 'ADMIN');
CREATE TYPE competition_type AS ENUM ('CP', 'WEB','UI/UX','ESPORT');

CREATE TABLE IF NOT EXISTS uploads
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    file_name  varchar(255)     NOT NULL UNIQUE,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);

CREATE TABLE IF NOT EXISTS users
(
    id                uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    id_kpm            uuid             NOT NULL UNIQUE REFERENCES uploads,
    nama              varchar(255)     NOT NULL,
    nim               varchar(255)     NOT NULL UNIQUE,
    email             varchar(255)     NOT NULL UNIQUE,
    password          varchar(255)     NOT NULL,
    university        varchar(255)     NOT NULL,
    role              role_type        NOT NULL DEFAULT 'BASE',
    is_email_verified bool             NOT NULL DEFAULT FALSE,
    no_wa             varchar(255)     NOT NULL,
    created_at        timestamp        NOT NULL DEFAULT (NOW()),
    updated_at        timestamp        NOT NULL DEFAULT (NOW())
);

CREATE TABLE IF NOT EXISTS members
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    id_kpm     uuid             NOT NULL UNIQUE REFERENCES uploads,
    nama       varchar(255)     NOT NULL,
    nim        varchar(255)     NOT NULL,
    email      varchar(255)     NOT NULL UNIQUE,
    university varchar(255)     NOT NULL,
    no_wa      varchar(255)     NOT NULL,
    created_at timestamp        NOT NULL DEFAULT (NOW()),
    updated_at timestamp        NOT NULL DEFAULT (NOW())
);



CREATE TABLE IF NOT EXISTS teams
(
    id           uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    team_name    varchar(255)     NOT NULL,
    id_lead      uuid             NOT NULL REFERENCES users,
    competition  competition_type NOT NULL,
    id_member_1  uuid REFERENCES members,
    id_member_2  uuid REFERENCES members,
    id_payment   uuid REFERENCES uploads,
    is_confirmed bool             NOT NULL DEFAULT FALSE,
    created_at   timestamp        NOT NULL DEFAULT (NOW()),
    updated_at   timestamp        NOT NULL DEFAULT (NOW())
);

