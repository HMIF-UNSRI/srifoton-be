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
    id_member3       uuid REFERENCES members ON DELETE SET NULL ON UPDATE CASCADE,
    id_member4       uuid REFERENCES members ON DELETE SET NULL ON UPDATE CASCADE,
    created_at       timestamp        NOT NULL DEFAULT (NOW()),
    updated_at       timestamp        NOT NULL DEFAULT (NOW())
);