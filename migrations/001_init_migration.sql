CREATE TABLE IF NOT EXISTS public.feature (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    description TEXT NOT NULL CONSTRAINT valid_description CHECK (length(description) > 0),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.tag (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    "name" TEXT NOT NULL CONSTRAINT valid_name CHECK (length("name") > 0),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.banner (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    feature_id INT REFERENCES public.feature ON DELETE CASCADE,
    "content" JSONB NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.banner_tag (
    banner_id INT REFERENCES public.banner ON DELETE CASCADE,
    tag_id INT REFERENCES public.tag ON DELETE CASCADE,
    PRIMARY KEY (banner_id, tag_id)
);

---- create above / drop below ----
DROP TABLE IF EXISTS public.banner_tag;

DROP TABLE IF EXISTS public.tag;

DROP TABLE IF EXISTS public.banner;

DROP TABLE IF EXISTS public.feature;
