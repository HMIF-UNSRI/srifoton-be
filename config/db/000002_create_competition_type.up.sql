DO
$$
    BEGIN
        CREATE TYPE competition_type AS ENUM ('CP', 'WEB','UI/UX','ESPORT');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;
