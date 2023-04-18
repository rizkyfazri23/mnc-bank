-- Table: public.authentication_logs

-- DROP TABLE IF EXISTS public.authentication_logs;

CREATE TABLE IF NOT EXISTS public.authentication_logs
(
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    status character varying COLLATE pg_catalog."default" NOT NULL,
    date_time time without time zone NOT NULL DEFAULT 'now()',
    CONSTRAINT authentication_logs_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.authentication_logs
    OWNER to postgres;
	
-- SEQUENCE: public.authentication_logs_id_seq

-- DROP SEQUENCE IF EXISTS public.authentication_logs_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.authentication_logs_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY authentication_logs.id;

ALTER SEQUENCE public.authentication_logs_id_seq
    OWNER TO postgres;
	
ALTER TABLE IF EXISTS public.authentication_logs
    ALTER COLUMN id SET DEFAULT nextval('authentication_logs_id_seq'::regclass);