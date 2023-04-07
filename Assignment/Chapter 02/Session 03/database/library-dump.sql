--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2023-04-04 01:27:11

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

--
-- TOC entry 3313 (class 1262 OID 58515)
-- Name: library; Type: DATABASE; Schema: -; Owner: jampirojam
--

CREATE DATABASE library WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Indonesia.1252';


ALTER DATABASE library OWNER TO jampirojam;

\connect library

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

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3314 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 66713)
-- Name: book; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.book (
    id integer NOT NULL,
    title text NOT NULL,
    author text,
    publisher text,
    year integer
);


ALTER TABLE public.book OWNER TO jampirojam;

--
-- TOC entry 209 (class 1259 OID 66712)
-- Name: book_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_id_seq OWNER TO jampirojam;

--
-- TOC entry 3315 (class 0 OID 0)
-- Dependencies: 209
-- Name: book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.book_id_seq OWNED BY public.book.id;


--
-- TOC entry 3164 (class 2604 OID 66716)
-- Name: book id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.book ALTER COLUMN id SET DEFAULT nextval('public.book_id_seq'::regclass);


--
-- TOC entry 3307 (class 0 OID 66713)
-- Dependencies: 210
-- Data for Name: book; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.book VALUES (1, 'Forensik Digital', 'Ojam', 'Hacktiv8', 2022);
INSERT INTO public.book VALUES (3, 'Cinta Searah', 'Aki P', 'AKB48', 2007);
INSERT INTO public.book VALUES (4, 'Cara Menjadi Jutawan', 'Mr. Omdo', 'Sura Press', 1996);
INSERT INTO public.book VALUES (5, 'SPSS', 'Mushlich', 'Hacktiv8', 2022);
INSERT INTO public.book VALUES (6, 'Hafalan Surat Denisa', 'Denisa', 'Tera Jana', 2004);
INSERT INTO public.book VALUES (7, 'Tendangan dari Laut', 'Tantri', 'Kotak', 2013);
INSERT INTO public.book VALUES (8, 'Malapetaka', 'Simple Girl', 'E-Pub', 2021);


--
-- TOC entry 3316 (class 0 OID 0)
-- Dependencies: 209
-- Name: book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.book_id_seq', 8, true);


--
-- TOC entry 3166 (class 2606 OID 66720)
-- Name: book book_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.book
    ADD CONSTRAINT book_pkey PRIMARY KEY (id);


-- Completed on 2023-04-04 01:27:12

--
-- PostgreSQL database dump complete
--

