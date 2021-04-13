CREATE TABLE public.users (
    id SERIAL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id)
);