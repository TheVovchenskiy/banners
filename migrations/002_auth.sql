CREATE TABLE IF NOT EXISTS public.role (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    "name" TEXT UNIQUE NOT NULL CONSTRAINT valid_name CHECK (
        length("name") > 0
        AND length("name") <= 50
    ),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.user_profile (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    role_id INT REFERENCES public.role ON DELETE CASCADE,
    username TEXT UNIQUE NOT NULL CONSTRAINT valid_username CHECK (
        length(username) > 0
        AND length(username) <= 150
    ),
    password_hash bytea NOT NULL,
    salt bytea NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

INSERT INTO
    public.role ("name")
VALUES
    ('admin'),
    ('user');

---- create above / drop below ----
DROP TABLE IF EXISTS public.user_profile;

DROP TABLE IF EXISTS public.role;
