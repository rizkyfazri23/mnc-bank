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

