-- Table: public.tokens

-- DROP TABLE IF EXISTS public.tokens;

CREATE TABLE IF NOT EXISTS public.tokens
(
    id integer NOT NULL ,
    token character varying COLLATE pg_catalog."default" NOT NULL,
    expire_at timestamp with time zone NOT NULL,
    status boolean,
    user_id bigint,
    CONSTRAINT token_blacklist_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tokens
    OWNER to postgres;
	
-- SEQUENCE: public.token_id_seq

-- DROP SEQUENCE IF EXISTS public.token_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.token_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY tokens.id;

ALTER SEQUENCE public.token_id_seq
    OWNER TO postgres;
	
ALTER TABLE IF EXISTS public.tokens
    ALTER COLUMN id SET DEFAULT nextval('token_id_seq'::regclass);