--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-1.pgdg100+1)
-- Dumped by pg_dump version 13.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: transactions_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_categories (
    id integer NOT NULL,
    name character varying(25) NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions_categories OWNER TO postgres;

--
-- Name: transactions_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_categories_id_seq OWNER TO postgres;

--
-- Name: transactions_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_categories_id_seq OWNED BY public.transactions_categories.id;


--
-- Name: transactions_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_data (
    id integer NOT NULL,
    name character varying(25) NOT NULL,
    description character varying(255) NOT NULL,
    transaction_quote real NOT NULL,
    transaction_date date NOT NULL,
    transaction_type integer NOT NULL,
    transaction_category integer NOT NULL,
    transaction_recurence integer DEFAULT 0 NOT NULL,
    repeat_until date DEFAULT '1900-01-01'::date NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions_data OWNER TO postgres;

--
-- Name: transactions_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_data_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_data_id_seq OWNER TO postgres;

--
-- Name: transactions_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_data_id_seq OWNED BY public.transactions_data.id;


--
-- Name: transactions_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_logs (
    id integer NOT NULL,
    transaction_data integer NOT NULL,
    transaction_quote real NOT NULL,
    transaction_date date NOT NULL,
    created_by integer NOT NULL,
    updated_by integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions_logs OWNER TO postgres;

--
-- Name: transactions_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_logs_id_seq OWNER TO postgres;

--
-- Name: transactions_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_logs_id_seq OWNED BY public.transactions_logs.id;


--
-- Name: transactions_recurence; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_recurence (
    id integer NOT NULL,
    name character varying(25) NOT NULL,
    description character varying(255) NOT NULL,
    addtime character varying(25) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions_recurence OWNER TO postgres;

--
-- Name: transactions_recurence_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_recurence_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_recurence_id_seq OWNER TO postgres;

--
-- Name: transactions_recurence_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_recurence_id_seq OWNED BY public.transactions_recurence.id;


--
-- Name: transactions_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions_types (
    id integer NOT NULL,
    name character varying(25) NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.transactions_types OWNER TO postgres;

--
-- Name: transactions_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_types_id_seq OWNER TO postgres;

--
-- Name: transactions_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_types_id_seq OWNED BY public.transactions_types.id;


--
-- Name: transactions_categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_categories ALTER COLUMN id SET DEFAULT nextval('public.transactions_categories_id_seq'::regclass);


--
-- Name: transactions_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_data ALTER COLUMN id SET DEFAULT nextval('public.transactions_data_id_seq'::regclass);


--
-- Name: transactions_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_logs ALTER COLUMN id SET DEFAULT nextval('public.transactions_logs_id_seq'::regclass);


--
-- Name: transactions_recurence id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_recurence ALTER COLUMN id SET DEFAULT nextval('public.transactions_recurence_id_seq'::regclass);


--
-- Name: transactions_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_types ALTER COLUMN id SET DEFAULT nextval('public.transactions_types_id_seq'::regclass);


--
-- Name: transactions_categories transactions_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_categories
    ADD CONSTRAINT transactions_categories_pkey PRIMARY KEY (id);


--
-- Name: transactions_data transactions_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_data
    ADD CONSTRAINT transactions_data_pkey PRIMARY KEY (id);


--
-- Name: transactions_logs transactions_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_logs
    ADD CONSTRAINT transactions_logs_pkey PRIMARY KEY (id);


--
-- Name: transactions_recurence transactions_recurence_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_recurence
    ADD CONSTRAINT transactions_recurence_pkey PRIMARY KEY (id);


--
-- Name: transactions_types transactions_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_types
    ADD CONSTRAINT transactions_types_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: transactions_data transactions_data_transactions_categories_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_data
    ADD CONSTRAINT transactions_data_transactions_categories_id_fk FOREIGN KEY (transaction_category) REFERENCES public.transactions_categories(id);


--
-- Name: transactions_data transactions_data_transactions_recurence_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_data
    ADD CONSTRAINT transactions_data_transactions_recurence_id_fk FOREIGN KEY (transaction_recurence) REFERENCES public.transactions_recurence(id);


--
-- Name: transactions_data transactions_data_transactions_types_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_data
    ADD CONSTRAINT transactions_data_transactions_types_id_fk FOREIGN KEY (transaction_type) REFERENCES public.transactions_types(id);


--
-- Name: transactions_logs transactions_logs_transactions_data_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions_logs
    ADD CONSTRAINT transactions_logs_transactions_data_id_fk FOREIGN KEY (transaction_data) REFERENCES public.transactions_data(id);


--
-- PostgreSQL database dump complete
--

