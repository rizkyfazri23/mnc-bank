-- Table: public.deposits

-- DROP TABLE IF EXISTS public.deposits;

CREATE TABLE IF NOT EXISTS public.deposits
(
    deposit_id integer NOT NULL,
    deposit_code character varying COLLATE pg_catalog."default" NOT NULL,
    user_id bigint NOT NULL,
    deposit_amount numeric NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    date_time date NOT NULL,
    CONSTRAINT deposits_pkey PRIMARY KEY (deposit_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.deposits
    OWNER to postgres;

-- SEQUENCE: public.deposit_seq

-- DROP SEQUENCE IF EXISTS public.deposit_seq;

CREATE SEQUENCE IF NOT EXISTS public.deposit_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.deposit_seq
    OWNER TO postgres;

-- SEQUENCE: public.deposits_deposit_id_seq

-- DROP SEQUENCE IF EXISTS public.deposits_deposit_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.deposits_deposit_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY deposits.deposit_id;

ALTER SEQUENCE public.deposits_deposit_id_seq
    OWNER TO postgres;

ALTER TABLE IF EXISTS public.deposits
    ALTER COLUMN deposit_id SET DEFAULT nextval('deposits_deposit_id_seq'::regclass);