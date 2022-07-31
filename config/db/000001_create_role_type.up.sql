DO
$$
    BEGIN
        CREATE TYPE role_type AS ENUM ('BASE', 'ADMIN');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;
