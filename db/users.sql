-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    user_id integer NOT NULL,
    username character varying(20) COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;
	
-- SEQUENCE: public.user_user_id_seq

-- DROP SEQUENCE IF EXISTS public.user_user_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.user_user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY users.user_id;

ALTER SEQUENCE public.user_user_id_seq
    OWNER TO postgres;
	
ALTER TABLE IF EXISTS public.users
    ALTER COLUMN user_id SET DEFAULT nextval('user_user_id_seq'::regclass);