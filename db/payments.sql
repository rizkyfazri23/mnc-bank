-- Table: public.payments

-- DROP TABLE IF EXISTS public.payments;

CREATE TABLE IF NOT EXISTS public.payments
(
    payment_id integer NOT NULL,
    payment_code character varying COLLATE pg_catalog."default" NOT NULL,
    sender_id bigint NOT NULL,
    receipt_id bigint NOT NULL,
    payment_amount numeric NOT NULL,
    description character varying COLLATE pg_catalog."default",
    date_time date NOT NULL,
    CONSTRAINT payment_pkey PRIMARY KEY (payment_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.payments
    OWNER to postgres;
	
-- SEQUENCE: public.payment_payment_id_seq

-- DROP SEQUENCE IF EXISTS public.payment_payment_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.payment_payment_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY payments.payment_id;

ALTER SEQUENCE public.payment_payment_id_seq
    OWNER TO postgres;
	
-- SEQUENCE: public.payment_seq

-- DROP SEQUENCE IF EXISTS public.payment_seq;

CREATE SEQUENCE IF NOT EXISTS public.payment_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.payment_seq
    OWNER TO postgres;
	
ALTER TABLE IF EXISTS public.payments
    ALTER COLUMN payment_id SET DEFAULT nextval('payment_payment_id_seq'::regclass);