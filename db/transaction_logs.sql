-- Table: public.transaction_logs

-- DROP TABLE IF EXISTS public.transaction_logs;

CREATE TABLE IF NOT EXISTS public.transaction_logs
(
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    transaction_code character varying COLLATE pg_catalog."default" NOT NULL,
    transaction_type character varying COLLATE pg_catalog."default" NOT NULL,
    amount numeric NOT NULL,
    date_time time without time zone NOT NULL DEFAULT 'now()',
    CONSTRAINT transaction_logs_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.transaction_logs
    OWNER to postgres;
	
-- SEQUENCE: public.transaction_logs_id_seq

-- DROP SEQUENCE IF EXISTS public.transaction_logs_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.transaction_logs_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY transaction_logs.id;

ALTER SEQUENCE public.transaction_logs_id_seq
    OWNER TO postgres;
	
ALTER TABLE IF EXISTS public.transaction_logs
    ALTER COLUMN id SET DEFAULT nextval('transaction_logs_id_seq'::regclass);