-- Table: public.wallets

-- DROP TABLE IF EXISTS public.wallets;

CREATE TABLE IF NOT EXISTS public.wallets
(
    wallet_id integer NOT NULL ,
    user_id bigint NOT NULL,
    amount numeric NOT NULL,
    CONSTRAINT wallet_pkey PRIMARY KEY (wallet_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.wallets
    OWNER to postgres;
	
-- SEQUENCE: public.wallet_wallet_id_seq

-- DROP SEQUENCE IF EXISTS public.wallet_wallet_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.wallet_wallet_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY wallets.wallet_id;

ALTER SEQUENCE public.wallet_wallet_id_seq
    OWNER TO postgres;

ALTER TABLE IF EXISTS public.wallets
    ALTER COLUMN wallet_id SET DEFAULT nextval('wallet_wallet_id_seq'::regclass);